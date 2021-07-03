package configuration

import (
	"fmt"
	"os"
	"path/filepath"
)

//Create a directory at the given path.
func CreateDirectory(path string) {
	os.Mkdir(path, os.FileMode(0522))
}

//Create a file at the given path.
func CreateFile() {
	
}

//Walk through the files and append to the array.
func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error on a file")
		} else {
			*files = append(*files, path)
		}
		return nil
	}
}

//Returns a complete list of files at the path specified.
func GetListOfFiles(path string) []string {
	var files []string

	err := filepath.Walk(path, visit(&files))
	if err != nil {
		panic(err)
	}

	return files
}
