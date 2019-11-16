package pages

import (
	"gg.gov.revenue.gonfluence/files"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"html/template"
	"testing"
)

type ProjectsPageTestSuite struct {
	suite.Suite
}

func TestProjectsPageTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectsPageTestSuite))
}

func (suite ProjectsPageTestSuite) TestProjectsPageCanGroupBy() {

	var fs []*files.MarkdownFile
	fs = append(fs, &files.MarkdownFile{
		ProjectPath:  "theProjectPath",
		AbsolutePath: "theAbsolutePath",
		ProjectName:  "theProjectName",
		Filename:     "theFile",
	})

	t := template.Must(template.New("tmpl").Parse("{{range .Projects}}{{.}}{{end}}"))

	f := NewProjectsPage(t, func() []*files.MarkdownFile { return fs })

	assert.Equal(suite.T(), template.HTML("{theProjectName [{theAbsolutePath theProjectPath theProjectName theFile}]}"), f)
}
