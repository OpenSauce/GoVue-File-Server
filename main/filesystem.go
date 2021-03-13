package main

import (
	"fmt"
	"os"

	humanize "github.com/dustin/go-humanize"
	"github.com/minio/minio/pkg/disk"
)

type HDDStats struct {
	TotalSpace string
	FreeSpace string
	Percentage float64
}

func GetHDDStats(path string) HDDStats {
	di, err := disk.GetInfo(path)
	if err != nil {
		
	}
	stats := HDDStats{
		TotalSpace: humanize.Bytes(di.Total),
		FreeSpace: humanize.Bytes(di.Total-di.Free),
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
	return (float64(free)/float64(total)) * 100
}

func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return name;
}