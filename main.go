package main

import (
	"bytes"
	"flag"

	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"ginkgo_gen/model"
	gen_parser "ginkgo_gen/parser"

	log "github.com/sirupsen/logrus"
)

var (
	fileName string
	outPath  string
	caseFile string
)

var parserFiles []string

func init() {
	flag.StringVar(&fileName, "f", "", "文件名")
	flag.StringVar(&outPath, "o", "", "输出路径")
	flag.StringVar(&caseFile, "cf", "", "对应的 case 文件路径")
	flag.Parse()
}

func main() {
	getwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	if len(fileName) == 0 {
		files, err := os.ReadDir(getwd)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if strings.Contains(file.Name(), ".go") {
				parserFiles = append(parserFiles, file.Name())
			}
		}
	} else {
		parserFiles = append(parserFiles, fileName)
	}

	log.Info("Will parser:", parserFiles)
	fset := token.NewFileSet()

	// TODO: support parser files
	f, err := parser.ParseFile(fset, parserFiles[0], nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ms := initModel(f)

	if len(caseFile) != 0 {
		m := gen_parser.ParseFileToContent(caseFile)
		ms[0].Contexts = m.Contexts
		ms[0].Describe = m.Describe
	}
	//ast.Print(fset, f)
	genContent(ms)
}

func initModel(f *ast.File) (ms []model.Model) {
	if len(caseFile) != 0 {

	}
	log.Info("initModel:", f.Name.String())
	// 遍历类型
	for _, decl := range f.Decls {
		switch d := decl.(type) {

		case *ast.GenDecl: // 必须为定义类型
			var m model.Model // 自定义模版
			m.Pkg = f.Name.String()
			m.Service = strings.ToUpper(string(m.Pkg[0])) + m.Pkg[1:]
			if d.Doc != nil {
				for _, comment := range d.Doc.List {
					com := strings.TrimSpace(comment.Text[2:])
					if strings.Contains(com, "Describe") {
						m.Describe = strings.Split(com, ":")[1][1:]
					} else {
						s := strings.Split(com, ",")
						if len(s) < 2 {
							s = strings.Split(com, "，")
						}
						m.Contexts = append(m.Contexts, model.ContextIt{
							Context: strings.Join(s[:len(s)-1], ", "),
							It:      s[len(s)-1],
						})
					}
				}

			}

			for _, spec := range d.Specs {
				t, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				m.Name = t.Name.String()
				structType, ok := t.Type.(*ast.StructType)
				// 只接受结构体类型
				if !ok {
					continue
				}
				for _, field := range structType.Fields.List {
					var modelField model.Field
					// 一个 filed 可能包含多个字段名
					for _, name := range field.Names {
						modelField.Name = append(modelField.Name, name.Name)
					}
					if field.Tag != nil {
						modelField.Tag = field.Tag.Value
					}
					//todo:暂时不支持嵌套
					switch filedT := field.Type.(type) {
					case *ast.Ident:
						modelField.Type = model.Normal
						modelField.TypeName = filedT.Name
					case *ast.StarExpr:
						modelField.Type = model.Pointer
						modelField.TypeName = filedT.X.(*ast.Ident).Name
					case *ast.ArrayType:
						modelField.Type = model.Array
						modelField.TypeName = filedT.Elt.(*ast.Ident).Name
					}

					// 加入fields
					m.Fs = append(m.Fs, modelField)
				}
			}

			if len(m.Fs) != 0 {
				ms = append(ms, m)
			}

		}

	}
	return ms
}

func genContent(ms []model.Model) {
	for _, v := range ms {
		parse, err := template.New(v.Name).Funcs(funcMap).Parse(ginkgoTemp)
		if err != nil {
			panic(err)
		}
		buf := bytes.NewBuffer(nil)
		err = parse.Execute(buf, v)
		if err != nil {
			panic(err)
		}
		r1 := strings.ReplaceAll(buf.String(), "&#34;&#34", "\"\"")
		r2 := strings.ReplaceAll(r1, "&amp;", "&")
		source, err := format.Source([]byte(r2))
		if err != nil {
			log.Println("format fail:", err)
			return
		}
		dir := outPath
		if len(outPath) == 0 {
			modelFileDir := strings.Split(fileName, "/")
			dir = filepath.Join(modelFileDir[:len(dir)-1]...)
		}
		f, err := os.OpenFile(dir+"/"+strings.ToLower(v.Name+".go"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			panic(err)
		}
		f.Write(source)
	}
}
