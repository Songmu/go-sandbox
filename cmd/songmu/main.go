package main

import (
	"flag"
	"log"
	"os"

	"github.com/Songmu/ghselfupdate"
)

const version = "0.0.3"

const (
	exitOK = iota
	exitErr
)

func main() {
	os.Exit(run())
}

func run() int {
	fs := flag.NewFlagSet("songmu", flag.ContinueOnError)
	var selfUpdate = fs.Bool("self-update", false, "self update")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			return exitOK
		}
		return exitErr
	}
	if *selfUpdate {
		err := ghselfupdate.Do(version)
		if err != nil {
			log.Println("Binary update failed:", err)
			return exitErr
		}
		return exitOK
	}
	log.Println("Hello! songmu. version:", version)
	return exitOK
}
