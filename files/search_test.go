package files_test

import (
	"gg.gov.revenue.gonfluence/files"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SearchTestSuite struct {
	suite.Suite
	fs afero.Fs
}

func (suite *SearchTestSuite) SetupTest() {
	var appFS afero.Fs = afero.NewMemMapFs()
	appFS.MkdirAll("src", 0755)
	suite.fs = appFS
}

func TestFileSearchTestSuite(t *testing.T) {
	suite.Run(t, new(SearchTestSuite))
}

func (suite *SearchTestSuite) TestSearch() {

	suite.fs.MkdirAll("src/a-directory", 0755)
	afero.WriteFile(suite.fs, "src/root-file.md", []byte("root file contents"), 0644)
	afero.WriteFile(suite.fs, "src/a-directory/a-file.md", []byte("a file contents"), 0644)

	var result, err = files.Search(".", suite.fs)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(result.Locations))
}

func (suite *SearchTestSuite) TestBadBasePath() {
	var _, err = files.Search("idontexist", suite.fs)

	if assert.NotNil(suite.T(), err) {
		assert.Equal(suite.T(), "The basePath [idontexist] does not exits", err.Error())
	}
}


func (suite *SearchTestSuite) TestDotFilesAreIgnored() {

	suite.fs.MkdirAll("src/.hidden-directory", 0755)
	afero.WriteFile(suite.fs, "src/.hidden-directory/hidden-file.md", []byte("this file should not be found"), 0644)

	var result, _ = files.Search(".", suite.fs)
	assert.Equal(suite.T(), 0, len(result.Locations), result.Locations)
}

func (suite *SearchTestSuite) TestIsWithinHiddenDirectory() {

	assert.True(suite.T(), files.IsWithinHiddenDir(".src/hidden-file.md"))
	assert.True(suite.T(), files.IsWithinHiddenDir("src/.hidden-directory/hidden-file.md"))
	assert.False(suite.T(), files.IsWithinHiddenDir("src/not-a-hidden-file.md"))
}