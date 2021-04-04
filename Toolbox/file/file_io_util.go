package file

import (
	"io/ioutil"
	"path/filepath"
)

func ReadStrFromFile(dataPath string) string {
	abs, _ := filepath.Abs(dataPath)
	file, e := ioutil.ReadFile(abs)
	if e != nil {
	}
	return string(file)
}
