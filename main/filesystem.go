package main

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
	"github.com/minio/minio/pkg/disk"
)

func printUsage(path string) float64 {
	di, err := disk.GetInfo(path)
	if err != nil {
		
	}
	percentage := (float64(di.Total-di.Free)/float64(di.Total))*100
	fmt.Printf("%s of %s disk space used (%0.2f%%)\n", 
		humanize.Bytes(di.Total-di.Free), 
		humanize.Bytes(di.Total), 
		percentage,
	)
	return 100 - percentage
}