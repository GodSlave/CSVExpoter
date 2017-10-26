package main

var ClientEnum= `// Fill out your copyright notice in the Description page of Project Settings.

#pragma once

#include "CoreMinimal.h"
#include "GeneratedEnums.generated.h"

{{range $index,$A := .Enums }}
UENUM(BlueprintType)
enum class E{{$A.Name}}Enum : uint8
{
	{{range $index,$B := $A.Attributes }}
	VE_{{$B.Name}} = {{$B.Type}}	UMETA(DisplayName = "{{$B.Desc}}"),
	{{end }}
};
{{end }}`
var ClientStruct=`// Fill out your copyright notice in the Description page of Project Settings.

#pragma once

#include "CoreMinimal.h"
#include "GameFramework/Character.h"
#include "AttributeSet.h"
#include "AbilitySystemInterface.h"
#include "GeneratedEnums.h"
#include "GeneratedStructs.generated.h"

{{range $index,$A := .All }}
USTRUCT(BlueprintType)
struct F{{$A.Name}}Data: public FTableRowBase
{
	GENERATED_USTRUCT_BODY()

public:

	F{{$A.Name}}Data()
	{}
	{{range $index,$V := $A.Attributes }}
	UPROPERTY(EditAnywhere, BlueprintReadWrite, Category = "{{$A.Name}}", meta = (DisplayName = "{{$V.Desc}}"))
	{{if $V.IsArray}}
	TArray<{{$V.Type}}> {{$V.Name}};{{else}}
	{{$V.Type}} {{$V.Name}};
	{{end}}{{end}}

};
{{end}}`

var ServerStruct=`package bean


// model
{{range $index,$A := .All }}
type {{$A.Name}} struct {
{{range $index,$V := $A.Attributes }}   {{if $V.IsArray}}{{$V.Name}}     []{{$V.Type}}//{{$V.Desc}}
{{else}}{{$V.Name}}     {{$V.Type}}//{{$V.Desc}}
{{end}}{{end}}
}{{end}}

{{range $index,$A := .Enums }}
//{{$A.Name}}
{{range $index,$V := $A.Attributes }}var {{$A.Name}}_{{$V.Name}} = {{$V.Type}} // {{$V.Desc}}
{{end}}{{end}}

{{range $index,$A := .All }}{{if $A.HasPrimalKey}}var {{$A.Name}}s  map[string]{{$A.Name}}{{else}}var {{$A.Name}}s  []{{$A.Name}}{{end}}
{{end}}

func init()  {
{{range $index,$A := .All }}
   {{if $A.HasPrimalKey}}
   {{range $index,$V := $A.Content }}
   {{$A.Name}}s["{{generatePrimalKey $A.Attributes $V.Values}}"]= {{$A.Name}}{ {{generateContent $A.Attributes $V.Values}} } {{end}}
   {{else}}
   {{range $index,$V := $A.Content }}
   {{$A.Name}}s = append({{$A.Name}}s , {{$A.Name}}{ {{generateContent $A.Attributes $V.Values}} }){{end}}
   {{end}}{{end}}
}
`
