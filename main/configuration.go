package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateDirectory(path string) {
	os.Mkdir(path, os.FileMode(0522))
}

func CreateFile() {

}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func GetListOfFiles(path string) []string {
	var files []string

	err := filepath.Walk(path, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	return files
}
