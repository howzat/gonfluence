package files

import (
	"fmt"
	"gg.gov.revenue.gonfluence/configuration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

const BaseDirectory = "/Users/boo/bar/bing"

type MarkdownFileTestSuite struct {
	suite.Suite
	BaseDirectory string
	config        configuration.Configuration
}

func (suite *MarkdownFileTestSuite) SetupTest() {}

func TestFileSearchTestSuite(t *testing.T) {
	suite.Run(t, new(MarkdownFileTestSuite))
}

func (suite *MarkdownFileTestSuite) TestFailWithoutAFile() {

	_, err := NewMarkdownFile(BaseDirectory+"/a/b/c/", BaseDirectory)
	if assert.Error(suite.T(), err) {
		assert.Equal(suite.T(), fmt.Errorf("the path provided does not point to a .md file '/a/b/c'"), err)
	}
}

func (suite *MarkdownFileTestSuite) TestErrorIfNotWithinBaseDir() {

	_, err := NewMarkdownFile("/a/b/c/foo.md", BaseDirectory)
	if assert.Error(suite.T(), err) {
		assert.Equal(suite.T(), fmt.Errorf("the path provided is not contained within the base directory '/a/b/c/foo.md' '/Users/boo/bar/bing'"), err)
	}
}
