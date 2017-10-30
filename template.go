package main

var ClientEnum = `// Fill out your copyright notice in the Description page of Project Settings.

#pragma once

#include "CoreMinimal.h"
#include "GeneratedEnums.generated.h"

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

#include "CoreMinimal.h"
#include "GameFramework/Character.h"
#include "AttributeSet.h"
#include "AbilitySystemInterface.h"
#include "GeneratedEnums.h"
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
	TArray< F{{$V.Type}}> {{$V.Name}};{{else}}
	{{$V.Type}} {{$V.Name}};
	{{end}}{{end}}

};
{{end}}`

var ServerStruct = `package bean



// model
{{range $index,$A := .All }}
type {{$A.Name}} struct {
{{range $index,$V := $A.Attributes }}   {{if $V.IsArray}}{{$V.Name}}     []{{$V.Type}}//{{$V.Desc}}
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
