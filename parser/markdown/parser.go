package markdown

import (
	"os"
	"strings"

	"ginkgo_gen/model"

	inf "github.com/fzdwx/infinite"
	log "github.com/sirupsen/logrus"

	"github.com/fzdwx/infinite/components/selection/singleselect"
)

type FileMd struct {
	PathFile string
	CaseName string // 功能
	Describe string // 具体case 的描述

	headerIdx int
}

func (p *FileMd) ParseFile() *model.Content {
	f, err := os.ReadFile(p.PathFile)
	if err != nil {
		log.Fatalln("Read File error:", err)
	}
	rows := strings.Split(string(f), "\n")

	p.selectCase(rows)

	idx := p.headerIdx
	var content model.Content
	top, bottom := -1, -1
	for i := 1; i < len(rows); i++ {
		row := strings.Split(strings.ReplaceAll(rows[i], " ", ""), "|")
		if strings.Contains(row[idx], p.Describe) {
			content.Describe = row[idx]
			top = i
		} else if len(row[idx]) != 0 {
			if top != -1 {
				bottom = i
				break
			}
		}

	}

	if top != -1 && bottom == -1 {
		bottom = len(rows)
	}

	for i := top; i < bottom; i++ {
		row := strings.Split(strings.ReplaceAll(rows[i], " ", ""), "|")
		if len(row) > idx+1 && len(row[idx+1]) > 0 {
			ctx := row[idx+1]
			ctxIt := strings.Split(ctx, ",")
			if strings.Contains(ctx, "，") {
				ctxIt = strings.Split(ctx, "，")
			}
			content.Contexts = append(content.Contexts, model.ContextIt{
				Context: strings.Join(ctxIt[:len(ctxIt)-1], ", "),
				It:      ctxIt[len(ctxIt)-1],
			})
		}
	}

	return &content
}

func (p *FileMd) selectCase(rows []string) {

	headers := strings.Split(strings.ReplaceAll(rows[0], " ", ""), "|")

	var cases []string
	for _, v := range headers {
		if len(v) != 0 {
			cases = append(cases, v)
		}
	}
	caseIdx, err := inf.NewSingleSelect(cases, singleselect.WithPrompt("Please select header"), singleselect.WithDisableFilter()).Display()
	for err != nil {
		log.Fatalln("Please use 'tab' finish selection")
	}

	for k, v := range headers {
		if v == cases[caseIdx] {
			p.headerIdx = k
			break
		}
	}
	p.CaseName = headers[p.headerIdx]
	var desc []string
	for i := 1; i < len(rows); i++ {
		row := strings.Split(strings.ReplaceAll(rows[i], " ", ""), "|")
		if len(row) > p.headerIdx && len(row[p.headerIdx]) != 0 {
			desc = append(desc, row[p.headerIdx])
		}
	}
	descIdx, err := inf.NewSingleSelect(desc, singleselect.WithPrompt("Please select your Describe"), singleselect.WithDisableFilter()).Display()
	for err != nil {
		log.Fatalln("Please use 'tab' finish selection")
	}
	p.Describe = desc[descIdx]
}
