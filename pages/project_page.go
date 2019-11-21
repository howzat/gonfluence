package pages

import (
	"fmt"
	"gg.gov.revenue.gonfluence/files"
	"html/template"
	"strings"
)

type ProjectPage struct {
	Files       []*files.MarkdownFile
	ProjectName string
}

func NewProjectPage(t *template.Template, rawHtml []byte) template.HTML {

	var buf strings.Builder
	err := t.Execute(&buf, template.HTML(rawHtml))

	if err != nil {
		panic(fmt.Errorf("an error occured processing the project page template %w", err))
	}

	return template.HTML(buf.String())
}
