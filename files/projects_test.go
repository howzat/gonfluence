package files

import (
	"gg.gov.revenue.gonfluence/configuration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"sort"
	"testing"
)

type ProjectsTestSuite struct {
	suite.Suite
	BaseDirectory string
	config        configuration.Configuration
}

func (suite *ProjectsTestSuite) SetupTest() {}

func TestFindMarkdownFileTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectsTestSuite))
}

func (suite *ProjectsTestSuite) TestAddingAFileToAProject() {

	m1, _ := NewMarkdownFile("/Users/workspace/project-a/A-README.md", "/Users/workspace")
	m2, _ := NewMarkdownFile("/Users/workspace/project-a/sub-a/SUB-A-README.md", "/Users/workspace")
	m3, _ := NewMarkdownFile("/Users/workspace/project-b/B-README.md", "/Users/workspace")
	m4, _ := NewMarkdownFile("/Users/workspace/project-b/B-README-2.md", "/Users/workspace")

	projects := Projects{}
	projects.AddMarkdown(m1)
	projects.AddMarkdown(m2)
	projects.AddMarkdown(m3)
	projects.AddMarkdown(m4)

	names := projects.Names()
	sort.Strings(names)
	assert.Equal(suite.T(), []string{"project-a", "project-b"}, names)

	projectAFiles := suite.collectFilenames(projects, "project-a")
	assert.Equal(suite.T(), []string{"A-README.md", "SUB-A-README.md"}, projectAFiles)

	projectBFiles := suite.collectFilenames(projects, "project-b")
	assert.Equal(suite.T(), []string{"B-README.md", "B-README-2.md"}, projectBFiles)
}

func (suite *ProjectsTestSuite) collectFilenames(projects Projects, s string) []string {
	var filenames []string
	for _, md := range projects.Projects[s].Files {
		filenames = append(filenames, md.Filename)
	}
	return filenames
}
