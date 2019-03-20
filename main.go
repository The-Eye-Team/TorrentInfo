package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/dustin/go-humanize"
	"github.com/labstack/gommon/color"
)

var stats = struct {
	Size    int
	Index   int
	NbFiles int
}{}

var checkPre = color.Yellow("[") + color.Green("✓") + color.Yellow("]")
var tildPre = color.Yellow("[") + color.Green("~") + color.Yellow("]")
var crossPre = color.Yellow("[") + color.Red("✗") + color.Yellow("]")

func main() {
	var worker sync.WaitGroup
	var count int
	var err error

	stats.Index = 1

	// Parse arguments
	parseArgs(os.Args)

	// Check if input folder exist
	if _, err := os.Stat(arguments.Input); os.IsNotExist(err) {
		fmt.Println(crossPre +
			color.Yellow(" [") +
			color.Red(arguments.Input) +
			color.Yellow("] ") +
			color.Red("Invalid input folder."))
	}

	fmt.Println(color.Green("Collecting Files..."))
	var files []string
	err = filepath.Walk(arguments.Input, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	})

	stats.NbFiles = len(files)
	for _, path := range files {
		worker.Add(1)
		count++
		go calculateSize(path, &worker)
		if count == arguments.Concurrency {
			worker.Wait()
			count = 0
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	worker.Wait()

	fmt.Println(checkPre +
		color.Yellow(" [") +
		color.Green(stats.Index-1) +
		color.Yellow("/") +
		color.Green(stats.NbFiles) +
		color.Yellow("] ") +
		color.Green(" Total size: ") +
		color.Yellow(humanize.Bytes(uint64(stats.Size))))
}
