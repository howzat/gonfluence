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
	"log"
	"net/http"
)

type HttpHandler = func(http.ResponseWriter, *http.Request)


const ProjectsPageTemplate = "projects-page-template.html"
const ProjectPageTemplate = "project-page-template.html"

func main() {

	config := configuration.ReadConfiguration("gonfluence.json")

	filesCache := findMarkdownFiles(config.BaseDir, config.Exclusions)
	files := func() []*files.ProjectMarkdownFile {return filesCache}
	//searchResult := findMarkdownFiles(config)

	//var output = github_flavored_markdown.Markdown(projectFiles[0].Read())
	//projectsTemplate, err := template.ParseFiles("site/body.html")
	//
	router := mux.NewRouter()
	router.PathPrefix("/site/").Handler(http.StripPrefix("/site/", http.FileServer(http.Dir("site"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(gfmstyle.Assets)))
	router.Path("/gonfluence").Handler(http.RedirectHandler("/gonfluence/projects", 302))

	router.HandleFunc("/gonfluence/projects", projectsPageHandler(config, files)).
		Methods("GET").
		Name("Projects")

	router.HandleFunc("/gonfluence/projects/{project}", projectPageHandler(config)).
		Methods("GET").
		Name("Project")

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}

func projectPageHandler(config configuration.Configuration) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		templ, _ := template.ParseFiles("site/" + ProjectPageTemplate)
		project := mux.Vars(r)["project"]
		log.Printf("serving files under project: %q\n", project)
		html := pages.NewProjectPage(templ, project, func() []*files.ProjectMarkdownFile {
			return findMarkdownFiles(config.BaseDir + "/" + project, config.Exclusions)
		})

		_, execute := fmt.Fprintf(w, string(html))
		if execute != nil {
			log.Panicln(execute)
		}
	}
}

func projectsPageHandler(config configuration.Configuration, files func() []*files.ProjectMarkdownFile) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		projectsTemplate, _ := template.ParseFiles("site/" + ProjectsPageTemplate)
		html := pages.NewProjectsPage(projectsTemplate, files)

		_, execute := fmt.Fprintf(w, string(html))
		if execute != nil {
			log.Panicln(execute)
		}

	}
}

func findMarkdownFiles(baseDir string, exclusions []string) []*files.ProjectMarkdownFile {

	log.Printf("looking for markdown files in %q with exclusions [%q]\n", baseDir, exclusions)
	result, err := files.Search(baseDir, exclusions, afero.NewReadOnlyFs(afero.NewOsFs()))
	if err != nil {
		panic(fmt.Errorf("failed to initialise Gonfluence [%w]", err))
	}

	var projectFiles []*files.ProjectMarkdownFile

	for _, f := range result.Locations {
		file, e := files.NewProjectMarkdownFile(f, baseDir)
		if e != nil {
			panic(fmt.Errorf("failed to create a representation of the discovered markdown file '%s' [%w]", f, e))
		}
		projectFiles = append(projectFiles, file)
	}

	return projectFiles
}
