package filesystem

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetHostname(t *testing.T) {
	got := GetHostname()
	want := "PC"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHDDStats(t *testing.T) {
	t.Helper()
	GetHDDStats("/")
}

func TestCommandExec(t *testing.T) {
	if CanRunCommands() {
		t.Helper()

		got, err := ExecuteCommand("whoami")
		want := "pi"
		log.Printf("%s", got)
		if err != nil {
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		} else {
			t.Failed()
		}
	}
}

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

func Test_visit(t *testing.T) {
	type args struct {
		files *[]string
	}
	tests := []struct {
		name string
		args args
		want filepath.WalkFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := visit(tt.args.files); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("visit() = %v, want %v", got, tt.want)
			}
		})
	}
}
