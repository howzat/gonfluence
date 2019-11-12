package main

import (
	"fmt"
	"gg.gov.revenue.gonfluence/configuration"
	"gg.gov.revenue.gonfluence/files"
	"gg.gov.revenue.gonfluence/pages"
	"github.com/gorilla/mux"
	"github.com/shurcooL/github_flavored_markdown/gfmstyle"
	"github.com/spf13/afero"
	"html/template"
	"net/http"
)

type MyHandler = func(http.ResponseWriter, *http.Request)

func serveContent(content []byte, t *template.Template) MyHandler {

	return func(w http.ResponseWriter, r *http.Request) {
		_ = t.Execute(w, template.HTML(string(content)))
	}
}

func servePage(t template.HTML) MyHandler {

	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, string(t))
	}
}

const ProjectPageTemplate = "project-page-template.html"

func main() {

	config := configuration.ReadConfiguration("gonfluence.json")
	//searchResult := findMarkdownFiles(config)

	t, _ := template.ParseFiles("site/" + ProjectPageTemplate)
	projectPage := pages.NewProjectPage(t, func() []*files.ProjectMarkdownFile {
		return findMarkdownFiles(config)
	})

	//var output = github_flavored_markdown.Markdown(projectFiles[0].Read())
	//t, err := template.ParseFiles("site/body.html")
	//
	router := mux.NewRouter()
	router.PathPrefix("/site/").Handler(http.StripPrefix("/site/", http.FileServer(http.Dir("site"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(gfmstyle.Assets)))
	//router.Path("/gonfluence").Handler(http.RedirectHandler("/gonfluence/projects", 302))
	router.HandleFunc("/gonfluence/projects", servePage(projectPage)).Methods("GET")

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	server.ListenAndServe()
}

func findMarkdownFiles(config configuration.Configuration) []*files.ProjectMarkdownFile {

	result, err := files.Search(config, afero.NewReadOnlyFs(afero.NewOsFs()))
	if err != nil {
		panic(fmt.Errorf("failed to initialise Gonfluence [%w]", err))
	}

	var projectFiles []*files.ProjectMarkdownFile

	for _, f := range result.Locations {
		file, e2 := files.NewProjectMarkdownFile(f, config.BaseDir)
		if e2 != nil {
			panic(fmt.Errorf("failed to create a representation of the discovered markdown file '%s' [%w]", f, e2))
		}
		projectFiles = append(projectFiles, file)
	}

	return projectFiles
}
