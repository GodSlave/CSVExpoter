package main

var ClientEnum = `// Fill out your copyright notice in the Description page of Project Settings.

#pragma once

#include "CoreMinimal.h"
#include "GeneratedEnums.generated.h"
{{range $index,$A := .Responses }}
UENUM(BlueprintType)
enum class {{$A.Name}} : uint8
{
	{{range $index,$C := $A.Attributes }}
	{{$C.Name}} = {{$C.Type}}	UMETA(DisplayName = "{{$C.Desc}}"),
	{{end }}
};
{{end }}
{{range $index,$A := .Enums }}
UENUM(BlueprintType)
enum class {{$A.Name}} : uint8
{
	{{range $index,$B := $A.Attributes }}
	{{$B.Name}} = {{$B.Type}}	UMETA(DisplayName = "{{$B.Desc}}"),
	{{end }}
};
{{end }}`
var ClientStruct = `// Fill out your copyright notice in the Description page of Project Settings.

#pragma once

#include "Engine.h"
#include "GameFramework/Character.h"
#include "AttributeSet.h"
#include "AbilitySystemInterface.h"
#include "GeneratedEnums.h"
#include "PussyDataStructs.h"
#include "GeneratedStructs.generated.h"

{{range $index,$A := .All }}
USTRUCT(BlueprintType)
struct F{{$A.Name}}: public FTableRowBase
{
	GENERATED_USTRUCT_BODY()

public:

	F{{$A.Name}}()
	{}
	{{range $index,$V := $A.Attributes }}
	UPROPERTY(EditAnywhere, BlueprintReadWrite, Category = "{{$A.Name}}", meta = (DisplayName = "{{$V.Desc}}"))
	{{if $V.IsArray}}
	{{if checkIsEnum $V.Type}}
	TArray< {{$V.Type}}> {{$V.Name}};{{else}}
	TArray< F{{$V.Type}}> {{$V.Name}};
	{{end}}
	{{else}}
	{{$V.Type}} {{$V.Name}};
	{{end}}{{end}}

};

UCLASS(Blueprintable)
class CLOUD_API U{{$A.Name}}ItemData : public UPussyItemData
{
	GENERATED_BODY()

public:
	UPROPERTY(EditAnywhere, BlueprintReadWrite, Category = "ItemBase", Meta = (DisplayName = "Data", ExposeOnSpawn = true))
		F{{$A.Name}} Data;
};

{{end}}`

var ServerStruct = `package bean

import (
	"github.com/GodSlave/MyGoServer/module"
	"github.com/go-xorm/xorm"
)

// model
{{range $index,$A := .All }}
type {{$A.Name}} struct {
{{range $index,$V := $A.Attributes }}   {{if $V.IsArray}}{{$V.Name}}     []{{$V.Type}}  "xorm:extends"//{{$V.Desc}}
{{else if $V.IsPrimalKey}}{{$V.Name}}     {{$V.Type}} "xorm:pk"//{{$V.Desc}}
{{else if  isLargeString ($V.Type) }}{{$V.Name}}     string "xorm:\"varchar(2048)\""//{{$V.Desc}}
{{else}}{{$V.Name}}     {{$V.Type}}//{{$V.Desc}}
{{end}}{{end}}
}{{end}}

{{range $index,$A := .KeyMapModule.Attributes }}
const {{$A.Name}} = {{$A.Desc}}
{{end}}

{{range $index,$A := .Enums }}
//{{$A.Name}}
type {{$A.Name}} int32
{{range $index,$V := $A.Attributes }}const {{$A.Name}}_{{$V.Name}} = {{$V.Type}} // {{$V.Desc}}
{{end}}{{end}}

func EnableDBCache(app module.App) {
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
{{range $index,$A := .All }}
	app.GetSqlEngine().MapCacher(&{{$A.Name}}{}, cacher){{end}}
}

func ClearDBChache(app module.App) {
{{range $index,$A := .All }}
	app.GetSqlEngine().ClearCache(&{{$A.Name}}{}){{end}}
}

func DisableDBCache(app module.App) {
{{range $index,$A := .All }}
	app.GetSqlEngine().MapCacher(&{{$A.Name}}{}, nil){{end}}
}

`
var client_keymap_head = `// Fill out your copyright notice in the Description page of Project Settings.

#pragma once

#include "CoreMinimal.h"
#include "Kismet/BlueprintFunctionLibrary.h"
#include "GeneratedConstVariables.generated.h"

UCLASS()
class CLOUD_API UGeneratedConstVariables : public UBlueprintFunctionLibrary
{
	GENERATED_BODY()
public:
	{{range $index,$A := .KeyMapModule.Attributes }}
	static  const {{$A.Type}} {{$A.Name}};

	UFUNCTION(BlueprintPure, Category = "Const Variables", meta = (DisplayName = "{{$A.Name}}"))
	static	{{$A.Type}} Get{{$A.Name}}() { return {{$A.Name}}; }
	{{end}}
};
`
var client_keymap_content = `
// Fill out your copyright notice in the Description page of Project Settings.
#include "GeneratedConstVariables.h"
{{range $index,$A := .KeyMapModule.Attributes }}
const {{$A.Type}} UGeneratedConstVariables::{{$A.Name}} = {{$A.Desc}};
{{end}}
`

var server_response_content = `
package bean

import "github.com/GodSlave/MyGoServer/base"

var (
{{range $index,$A := .Responses }}
	{{range $index1,$B :=  $A.Attributes }}
	{{$A.Name}}_{{$B.Name}}     = base.NewError({{$B.Type}}, "{{$B.Desc}}")
	{{end}}
{{end}}
)`

var server_InitData_content = `
package bean

import (
	"github.com/GodSlave/MyGoServer/db"
	"os/exec"
	"os"
	"path/filepath"
	"fmt"
	"flag"
	"github.com/GodSlave/MyGoServer/conf"
)

func Init()  {
	file, _ := exec.LookPath(os.Args[0])
	ApplicationPath, _ := filepath.Abs(file)
	ApplicationDir, _ := filepath.Split(ApplicationPath)
	defaultPath := fmt.Sprintf("%sconf"+string(filepath.Separator)+"server.json", ApplicationDir)
	confPath := flag.String("conf", defaultPath, "Server configuration file path")
	flag.Parse() //解析输入的参数
	f, err := os.Open(*confPath)
	if err != nil {
		panic(err)
	}
	conf.LoadConfig(f.Name()) //加载配置文件
	//sql
	sql := db.BaseSql{
	}
	sql.Url = conf.Conf.DB.SQL
	sql.InitDB()
	sql.CheckMigrate()
	defer sql.Engine.Close()

{{range $index,$A := .All }}
	sql.Engine.DropTables(&{{$A.Name}}{})
	sql.Engine.Sync2(&{{$A.Name}}{})
   {{range $index,$V := $A.Content }}
 	sql.Engine.Insert( {{$A.Name}}{ {{generateContent $A.Attributes $V.Values}} })
   {{end}}{{end}}
}
`
