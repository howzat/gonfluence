package pages

import (
	"fmt"
	"gg.gov.revenue.gonfluence/files"
	"gg.gov.revenue.gonfluence/filtering"
	"html/template"
	"strings"
)

func NewProjectsPage(t *template.Template, findMarkdownFiles files.FindMarkdownsFn) template.HTML {

	projects := files.Projects{}
	for _, markdown := range findMarkdownFiles() {
		projects.AddMarkdown(markdown)
	}

	var buf strings.Builder
	err := t.Execute(&buf, projects)

	if err != nil {
		panic(fmt.Errorf("an error occured processing the project page template %w", err))
	}

	return template.HTML(buf.String())
}

func FindProjectDirectories(fn files.FindMarkdownsFn) []string {

	var projects []string
	for _, f := range fn() {
		projects = append(projects, f.ProjectName)
	}

	return filtering.Distinct(projects)
}
