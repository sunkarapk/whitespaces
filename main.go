package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	Version = "0.1.0"
)

type Options struct {
	Version bool `short:"v" long:"version" description:"Show version information"`
	Check   bool `short:"c" long:"check" description:"Check instead of overwriting"`
	Force   bool `short:"f" long:"force" description:"Force whitespace checking all files"`
}

func main() {

	// Options for flags package
	var opts Options

	// Build the parser
	parser := flags.NewParser(&opts, flags.Default)

	// Set usage string
	parser.Usage = "[options] FILES"

	// Parse the arguments
	args, err := parser.Parse()
	handle(err)

	// Print version and exit
	if opts.Version {
		fmt.Println(Version)
		os.Exit(0)
	}

	// If no argument is given
	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	whitespace := regexp.MustCompile("\\s+$")

	for _, v := range args {

		name, err := filepath.Abs(v)
		handle(err)

		info, err := os.Stat(name)
		handle(err)

		if info.IsDir() {
			fmt.Println(name + " seems to be a directory. Skipping...")
		} else {
			data, err := ioutil.ReadFile(name)
			handle(err)

			lines := strings.Split(string(data), "\n")

			for i, line := range lines {
				if opts.Check {
					if whitespace.FindStringIndex(line) != nil {
						check(name, opts)
						break
					}
				} else {
					lines[i] = whitespace.ReplaceAllString(line, "")
				}
			}

			if !opts.Check {
				data := []byte(strings.Join(lines, "\n"))

				ioutil.WriteFile(name, data, info.Mode())
			}
		}
	}
}

func check(name string, opts Options) {
	fmt.Println("Found whitespace in " + name)

	if !opts.Force {
		os.Exit(0)
	}
}

func handle(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
