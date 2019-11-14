package pages

import (
	"fmt"
	"gg.gov.revenue.gonfluence/files"
	"gg.gov.revenue.gonfluence/filtering"
	"html/template"
	"strings"
)

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



type Project struct {
	Name string
	Files []files.ProjectMarkdownFile
}

type Projects struct {
	Projects []Project
}

func FindProjects(fn SearchFn) Projects {

	projectsAndFiles := make(map[string][]files.ProjectMarkdownFile)

	for _, mdFile := range fn() {
		mdFiles, ok := projectsAndFiles[mdFile.ProjectName]
		if !ok {
			projectsAndFiles[mdFile.ProjectName] = append(mdFiles, *mdFile)
		} else {
			markdownFiles := make([]files.ProjectMarkdownFile, 1)
			projectsAndFiles[mdFile.ProjectName] = append(markdownFiles, *mdFile)
		}
	}

	var ppp = make([]Project, 0)

	for k,v := range projectsAndFiles {
		ppp = append(ppp, Project{Name:k, Files:v})
	}


	return Projects{ppp}
}
