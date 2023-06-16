package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexflint/go-arg"

	gologging "github.com/sigmonsays/go-logging"
)

type Options struct {
	Verbose  bool     `arg:"-v, --verbose" help:"be verbose"`
	DirsOnly bool     `arg:"-D, --dirs-only" help:"match only directories"`
	NoCase   bool     `arg:"-i, --no-case" help:"case insensitive match"`
	Dir      string   `arg:"-d, --dir" help:"directory to look in"`
	LogLevel string   `arg:"-l, --loglevel" help:"log level"`
	Args     []string `arg:"positional"`
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Debugf("getwd %s", err)
	}

	opts := &Options{
		Verbose:  false,
		DirsOnly: false,
		Dir:      cwd,
		LogLevel: "INFO",
	}
	arg.MustParse(opts)

	gologging.SetLogLevel(opts.LogLevel)

	err = run(opts)
	if err != nil {
		log.Debugf("run %s", err)
	}

}

func run(opts *Options) error {
	paths, err := GetPaths(opts.Dir, opts)
	if err != nil {
		return err
	}
	log.Tracef("loaded %d paths", len(paths))

	var (
		offset        int
		compare_value string
		matches       []string
	)

	log.Tracef("args %v", opts.Args)
	for _, arg := range opts.Args {
		offset = findCommonPrefixOffset(paths)
		log.Tracef("process arg %s offset:%d num_paths:%d",
			arg, offset, len(paths))

		matches = make([]string, 0)
		for _, path := range paths {

			compare_value = path[offset:]

			if strings.Contains(compare_value, arg) {
				matches = append(matches, path)
			}
		}

		if log.IsTrace() {
			for _, m := range matches {
				log.Tracef("MATCH %s", m)
			}
		}

		// print final match and exit
		if len(matches) == 1 {
			fmt.Printf("%s\n", matches[0])
			return nil
		}

		paths = matches
	}

	// print all matches
	for _, m := range matches {
		fmt.Printf("%s\n", m)
	}
	return nil
}

// find the common prefix in matches
func findCommonPrefixOffset(matches []string) int {
	offset := 0
	var v byte
	for {
		v = 0
		for _, m := range matches {

			if v == 0 {
				if offset >= len(m) {
					return offset
				}
				v = m[offset]
				continue
			}
			if offset > len(m) {
				return offset
			}
			if v != m[offset] {
				return offset
			}
		}
		offset += 1
	}

	return offset
}

func GetPaths(path string, opts *Options) ([]string, error) {
	paths := make([]string, 0)
	walkfn := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Debugf("walk %s: %s", path, err)
		}
		if opts.DirsOnly && info.IsDir() == false {
			return nil
		}
		paths = append(paths, path)
		return nil
	}

	err := filepath.Walk(path, walkfn)
	if err != nil {
		log.Debugf("Walk %s", err)
	}
	return paths, nil
}
