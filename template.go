package main

import (
	"ginkgo_gen/model"
	"html/template"
	"strings"
)

var funcMap = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},
	"join": func(params []model.Field) string {
		res := ``
		for k, v := range params {
			param := ""
			switch v.Type {
			case model.Normal:
				param = strings.ToLower(string(v.Name[0][0])) + v.Name[0][1:]
			case model.Pointer:
				param = `&` + strings.ToLower(string(v.Name[0][0])) + v.Name[0][1:]
			case model.Array:
				param = "[]" + v.TypeName + "{" + strings.ToLower(string(v.Name[0][0])) + v.Name[0][1:] + "}"
			}
			if k == len(params)-1 {
				res += "With" + v.Name[0] + "(" + param + ")"
			} else {
				res += "With" + v.Name[0] + "(" + param + ").\n"
			}
		}
		return res
	},
	"initParam": func(params []model.Field) string {
		res := ""
		for _, v := range params {
			line := strings.ToLower(string(v.Name[0][0])) + v.Name[0][1:] + ":=\"\"" + "\n"
			res += line
		}
		return res
	},
}

var ginkgoTemp = `
package {{.Pkg}}

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	{{.Pkg}}_client "insight.io/api/sdk/v1alpha1/client/{{.Pkg}}"
	"insight.io/test/tools"

	"k8s.io/klog/v2"
)
//go:generate ginkgo --focus={{.Describe}}
var _ = Describe("{{.Describe}}", func() {
	client := tools.Client{Client: tools.NewClient(nil, tools.Cfg)}
	
	//TODO: var param
	{{ initParam .Fs}}

	defaultParams :=func() *{{.Pkg}}_client.{{.Name}}Params{
		return {{.Pkg}}_client.New{{.Name}}Params().
		{{ join .Fs}}
	}

{{range $idx,$val := .Contexts}}
	Context("{{inc $idx}}. {{$val.Context}}", func() {	
		params := defaultParams()
		resp, err := client.Client.{{$.Service}}.{{$.Name}}(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("{{$.Name}} error:", err)
		}
		It("{{$val.It}}", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})		
{{end}}

})

`
