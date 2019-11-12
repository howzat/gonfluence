package files

import (
	"fmt"
	"gg.gov.revenue.gonfluence/configuration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

const BaseDirectory = "/Users/boo/bar/bing"

type GonfluencePageTestSuite struct {
	suite.Suite
	BaseDirectory string
	config        configuration.Configuration
}

func (suite *GonfluencePageTestSuite) SetupTest() {}

func TestFileSearchTestSuite(t *testing.T) {
	suite.Run(t, new(GonfluencePageTestSuite))
}

func (suite *GonfluencePageTestSuite) TestFailWithoutAFile() {

	_, err := NewProjectMarkdownFile(BaseDirectory+"/a/b/c/", BaseDirectory)
	if assert.Error(suite.T(), err) {
		assert.Equal(suite.T(), fmt.Errorf("the path provided does not point to a .md file '/a/b/c'"), err)
	}
}

func (suite *GonfluencePageTestSuite) TestErrorIfNotWithinBaseDir() {

	_, err := NewProjectMarkdownFile("/a/b/c/foo.md", BaseDirectory)
	if assert.Error(suite.T(), err) {
		assert.Equal(suite.T(), fmt.Errorf("the path provided is not contained within the base directory '/a/b/c/foo.md' '/Users/boo/bar/bing'"), err)
	}
}
