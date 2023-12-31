package main

import (
	"io/fs"
	"path/filepath"
	"sync"
)

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

func GetPathsChan(wg *sync.WaitGroup, paths string, opts *Options, work chan string) (int, error) {
	expected := 0
	walkfn := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Debugf("walk %s: %s", path, err)
		}
		if opts.DirsOnly && info.IsDir() == false {
			return nil
		}
		expected++
		wg.Add(1)
		work <- path
		return nil
	}

	// resolve symlink
	dir := opts.Dir
	res, err := filepath.EvalSymlinks(opts.Dir)
	if err == nil && res != "" {
		log.Tracef("%s is symlink, walking %s", opts.Dir, res)
		dir=res
	}

	err = filepath.Walk(dir, walkfn)
	if err != nil {
		log.Debugf("Walk %s", err)
	}
	return expected, nil
}
