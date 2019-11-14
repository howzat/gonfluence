package pages

import (
	"gg.gov.revenue.gonfluence/files"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"html/template"
	"strconv"
	"testing"
)

type ProjectPageTestSuite struct {
	suite.Suite
}

func TestProjectPageTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectPageTestSuite))
}

func (suite ProjectPageTestSuite) TestProjectPageListsAllFoundFiles() {

	var fs []*files.ProjectMarkdownFile
	for i := 0; i < 3; i++ {
		fs = append(fs, &files.ProjectMarkdownFile{
			Filename: "markdown-" + strconv.Itoa(i),
			ProjectName: "project-1",
		})
	}

	t := template.Must(template.New("tmpl").Parse("{{.ProjectName}},{{range .Files}}{{.Filename}},{{end}}"))
	f := NewProjectPage(t, ProjectPage{fs, "name"})
	assert.Equal(suite.T(), template.HTML("name,markdown-0,markdown-1,markdown-2,"), f)
}
