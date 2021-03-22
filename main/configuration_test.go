package main

import (
	"os"
	"testing"
)

func TestGetFileListing(t *testing.T) {

	t.Run("Testing create directory", func(t *testing.T) {
		CreateDirectory("testDir/")

		if _, err := os.Stat("testDir/"); os.IsNotExist(err) {
			t.Fail()
		}

	})

	t.Run("Testing adding files to directory", func(t *testing.T) {

	})

	t.Run("Testing list of files", func(t *testing.T) {

	})

}
