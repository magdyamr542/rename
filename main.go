package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func getUsage(appName string) string {
	return fmt.Sprintf("usage %v --replace [pattern] --with [substring] [directory]\nexample: %v --replace main --with index ./", appName, appName)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "error: %v\n", getUsage(os.Args[0]))
		os.Exit(1)
	}

	replace := flag.String("replace", "", "--replace [pattern]. will search for the pattern in the files")
	with := flag.String("with", "", "--with [substring]. the found patterns in the files will be replaced with the substring")
	flag.Parse()

	if len(*replace) == 0 || len(*with) == 0 {
		fmt.Fprintf(os.Stderr, "error: %v\n", getUsage(os.Args[0]))
		os.Exit(1)
	}

	if len(os.Args) < 6 {
		fmt.Fprintf(os.Stderr, "error: no directory specified: %v\n", getUsage(os.Args[0]))
		os.Exit(1)
	}

	dir := os.Args[len(os.Args)-1]
	pattern := *replace
	newPattern := *with

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
		os.Exit(1)
	}
	matches := []fs.FileInfo{}
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), pattern) {
			matches = append(matches, file)

		}
	}

	if len(matches) == 0 {
		fmt.Fprintf(os.Stderr, "error: no matches found for the pattern '%v'\n", pattern)
		os.Exit(1)
	}

	for _, file := range matches {
		if !file.IsDir() && strings.Contains(file.Name(), pattern) {
			newName := strings.Replace(file.Name(), pattern, newPattern, -1)
			oldPath := path.Join(dir, file.Name())
			newPath := path.Join(dir, newName)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: failed to rename %v. %v\n", oldPath, err.Error())
			}
		}
	}

}
