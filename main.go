package main

import "github.com/alexflint/go-arg"

type Options struct {
	Verbose  bool
	LogLevel string
}

func main() {

	opts := &Options{
		Verbose: false,
	}
	arg.MustParse(&opts)

}
