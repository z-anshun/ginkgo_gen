package markdown

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	m := FileMd{
		PathFile: "../../example/tracing_testcase.md",
		CaseName: "功能点",
		Describe: "service-graph/graph",
	}
	fmt.Println(m.ParseFile())
}
