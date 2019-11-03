package files

import (
	"errors"
	"fmt"
	"gg.gov.revenue.gonfluence/files/filtering"
	"github.com/spf13/afero"
	"os"
	"path/filepath"
	"strings"

)

type Result struct {
	Locations [] string
}

func isHiddenDirectory(dir string) bool {
	return strings.HasPrefix(dir, ".")
}

func IsWithinHiddenDir(path string) bool {
	dir, _ := filepath.Split(path)
	return filtering.Any(strings.Split(dir, "/"), isHiddenDirectory)
}

func Search(basePath string, fs afero.Fs)  (Result, error) {

	results := make([]string, 0)
	emptyResult := Result{results}

	var exists, _ = afero.DirExists(fs, basePath)
	if !exists {
		return emptyResult, errors.New(fmt.Sprintf("The basePath [%s] does not exits", basePath))
	}

	walker := func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if filepath.Ext(path) == ".md" {
			results = append(results, path)
		}

		return nil
	}

	err := afero.Walk(fs, basePath, walker)
	if err != nil {
		return emptyResult, errors.New(fmt.Sprintf("failed to traverse file system [%w]\n", err))
	}

	visibleFiles := filtering.NotMatching(results, IsWithinHiddenDir)
	return Result{visibleFiles}, nil
}