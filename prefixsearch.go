package main

import (
	"fmt"
	"strings"
)

func runPrefixSearch(opts *Options) error {
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
		if opts.NoCase {
			arg = strings.ToLower(arg)
		}
		offset = findCommonPrefixOffset(paths)
		log.Tracef("process arg %s offset:%d num_paths:%d",
			arg, offset, len(paths))

		matches = make([]string, 0)
		for _, path := range paths {
			if opts.NoCase {
				compare_value = strings.ToLower(path[offset:])
			} else {
				compare_value = path[offset:]
			}

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
