package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var arguments = struct {
	Input       string
	Concurrency int
}{}

func parseArgs(args []string) {
	// Create new parser object
	parser := argparse.NewParser("TorrentInfo", "Get infos of a folder full of .torrent")

	// Create flags
	input := parser.String("i", "input", &argparse.Options{
		Required: true,
		Help:     "Input directory"})

	concurrency := parser.Int("c", "concurrency", &argparse.Options{
		Required: false,
		Help:     "Concurrency",
		Default:  4})

	// Parse input
	err := parser.Parse(args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	// Fill arguments structure
	arguments.Input = *input
	arguments.Concurrency = *concurrency
}
