package main

import (
	"os"
	"text/template"
)

type (
	TypeInfo struct {
		Name     string
		TypeName string
	}
)

var header string = `// Generated by gen_types. Don't modify manually!

package main

import (
  "fmt"
)
`

var assertTemplate string = `
func assert{{.Name}}(obj Object, msg string) {{.TypeName}} {
  switch c := obj.(type) {
  case {{.TypeName}}:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "{{.Name}}", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}
`

var ensureTemplate string = `
func ensure{{.Name}}(args []Object, index int) {{.TypeName}} {
  switch c := args[index].(type) {
  case {{.TypeName}}:
    return c
  default:
    panic(RT.newArgTypeError(index, "{{.Name}}"))
  }
}
`

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := "types_gen.go"
	f, err := os.Create(filename)
	checkError(err)
	defer f.Close()

	var assert = template.Must(template.New("assert").Parse(assertTemplate))
	var ensure = template.Must(template.New("ensure").Parse(ensureTemplate))
	f.WriteString(header)
	for _, t := range os.Args[1:] {
		typeInfo := TypeInfo{
			Name:     t,
			TypeName: t,
		}
		if t[0] == '*' {
			typeInfo.Name = t[1:]
		}
		assert.Execute(f, typeInfo)
		ensure.Execute(f, typeInfo)
	}
}
