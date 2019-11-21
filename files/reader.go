package files

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) []byte {

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	info, err := file.Stat()
	check(err)

	contents := make([]byte, info.Size())

	contents, err = ioutil.ReadFile(filename)
	check(err)

	return contents
}

func check(e error) {
	if e != nil {
		panic(fmt.Errorf("an error occured reading the file %w", e))
	}
}
