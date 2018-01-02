package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const version = "0.0.2"

const (
	exitOK = iota
	exitErr
)

var slugReg = regexp.MustCompile(`/src/github.com/([^/]+/[^/]+)`)

func main() {
	os.Exit(run())
}

func run() int {
	semv, err := semver.Parse(version)
	if err != nil {
		log.Println(err)
		return exitErr
	}
	fpath := filepath.ToSlash(origFilePath())
	matches := slugReg.FindStringSubmatch(fpath)
	if len(matches) < 2 {
		log.Println("slug not detected")
		return exitErr
	}
	slug := matches[1]

	fs := flag.NewFlagSet("songmu", flag.ContinueOnError)
	var selfUpdate = fs.Bool("self-update", false, "self update")

	err = fs.Parse(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			return exitOK
		}
		return exitErr
	}
	if *selfUpdate {
		latest, err := selfupdate.UpdateSelf(semv, slug)
		if err != nil {
			log.Println("Binary update failed:", err)
			return exitErr
		}
		if latest.Version.Equals(semv) {
			// latest version is the same as current version. It means current binary is up to date.
			log.Println("Current binary is the latest version", version)
		} else {
			log.Println("Successfully updated to version", latest.Version)
		}
		return exitOK
	}
	log.Println("Hello! songmu")
	return exitOK
}

func origFilePath() string {
	_, fname, _, _ := runtime.Caller(1)
	return fname
}
