package gengo

// 报错行号+3
const goCodeTemplate = `// Auto generated by github.com/davyxu/cellmesh/protogen
// DO NOT EDIT!

package {{.PackageName}}

import (
	"fmt"
	"reflect"	
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	{{if HasJsonCodec $}}_ "github.com/davyxu/cellnet/codec/json"{{end}}
	_ "github.com/davyxu/cellnet/codec/binary"
)

// Make compiler import happy
var (
	_ cellnet.Peer
	_ cellnet.Codec
	_ reflect.Type
	_ fmt.Formatter
)


{{range $a, $enumobj := .Enums}}
type {{.Name}} int32
const (	{{range .Fields}}
	{{$enumobj.Name}}_{{.Name}} {{$enumobj.Name}} = {{TagNumber $enumobj .}} {{end}}
)

var (
{{$enumobj.Name}}MapperValueByName = map[string]int32{ {{range .Fields}}
	"{{.Name}}": {{TagNumber $enumobj .}}, {{end}}
}

{{$enumobj.Name}}MapperNameByValue = map[int32]string{ {{range .Fields}}
	{{TagNumber $enumobj .}}: "{{.Name}}" , {{end}}
}

{{$enumobj.Name}}MapperTrailingCommentByValue = map[int32]string{ {{range .Fields}}
	{{TagNumber $enumobj .}}: "{{.Trailing}}" , {{end}}
}
)

func (self {{$enumobj.Name}}) String() string {
	return {{$enumobj.Name}}MapperNameByValue[int32(self)]
}
{{end}}

{{range .Structs}}
{{ObjectLeadingComment .}}
type {{.Name}} struct{	{{range .Fields}}
	{{GoFieldName .}} {{GoTypeName .}} {{GoStructTag .}}{{FieldTrailingComment .}} {{end}}
}
{{end}}
{{range .Structs}}
func (self *{{.Name}}) String() string { return fmt.Sprintf("%+v",*self) } {{end}}

{{range ServiceGroup $}}
// {{$svcName := .Key}}{{$svcName}}
var ( {{range .Group}}
	Handle_{{ExportSymbolName $svcName}}_{{.Name}} = func(ev cellnet.Event){ panic("'{{.Name}}' not handled") } {{end}}
	Handle_{{ExportSymbolName $svcName}}_Default func(ev cellnet.Event)
)
{{end}}

func GetMessageHandler(svcName string) cellnet.EventCallback {

	switch svcName { {{range ServiceGroup $}}
	case "{{$svcName := .Key}}{{$svcName}}":
		return func(ev cellnet.Event) {
			switch ev.Message().(type) { {{range .Group}}
			case *{{.Name}}:
				Handle_{{ExportSymbolName $svcName}}_{{.Name}}(ev) {{end}}
			default:
				if Handle_{{ExportSymbolName $svcName}}_Default != nil {
					Handle_{{ExportSymbolName $svcName}}_Default(ev)
				}
			}
		} {{end}}
	} 

	return nil
}

func init() {
	{{range .Structs}} {{ if IsMessage . }}
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("{{StructCodec .}}"),
		Type:  reflect.TypeOf((*{{.Name}})(nil)).Elem(),
		ID:    {{StructMsgID .}},
	})
	{{end}} {{end}}
}

`
