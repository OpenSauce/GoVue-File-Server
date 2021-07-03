package filesystem

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/minio/minio/pkg/disk"
)

type HDDStats struct {
	TotalSpace string
	FreeSpace  string
	Percentage float64
}

func GetHDDStats(path string) HDDStats {
	di, err := disk.GetInfo(path)
	if err != nil {

	}
	stats := HDDStats{
		TotalSpace: humanize.Bytes(di.Total),
		FreeSpace:  humanize.Bytes(di.Total - di.Free),
		Percentage: avaliableSpacePercentage(float64(di.Total), float64(di.Total-di.Free)),
	}

	log.Printf("%s of %s disk space used (%0.2f%%)\n",
		stats.FreeSpace,
		stats.TotalSpace,
		stats.Percentage,
	)
	return stats
}

func avaliableSpacePercentage(total float64, free float64) float64 {
	return (float64(free) / float64(total)) * 100
}

func ExecuteCommand(command string) (string, error) {
	if !CanRunCommands() {
		return "", errors.New("command can't be ran on this machine")
	}

	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "", err
	}

	output := string(out[:])
	output = strings.Replace(output, "\n", "\\n", -1)
	return output, nil

}

func CanRunCommands() bool {
	return runtime.GOOS != "windows"
}

func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

//Walk through the files and append to the array.
func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error on a file")
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


//Create a directory at the given path.
func CreateDirectory(path string) {
	os.Mkdir(path, os.FileMode(0522))
}

//Create a file at the given path.
func CreateFile() {
	
}


