package pages

import (
	"fmt"
	"gg.gov.revenue.gonfluence/files"
	"gg.gov.revenue.gonfluence/filtering"
	"html/template"
	"strings"
)

type ProjectNames struct {
	Names []string
}

type SearchFn = func() []*files.ProjectMarkdownFile

func NewProjectsPage(t *template.Template, findMarkdownFiles SearchFn) template.HTML {

	var projectDirectories = FindProjectDirectories(findMarkdownFiles)
	var buf strings.Builder
	err := t.Execute(&buf, projectDirectories)

	if err != nil {
		panic(fmt.Errorf("an error occured processing the project page template %w", err))
	}

	return template.HTML(buf.String())
}

func FindProjectDirectories(fn SearchFn) []string {

	var projects []string
	for _, f := range fn() {
		projects = append(projects, f.ProjectName)
	}

	return filtering.Distinct(projects)
}
