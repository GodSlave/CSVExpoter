package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"strings"
	"errors"
	"os/exec"
	"os"
	"path/filepath"
	"flag"
	"io/ioutil"
	"html/template"
	"bufio"
)

var delimiter = flag.String("d", ",", "Delimiter to use between fields")

type outputer func(s string)

type AllModule struct {
	All          []Module
	Enums        []Module
	KeyMapModule Module
	Responses    []Module
}

type Module struct {
	Name         string
	Attributes   []Attr
	Content      []Contents
	HasPrimalKey bool
}

type Contents struct {
	Values []string
	Name   string
}

type Attr struct {
	Name        string
	Desc        string
	Type        string
	IsPrimalKey bool
	IsArray     bool
}

type AllEnum struct {
	Name string
	Id   int32
}

var ClientTypeMap map[string]string
var ServerTypeMap map[string]string
var TableNameMap map[string]string
var nameSort []string
var ClientAllModule *AllModule
var ServerAllModule *AllModule

func main() {
	ClientTypeMap = map[string]string{"i": "int32", "f": "float", "s": "FName", "b": "bool", "e": "EVictoryEnum", "a": "objectArray", "o": "object"}
	ServerTypeMap = map[string]string{"i": "int32", "f": "float", "s": "string", "b": "bool", "e": "int32", "a": "objectArray", "o": "object"}
	TableNameMap = map[string]string{}

	file, _ := exec.LookPath(os.Args[0])
	ApplicationPath, _ := filepath.Abs(file)
	ApplicationDir, _ := filepath.Split(ApplicationPath)
	var err error
	loadNameMapTable(ApplicationDir)
	err = os.MkdirAll(ApplicationDir+"out", os.ModePerm)
	err = os.MkdirAll(ApplicationDir+"out"+string(filepath.Separator)+"client"+string(filepath.Separator)+"CODE", os.ModePerm)
	err = os.MkdirAll(ApplicationDir+"out"+string(filepath.Separator)+"client"+string(filepath.Separator)+"CSV", os.ModePerm)
	err = os.MkdirAll(ApplicationDir+"out"+string(filepath.Separator)+"server", os.ModePerm)
	checkError(err)
	ClientAllModule = &AllModule{}
	ServerAllModule = &AllModule{}
	process(ApplicationDir, ApplicationDir+"out"+string(filepath.Separator)+"client"+string(filepath.Separator)+"CSV", true)
	for _, module := range ClientAllModule.All {
		generateClientCSVFile(module, ApplicationDir+"out"+string(filepath.Separator)+"client"+string(filepath.Separator)+"CSV")
	}
	GenerateClientHeadFile(ClientAllModule, ApplicationDir+"out"+string(filepath.Separator)+"client"+string(filepath.Separator)+"CODE")
	generateClientEnumFile(ClientAllModule, ApplicationDir+"out"+string(filepath.Separator)+"client"+string(filepath.Separator)+"CODE")
	generateClientConstValueCppFile(ClientAllModule, ApplicationDir+"out"+string(filepath.Separator)+"client"+string(filepath.Separator)+"CODE")
	generateClientConstValueHeadFile(ClientAllModule, ApplicationDir+"out"+string(filepath.Separator)+"client"+string(filepath.Separator)+"CODE")

	process(ApplicationDir, ApplicationDir+"out"+string(filepath.Separator)+"server", false)
	generateServerFile(ServerAllModule, ApplicationDir+"out"+string(filepath.Separator)+"server")
	generateServerResponsFile(ServerAllModule, ApplicationDir+"out"+string(filepath.Separator)+"server")
	generateServerInitDataFile(ServerAllModule, ApplicationDir+"out"+string(filepath.Separator)+"server")
}

func loadNameMapTable(dirPath string) {
	xlFile, error := xlsx.OpenFile(dirPath + string(filepath.Separator) + "对照表.xlsx")
	checkError(error)
	if len(xlFile.Sheets) < 1 {
		fmt.Println("对照表内没有数据")
	}
	sheet := xlFile.Sheets[0]
	for _, row := range sheet.Rows {
		var key string
		var value string
		for rowIndex, cell := range row.Cells {
			str, err := cell.FormattedValue()
			if err != nil {
				str = err.Error()
			}
			switch rowIndex {
			case 0:
				key = str
			case 1:
				value = str
			}
		}
		TableNameMap[key] = value
		nameSort = append(nameSort, value)
	}
}

func generateClientCSVFromXLSXFile2(xlFile *xlsx.File, sheetIndex int, outputf outputer, isEnum bool, outFileName string) (error) {
	module := Module{}
	enum := Module{}
	sheet := xlFile.Sheets[sheetIndex]
	keymap := map[int]string{}
	rowNameMap := map[int]string{}
	var allDatas []string

	for index, row := range sheet.Rows {
		var vals []string
		if index == 2 {
			continue
		}

		if row != nil {
			var attr Attr
			var enumAttr = Attr{}
			for rowIndex, cell := range row.Cells {
				attr = Attr{}
				str, err := cell.FormattedValue()
				if err != nil {
					str = err.Error()
				}
				if index == 0 {
					rowNameMap[rowIndex] = str
				} else if index == 1 {
					if str != "" {
						keymap[rowIndex] = str
						Name := str[strings.IndexByte(str, '_')+1:]
						if strings.HasSuffix(Name, "*") {
							attr.IsPrimalKey = true
							module.HasPrimalKey = true
							Name = Name[0:len(Name)-1]
						}
						if strings.HasSuffix(Name, "[]") {
							attr.IsArray = true
							Name = Name[0:len(Name)-2]
						}
						attr.Name = Name
						attr.Desc = rowNameMap[rowIndex]
						contentTypeKey := str[0:strings.IndexByte(str, '_')]
						if contentType, ok := ClientTypeMap[contentTypeKey]; ok {
							attr.Type = contentType
						} else {
							attr.Type = contentTypeKey
							fmt.Println(attr.Type)
						}
						vals = append(vals, Name)
						module.Attributes = append(module.Attributes, attr)
					}
				} else {
					if _, ok := keymap[rowIndex]; ok {
						vals = append(vals, str)
						if index >= 3 {
							switch rowIndex {
							case 0:
								enumAttr.Type = str
							case 1:
								enumAttr.Name = str
							case 2:
								enumAttr.Desc = str
							}
						}
					}
				}
			}
			if index == 0 {
				continue
			}
			if enumAttr.Desc == "" {
				enumAttr.Desc = attr.Name
			}
			if enumAttr.Name != "" {
				if isEnum {
					enum.Attributes = append(enum.Attributes, enumAttr)
				}
			}
			content := strings.Join(vals, *delimiter) + "\n"
			if index != 1 {
				if module.HasPrimalKey {
					module.Content = append(module.Content,
						Contents{vals, ParsPrimalKey(module.Attributes, vals),})
				} else {
					module.Content = append(module.Content,
						Contents{vals, "",})
				}
			}
			allDatas = append(allDatas, content)
		}
	}
	for _, value := range allDatas {
		if outputf != nil {
			outputf(value)
		}
	}

	if isEnum {
		enum.Name = outFileName
		ClientAllModule.Enums = append(ClientAllModule.Enums, enum)
	} else {
		module.Name = outFileName
		ClientAllModule.All = append(ClientAllModule.All, module)
	}

	return nil
}

func generateServerCSVFromXLSXFile2(xlFile *xlsx.File, sheetIndex int, outputf outputer, isEnum bool, outFileName string) (error) {
	module := Module{}
	enum := Module{}
	sheetLen := len(xlFile.Sheets)
	switch {
	case sheetLen == 0:
		return errors.New("This XLSX file contains no sheets.")
	case sheetIndex >= sheetLen:
		return fmt.Errorf("No sheet %d available, please select a sheet between 0 and %d\n", sheetIndex, sheetLen-1)
	}
	sheet := xlFile.Sheets[sheetIndex]
	keymap := map[int]string{}
	var allDatas []string
	for index, row := range sheet.Rows {
		var vals []string
		if index == 0 {
			//allDatas = append(allDatas, "")
			continue
		}
		if row != nil {
			attr := Attr{}
			for rowIndex, cell := range row.Cells {

				str, err := cell.FormattedValue()
				if err != nil {
					//vals = append(vals, err.Error())
					str = err.Error()
				}
				if index == 1 {
					keymap[rowIndex] = str
				} else if index == 2 {
					if str == "" {
						str = keymap[rowIndex]
					} else {
						keymap[rowIndex] = str
					}
					str = str[strings.IndexByte(str, '_')+1:]
					vals = append(vals, str)
				} else {
					if _, ok := keymap[rowIndex]; ok {
						vals = append(vals, str)

						if index >= 3 {
							switch rowIndex {
							case 0:
								attr.Type = str
							case 1:
								attr.Name = str
							case 2:
								attr.Desc = str
							}
						}
					}
				}
			}
			if attr.Desc == "" {
				attr.Desc = attr.Name
			}

			if isEnum && attr.Name != "" {
				enum.Attributes = append(enum.Attributes, attr)
			}

			if index == 1 {
				continue
			}
			if index != 2 {
				module.Content = append(module.Content, Contents{vals, "testName"})
			}
			content := strings.Join(vals, *delimiter)
			fmt.Println(content)
			allDatas = append(allDatas, content+"\n")

		}
	}
	for _, value := range allDatas {
		//fmt.Printf("%s", value)
		if outputf != nil {
			outputf(value)
		}
	}
	var descVals []string
	if len(sheet.Rows) > 0 {
		for rowIndex, cell := range sheet.Rows[0].Cells {
			if indeName, ok := keymap[rowIndex]; ok {
				str, err := cell.FormattedValue()
				if err != nil {
					str = err.Error()
				}
				descVals = append(descVals, str)
				if !isEnum && len(indeName) > 0 {
					attr := Attr{}
					Name := indeName[strings.IndexByte(indeName, '_')+1:]
					if strings.HasSuffix(Name, "*") {
						attr.IsPrimalKey = true
						module.HasPrimalKey = true
						Name = Name[0:len(Name)-1]
					}
					if strings.HasSuffix(Name, "[]") {
						attr.IsArray = true
						Name = Name[0:len(Name)-2]
					}
					attr.Name = Name
					contentTypeKey := indeName[0:strings.IndexByte(indeName, '_')]
					var TypeMap = map[string]string{}
					TypeMap = ServerTypeMap
					if contentType, ok := TypeMap[contentTypeKey]; ok {
						attr.Type = contentType
					} else {
						attr.Type = contentTypeKey
					}
					attr.Desc = str
					module.Attributes = append(module.Attributes, attr)
				}
			}
		}
	}

	module.Name = outFileName
	enum.Name = outFileName
	if isEnum {
		ServerAllModule.Enums = append(ServerAllModule.Enums, enum)
	} else {
		ServerAllModule.All = append(ServerAllModule.All, module)
	}

	return nil
}

func process(dirPath string, outdir string, clientMode bool) error {
	files, error := ioutil.ReadDir(dirPath)
	checkError(error)
	for _, info := range files {
		if info.IsDir() {
			if info.Name()[0] != '.' && info.Name() != "out" {
				process(dirPath+string(filepath.Separator)+info.Name(), outdir+string(filepath.Separator)+info.Name(),
					clientMode)
			}
		} else {
			if len(info.Name()) > 4 && info.Name()[0:2] != "~$" && info.Name() != "对照表.xlsx" {
				switch {
				case strings.HasSuffix(info.Name(), "配置.xlsx"):
					readKeyMap(dirPath+string(filepath.Separator)+info.Name(), clientMode)
				case strings.HasPrefix(info.Name(), "协议响应状态枚举"):

					readResponses(dirPath+string(filepath.Separator)+info.Name(), clientMode)
				case strings.HasSuffix(info.Name(), "xlsx"):
					fmt.Println(info.Name())
					xlFile, error := xlsx.OpenFile(dirPath + string(filepath.Separator) + info.Name())
					checkError(error)
					sheetLen := len(xlFile.Sheets)
					isEnum := strings.HasSuffix(info.Name(), "枚举.xlsx")
					switch {
					case sheetLen == 0:
						return errors.New("This XLSX file contains no sheets.")
					case sheetLen == 1:
						outPutFileName := info.Name()[0:strings.LastIndexByte(info.Name(), '.')]
						if contentType, ok := TableNameMap[outPutFileName]; ok {
							outPutFileName = contentType
						}
						if isEnum {
							outPutFileName = xlFile.Sheets[0].Name
						}
						DoGenerateFile(xlFile, 0, isEnum, outPutFileName, clientMode)
					case sheetLen > 1:
						for i := 0; i < sheetLen; i++ {
							sheet := xlFile.Sheets[i]
							outPutFileName := sheet.Name
							if contentType, ok := TableNameMap[outPutFileName]; ok {
								outPutFileName = contentType
							}
							DoGenerateFile(xlFile, i, isEnum, outPutFileName, clientMode)
						}
					}

				}

			}
		}
	}

	return nil
}

func DoGenerateFile(xlFile *xlsx.File, sheetIndex int, isEnum bool, outFileName string, clientMode bool) error {
	var error error
	if clientMode {
		error = generateClientCSVFromXLSXFile2(xlFile, sheetIndex, nil, isEnum, outFileName)
	} else {
		error = generateServerCSVFromXLSXFile2(xlFile, sheetIndex, nil, isEnum, outFileName)
	}

	checkError(error)
	return nil
}

func GenerateClientHeadFile(all *AllModule, path string) error {
	var alll []Module
	for _, content := range nameSort {
		for _, module := range all.All {
			if module.Name == content {
				alll = append(alll, module)
				fmt.Println(content)
			}
		}
	}
	all.All = alll
	tpl, err := template.New("client_struct.template").Funcs(template.FuncMap{
		"checkIsEnum": checkIsEnum,
	}).Parse(ClientStruct)
	checkError(err)
	var printer = Printer{}
	printer.f, err = os.OpenFile(path+string(filepath.Separator)+"GeneratedStructs.h", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	err = tpl.Execute(&printer, all)
	checkError(err)
	return err
}

func generateClientConstValueCppFile(all *AllModule, path string) {
	tpl, err := template.New("client_struct.template").Parse(client_keymap_content)
	checkError(err)
	var printer = Printer{}
	printer.f, err = os.OpenFile(path+string(filepath.Separator)+"GeneratedConstVariables.cpp", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	err = tpl.Execute(&printer, all)
	checkError(err)
}

func generateClientConstValueHeadFile(all *AllModule, path string) {
	tpl, err := template.New("client_struct.template").Parse(client_keymap_head)
	checkError(err)
	var printer = Printer{}
	printer.f, err = os.OpenFile(path+string(filepath.Separator)+"GeneratedConstVariables.h", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	err = tpl.Execute(&printer, all)
	checkError(err)
}

func generateClientEnumFile(all *AllModule, path string) error {
	tpl, err := template.New("client_enum.template").Parse(ClientEnum)
	checkError(err)
	var printer = Printer{}
	printer.f, err = os.OpenFile(path+string(filepath.Separator)+"GeneratedEnums.h", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	err = tpl.Execute(&printer, all)
	checkError(err)
	return err
}

func generateServerFile(all *AllModule, path string) error {
	tpl, err := template.New("server_struct.template").Funcs(template.FuncMap{
		"generateContent":   parseForServer,
		"generatePrimalKey": ParsPrimalKey,
	}).Parse(ServerStruct)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	var printer = Printer{}
	printer.f, err = os.OpenFile(path+string(filepath.Separator)+"bean.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	err = tpl.Execute(&printer, all)
	checkError(err)
	return err
}

func generateServerResponsFile(all *AllModule, path string) error {
	tpl, err := template.New("server_struct.template").Parse(server_response_content)
	checkError(err)
	var printer = Printer{}
	printer.f, err = os.OpenFile(path+string(filepath.Separator)+"errorCode.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	err = tpl.Execute(&printer, all)
	checkError(err)
	return nil
}

func generateServerInitDataFile(all *AllModule, path string) error {
	tpl, err := template.New("server_struct.template").Funcs(template.FuncMap{
		"generateContent":   parseForServer,
		"generatePrimalKey": ParsPrimalKey,
	}).Parse(server_InitData_content)
	checkError(err)
	var printer = Printer{}
	printer.f, err = os.OpenFile(path+string(filepath.Separator)+"initData.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	err = tpl.Execute(&printer, all)
	checkError(err)
	return nil
}

type Printer struct {
	f *os.File
}

func (printer *Printer) Write(p []byte) (n int, err error) {
	content := string(p)
	content = strings.Replace(content, "&#34;", "\"", -1)
	content = strings.Replace(content, "&lt;", "<", -1)
	_, err = printer.f.Write([]byte(content))
	if err != nil {
		fmt.Println(err.Error())
	}
	return 0, err
}

func checkIsEnum(Type string) bool {
	if len(Type) > 2 && Type[0] == 'E' && Type[1] >= 'A' && Type[1] <= 'Z' {
		return true
	}
	return false
}

func parseForServer(types [] Attr, contents []string) string {
	len1 := len(types)
	len2 := len(contents)
	if len1 != len2 {
		fmt.Println("Attr && content length not match")
	}
	result := ""
	for i := 0; i < len1; i++ {
		var content string
		switch types[i].Type {
		case "string":
			if types[i].IsArray {
				value := contents[i]
				value = value[2:len(value)-2]
				subcontents := strings.Split(value, ",")
				var realSubs []string
				for _, subContent := range subcontents {
					realSubs = append(realSubs, fmt.Sprintf(`"%s"`, subContent))
				}
				content = strings.Join(realSubs, ",")
				content = fmt.Sprintf("[]string{%s}", content)
			} else {
				content = fmt.Sprintf(`"%s"`, contents[i])
			}
		case "int32", "float", "int64":
			if types[i].IsArray {
				content = fmt.Sprintf("[]%s{%s}", types[i].Type, contents[i][2:len(contents)-2])
			} else {
				content = contents[i]
			}
		default:
			if types[i].IsArray {
				value := contents[i]
				if strings.HasPrefix(value, "\"(") && strings.HasSuffix(value, ")\"") {
					value = value[2:len(value)-2]
					var subValues []string
					if value[0] == '(' {
						subValues = strings.Split(value, "),(")
						subValues[0] = subValues[0][1:]
						subValues[len(subValues)-1] = subValues[len(subValues)-1][0: len(subValues[len(subValues)-1])-1]
					} else {
						subValues = strings.Split(value, ",")
					}

					ItemType := types[i].Type
					var matched = false
					for _, module := range ServerAllModule.All {
						if module.Name == ItemType {
							for _, subValue := range subValues {
								subSubValues := strings.Split(subValue, ",")
								tempresult := parseForServer(module.Attributes, subSubValues)
								content += fmt.Sprintf("{%s},", tempresult)
							}
							matched = true
							break
						}
					}
					if !matched {
						content = strings.Join(subValues, ",")
						fmt.Println("can't find assign Type")
					}
					content = fmt.Sprintf("[]%s{%s}", ItemType, content)

				} else {
					return "format Error"
				}
			} else {

				value := contents[i]
				if len(value) < 4 {
					content = value
				} else {
					value = value[2:len(value)-2]
					ItemType := types[i].Type
					var matched = false
					for _, module := range ServerAllModule.All {
						if module.Name == ItemType {
							subSubValues := strings.Split(value, ",")
							tempresult := parseForServer(module.Attributes, subSubValues)
							content = fmt.Sprintf("%s{%s},", ItemType, tempresult)
							matched = true
							break
						}
					}
					if !matched {
						content = fmt.Sprintf("{%s}", value)
						fmt.Println("can't find assign Type")
					}
				}

			}
		}
		result += content + ","
	}
	return result
}

func ParsPrimalKey(types [] Attr, contents []string) string {
	len1 := len(types)
	len2 := len(contents)
	if len1 != len2 {
		fmt.Println("error lenth not match")
	}
	result := ""
	for i := 0; i < len1; i++ {
		if (types[i].IsPrimalKey) {
			result += contents[i]
			result += "_"
		}
	}
	if len(result) > 0 {
		result = result[0:len(result)-1]
	}
	return result
}

func generateClientCSVFile(module Module, outPutFileDir string) error {
	var allContent []string
	var firstLine string
	if module.HasPrimalKey {
		firstLine = ","
	} else {
		firstLine = ""
	}

	for _, ContentType := range module.Attributes {
		firstLine += ContentType.Name + ","
	}
	if len(firstLine) > 0 {
		firstLine = firstLine[0:len(firstLine)-1]
	}
	allContent = append(allContent, firstLine)

	for _, content := range module.Content {
		lineContent := ""
		if module.HasPrimalKey {
			lineContent = content.Name + ","
		}
		lineContent += buildClientCSVContent(module.Attributes, content.Values, false)
		allContent = append(allContent, lineContent)
	}

	result := strings.Join(allContent, "\n")
	error := ioutil.WriteFile(outPutFileDir+string(filepath.Separator)+module.Name+".csv", []byte(result), os.ModePerm)
	checkError(error)
	return nil
}

func buildClientCSVContent(types [] Attr, contents []string, withName bool) string {
	len1 := len(types)
	len2 := len(contents)
	if len1 != len2 {
		fmt.Println("error length not match")
	}
	result := ""
	for i := 0; i < len1; i++ {
		attr := types[i]
		var content string;
		switch attr.Type {
		case "FName":

			content = fmt.Sprintf("\"%s\"", contents[i])

		case "int32", "float", "bool":
			content = contents[i]
		default:

			ItemType := attr.Type
			var matched = false
			for _, module := range ClientAllModule.All {
				if module.Name == ItemType {
					if attr.IsArray {

						subValues := strings.Split(contents[i], "),(")
						subValues[0] = subValues[0][3:]
						subValues[len(subValues)-1] = subValues[len(subValues)-1][0: len(subValues[len(subValues)-1])-3]
						for _, subValue := range subValues {
							subsubValues := strings.Split(subValue, ",")
							tempContent := buildClientCSVContent(module.Attributes, subsubValues, true)
							tempContent = tempContent[0:len(tempContent)-1]
							tempContent = fmt.Sprintf("(%s),", tempContent)
							content += tempContent
						}
						content = content[0:len(content)-1]
					} else {
						subValue := contents[i][2:len(contents[i])-2]
						subValues := strings.Split(subValue, ",")
						content = buildClientCSVContent(module.Attributes, subValues, true)
						content = content[0:len(content)-1]
					}
					content = fmt.Sprintf("\"(%s)\"", content)
					matched = true
				}
			}

			for _, enum := range ClientAllModule.Enums {
				if enum.Name == ItemType {
					if attr.IsArray {
						currentContent := contents[i];
						currentContent = currentContent[2: len(currentContent)-2]
						subValues := strings.Split(currentContent, ",")
						for _, subValue := range subValues {
							for _,enumItem := range enum.Attributes{
								if subValue == enumItem.Type{
									content += enumItem.Name
									content += ","
									matched = true
								}
							}
						}
						content = content[0:len(content)-1]
						content = fmt.Sprintf("\"(%s)\"", content)
					} else {
						for _,enumItem := range enum.Attributes{
							if  contents[i] == enumItem.Type{
								content = enumItem.Name
								matched = true
							}
						}
					}
				}
			}

			if !matched {
				content = contents[i]
			}
		}
		if withName {
			result += attr.Name + "=" + content + ","
		} else {
			result += content + ","
		}
	}
	if len(result) > 1 {
		result = result[0:len(result)-1]
	}

	return result
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
	}

}

func readKeyMap(filePath string, isClientMode bool) {
	xlFile, error1 := xlsx.OpenFile(filePath)
	checkError(error1)
	sheetLen := len(xlFile.Sheets)

	for i := 0; i < sheetLen; i++ {
		sheet := xlFile.Sheets[i]
		for _, row := range sheet.Rows {
			attr := Attr{}
			var err error
			for index, cell := range row.Cells {
				switch index {
				case 0:
					attr.Name, err = cell.FormattedValue()
				case 1:
					attr.Type, err = cell.FormattedValue()
					checkError(err)
					if isClientMode {
						if contentType, ok := ClientTypeMap[attr.Type]; ok {
							attr.Type = contentType
						}
					} else {
						if contentType, ok := ServerTypeMap[attr.Type]; ok {
							attr.Type = contentType
						}
					}
				case 2:
					attr.Desc, err = cell.FormattedValue()
					switch attr.Type {
					case "int32", "float", "int64":

					case "bool":
						if attr.Desc == "1" {
							attr.Desc = "true"
						} else {
							attr.Desc = "false"
						}
					case "string":
						attr.Desc = fmt.Sprintf("`%s`", attr.Desc)
					case "FName":
						attr.Desc = strings.Replace(attr.Desc, "\"", "\\\"", -1)
						attr.Desc = fmt.Sprintf("FName(\"%s\")", attr.Desc)
					}
				}
			}
			if isClientMode {
				ClientAllModule.KeyMapModule.Attributes = append(ClientAllModule.KeyMapModule.Attributes, attr)
			} else {
				ServerAllModule.KeyMapModule.Attributes = append(ServerAllModule.KeyMapModule.Attributes, attr)
			}
		}

	}

}

func readResponses(filePath string, isClientMode bool) {
	xlFile, error1 := xlsx.OpenFile(filePath)
	checkError(error1)
	sheetLen := len(xlFile.Sheets)

	for i := 0; i < sheetLen; i++ {
		sheet := xlFile.Sheets[i]
		responsModule := Module{}
		if isClientMode {
			responsModule.Name = fmt.Sprintf("E%sResponseState", sheet.Name)
		} else {
			responsModule.Name = sheet.Name
		}
		for _, row := range sheet.Rows {
			attr := Attr{}
			var err error
			var value string
			for index, cell := range row.Cells {
				value, err = cell.FormattedValue()
				checkError(err)
				switch index {
				case 0:
					attr.Type = value
				case 1:
					attr.Name = value
				case 2:
					attr.Desc = value
				}
			}
			responsModule.Attributes = append(responsModule.Attributes, attr)
		}
		if isClientMode {
			ClientAllModule.Responses = append(ClientAllModule.Responses, responsModule)
		} else {
			ServerAllModule.Responses = append(ServerAllModule.Responses, responsModule)
		}
	}
}
