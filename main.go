package main

import (
	"os"

	"github.com/alexflint/go-arg"

	gologging "github.com/sigmonsays/go-logging"
)

type Options struct {
	Verbose       bool     `arg:"-v, --verbose" help:"be verbose"`
	DirsOnly      bool     `arg:"-D, --dirs-only" help:"match only directories"`
	CaseSensitive bool     `arg:"-C, --case" help:"case sensitive match"`
	PrefixSearch  bool     `arg:"-p, --prefix-search" help:"perform prefix search"`
	Concurrency   int      `arg:"-c, --concurrency" help:"worker concurrency"`
	Dir           string   `arg:"-d, --dir" help:"directory to look in"`
	LogLevel      string   `arg:"-l, --loglevel" help:"log level"`
	Args          []string `arg:"positional"`
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Debugf("getwd %s", err)
	}

	opts := &Options{
		Verbose:     false,
		DirsOnly:    false,
		Dir:         cwd,
		Concurrency: 10, // only for runSearch (not Prefix search)
		LogLevel:    "INFO",
	}
	arg.MustParse(opts)

	gologging.SetLogLevel(opts.LogLevel)

	err = run(opts)
	if err != nil {
		log.Debugf("run %s", err)
	}

}

func run(opts *Options) error {
	if opts.PrefixSearch {
		return runPrefixSearch(opts)
	}
	return runSearch(opts)
}
