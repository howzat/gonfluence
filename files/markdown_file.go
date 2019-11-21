package files

import (
	"fmt"
	"gg.gov.revenue.gonfluence/filtering"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FindMarkdownsFn = func() []*MarkdownFile

type MarkdownFile struct {
	AbsolutePath string
	ProjectPath  string
	ProjectName  string
	Filename     string
}

func FindMarkdownFile(path string, files []*MarkdownFile) *MarkdownFile {
	for _, mdFile := range files {
		if strings.Contains(mdFile.AbsolutePath, path) {
			return mdFile
		}
	}

	return nil
}

func (f *MarkdownFile) Read() []byte {

	s := f.AbsolutePath

	log.Printf("opening file %v", s)
	return ReadFile(s)
}

func NewMarkdownFile(path string, baseDirectory string) (*MarkdownFile, error) {

	absolutePath, withoutBaseDir, e := definePaths(path, baseDirectory)

	if e != nil {
		return nil, e
	}

	dir, file := filepath.Split(withoutBaseDir)

	name := filtering.FirstMatching(
		strings.Split(dir, string(os.PathSeparator)),
		func(elem string) bool {
			return elem != ""
		})

	if dir == "/" {
		dir = ""
	}

	filename := strings.Replace(file, "/", "", 0)

	page := &MarkdownFile{
		AbsolutePath: absolutePath,
		ProjectName:  name,
		ProjectPath:  dir,
		Filename:     filename,
	}

	return page, nil
}

func definePaths(path string, baseDirectory string) (string, string, error) {
	absolutePath, _ := filepath.Abs(path)
	withoutBaseDir, err := pathRelativeToBaseDirectory(absolutePath, baseDirectory)
	if err != nil {
		return "", "", err
	} else if filepath.Ext(withoutBaseDir) != ".md" {
		return "", "", fmt.Errorf("the path provided does not point to a .md file '%v'", withoutBaseDir)
	}
	return absolutePath, withoutBaseDir, nil
}

func pathRelativeToBaseDirectory(path string, baseDir string) (string, error) {
	if !strings.Contains(path, baseDir) {
		return "", fmt.Errorf("the path provided is not contained within the base directory '%v' '%v'", path, baseDir)
	}

	withoutBaseDir := strings.Replace(path, baseDir, "", 1)
	return withoutBaseDir, nil
}
