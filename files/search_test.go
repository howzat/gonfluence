package files_test

import (
	"gg.gov.revenue.gonfluence/configuration"
	"gg.gov.revenue.gonfluence/files"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SearchTestSuite struct {
	suite.Suite
	fs     afero.Fs
	config configuration.Configuration
}

func (suite *SearchTestSuite) SetupTest() {

	var appFS = afero.NewMemMapFs()
	_ = appFS.MkdirAll("src", 0755)
	suite.fs = appFS
	suite.config = configuration.Configuration{[]string{}, true, "."}
}

func TestFileSearchTestSuite(t *testing.T) {
	suite.Run(t, new(SearchTestSuite))
}

func (suite *SearchTestSuite) TestSearch() {

	_ = suite.fs.MkdirAll("src/a-directory", 0755)
	_ = afero.WriteFile(suite.fs, "src/root-file.md", []byte("root file contents"), 0644)
	_ = afero.WriteFile(suite.fs, "src/a-directory/a-file.md", []byte("a file contents"), 0644)

	var result, err = files.Search(suite.config, suite.fs)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(result.Locations))
}

func (suite *SearchTestSuite) TestBadBasePath() {

	config := configuration.Configuration{BaseDir: "idontexist"}
	var _, err = files.Search(config, suite.fs)

	if assert.NotNil(suite.T(), err) {
		assert.Equal(suite.T(), "The basePath [idontexist] does not exits", err.Error())
	}
}

func (suite *SearchTestSuite) TestDotFilesAreIgnored() {

	_ = suite.fs.MkdirAll("src/.hidden-directory", 0755)
	_ = afero.WriteFile(suite.fs, "src/.hidden-directory/hidden-file.md", []byte("this file should not be found"), 0644)

	var result, _ = files.Search(suite.config, suite.fs)
	assert.Equal(suite.T(), 0, len(result.Locations), result.Locations)
}

func (suite *SearchTestSuite) TestIsWithinHiddenDirectory() {

	assert.True(suite.T(), files.IsWithinHiddenDir(".src/hidden-file.md"))
	assert.True(suite.T(), files.IsWithinHiddenDir("src/.hidden-directory/hidden-file.md"))
	assert.False(suite.T(), files.IsWithinHiddenDir("src/not-a-hidden-file.md"))
}

func (suite *SearchTestSuite) TestIncludes() {

	_ = suite.fs.MkdirAll("src/.hidden-directory", 0755)
	_ = afero.WriteFile(suite.fs, "src/.hidden-directory/hidden-file.md", []byte("this file should not be found"), 0644)

	assert.True(suite.T(), files.IsWithinHiddenDir(".src/hidden-file.md"))
	assert.True(suite.T(), files.IsWithinHiddenDir("src/.hidden-directory/hidden-file.md"))
	assert.False(suite.T(), files.IsWithinHiddenDir("src/not-a-hidden-file.md"))
}

func (suite *SearchTestSuite) TestPathContainsDir() {

	excludedDirs := []string{"foo", "bar"}
	assert.False(suite.T(), files.PathContainsDir("/a/b/c/file.md", excludedDirs))
	assert.True(suite.T(), files.PathContainsDir("/foo/b/c/file.md", excludedDirs))
	assert.True(suite.T(), files.PathContainsDir("/a/foo/c/file.md", excludedDirs))
	assert.True(suite.T(), files.PathContainsDir("/a/b/foo/file.md", excludedDirs))
	assert.True(suite.T(), files.PathContainsDir("/bar/b/c/file.md", excludedDirs))
	assert.True(suite.T(), files.PathContainsDir("/a/bar/c/file.md", excludedDirs))
	assert.True(suite.T(), files.PathContainsDir("/a/b/bar/file.md", excludedDirs))
}

func (suite *SearchTestSuite) TestExcludeDirs() {

	_ = suite.fs.MkdirAll("src/foo/bar/baz/qux", 0755)
	_ = afero.WriteFile(suite.fs, "src/foo/bar/baz/qux/hidden-file.md", []byte("this file should not be found"), 0644)

	assert.True(suite.T(), files.IsWithinHiddenDir(".src/hidden-file.md"))
}
