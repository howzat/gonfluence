package pages

import (
	"fmt"
	"gg.gov.revenue.gonfluence/files"
	"html/template"
	"strings"
)

type ProjectPage struct {
	Files       []*files.ProjectMarkdownFile
	ProjectName string
}

func NewProjectPage(t *template.Template, projectName string, findMarkdownFiles SearchFn) template.HTML {

	var buf strings.Builder

	page := ProjectPage{findMarkdownFiles(), projectName}
	err := t.Execute(&buf, page)

	if err != nil {
		panic(fmt.Errorf("an error occured processing the project page template %w", err))
	}

	return template.HTML(buf.String())
}
