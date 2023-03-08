package parserFile

import (
	"ginkgo_gen/model"
	"ginkgo_gen/parser/markdown"
	"strings"
)

type Parse interface {
	ParseFile() *model.Content
}

func ParseFileToContent(fileName string) *model.Content {
	if strings.Contains(fileName, ".md") {
		m := markdown.FileMd{
			PathFile: fileName,
		}
		return m.ParseFile()
	}
	return &model.Content{}
}
