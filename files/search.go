package files

import (
	"errors"
	"fmt"
	"gg.gov.revenue.gonfluence/filtering"
	"github.com/spf13/afero"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Result struct {
	Locations []string
}

const HiddenFilePrefix = "."
const PathDelimiter = "/"

func IsWithinHiddenDir(path string) bool {

	predicate := func(pathPart string) bool {
		return strings.HasPrefix(pathPart, HiddenFilePrefix)
	}

	dirs, _ := filepath.Split(path)
	return filtering.Any(strings.Split(dirs, PathDelimiter), predicate)
}

func PathContainsDir(path string, excluded []string) bool {

	makePredicate := func(excludedDirs []string) filtering.StringPredicate {
		return func(dir string) bool {
			var ed = excludedDirs
			if "" == dir {
				return false
			} else {
				return filtering.Any(ed, func(excluded string) bool {
					return excluded == dir
				})
			}
		}
	}

	return filtering.Any(strings.Split(path, "/"), makePredicate(excluded))
}

func Search(baseDir string, exclusions []string, fs afero.Fs) (Result, error) {

	results := make([]string, 0)
	emptyResult := Result{}

	var exists, _ = afero.DirExists(fs, baseDir)
	if !exists {
		return emptyResult, errors.New(fmt.Sprintf("The basePath [%s] does not exits", baseDir))
	}

	walker := func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if filepath.Ext(path) == ".md" && !PathContainsDir(path, exclusions) {
			results = append(results, path)
		}

		return nil
	}

	err := afero.Walk(fs, baseDir, walker)
	if err != nil {
		return emptyResult, errors.New(fmt.Sprintf("failed to traverse file system [%w]\n", err))
	}

	visibleFiles := filtering.NotMatching(results, IsWithinHiddenDir)
	log.Printf("visible markdown files %v (excluded files %v)", len(visibleFiles), len(results)-len(visibleFiles))
	return Result{visibleFiles}, nil
}
