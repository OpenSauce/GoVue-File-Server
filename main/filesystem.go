package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

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

	fmt.Printf("%s of %s disk space used (%0.2f%%)\n",
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
	output := ""
	if CanRunCommands() {
		out, err := exec.Command("ls").Output()
		if err != nil {
			return "", err
		}
		fmt.Println("Command Successfully Executed")
		output := string(out[:])
		fmt.Println(output)
	}
	return output, nil
}

func CanRunCommands() bool {
	if runtime.GOOS == "windows" {
		return false
	} else {
		return true
	}
}

func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return name
}
