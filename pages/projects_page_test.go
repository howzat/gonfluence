package pages

import (
	"gg.gov.revenue.gonfluence/files"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"html/template"
	"strconv"
	"testing"
)

type ProjectsPageTestSuite struct {
	suite.Suite
}

func TestProjectsPageTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectsPageTestSuite))
}

func (suite ProjectsPageTestSuite) TestProjectsPageCanGroupBy() {

	var fs []*files.ProjectMarkdownFile
	for i := 0; i < 3; i++ {
		fs = append(fs, &files.ProjectMarkdownFile{
			ProjectName: "project-" + strconv.Itoa(i),
		})
	}

	t := template.Must(template.New("tmpl").Parse("{{range .}}{{.}},{{end}}"))

	f := NewProjectPage(t, func() []*files.ProjectMarkdownFile { return fs })

	assert.Equal(suite.T(), template.HTML("project-0,project-1,project-2,"), f)
}
