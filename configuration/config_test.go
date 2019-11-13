package configuration

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
	fs afero.Fs
}

func newFileWithContents(contents []byte, fs afero.Fs) string {
	filename := "config.json"
	err := afero.WriteFile(fs, filename, contents, 0644)
	if err != nil {
		panic("failed to initialise the test...")
	}
	return filename
}

func (suite *ConfigTestSuite) newConfigurationWithContents(contents []byte) Configuration {
	var file, _ = suite.fs.Open(newFileWithContents(contents, suite.fs))
	config := parseConfig(file)
	defer file.Close()
	return config
}

func (suite *ConfigTestSuite) SetupTest() {
	suite.fs = afero.NewMemMapFs()
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (suite *ConfigTestSuite) TestParseExclusions() {

	config := suite.newConfigurationWithContents([]byte("{ \"exclusions\" : [\"a\", \"b\"]}"))
	assert.Equal(suite.T(), []string{"a", "b"}, config.Exclusions, config)
}

func (suite *ConfigTestSuite) TestParseBaseDir() {

	config := suite.newConfigurationWithContents([]byte("{ \"baseDir\" : \"mydir\"}}"))
	assert.Equal(suite.T(), "mydir", config.BaseDir, config.BaseDir)
}
