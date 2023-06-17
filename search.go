package main

import (
	"fmt"
	"strings"
	"sync"
)

func runSearch(opts *Options) error {
	paths := make(chan string)
	results := make(chan string)
	control := make(chan int, 1)

	log.Tracef("runSearch %s args:%v", opts.Dir, opts.Args)

	var wg sync.WaitGroup

	// send work
	wg.Add(1)
	go func() {
		expected, err := GetPathsChan(&wg, opts.Dir, opts, paths)
		if err != nil {
			log.Warnf("GetPaths %s: %s", opts.Dir, err)
			wg.Done()
			return
		}
		log.Tracef("GetPathsChan returned expected:%d", expected)
		wg.Done()
	}()

	// start workers
	for w := 0; w < opts.Concurrency; w++ {
		go searchWorker(&wg, w, opts, paths, results)
	}

	// wait for results
	go func() {
		received := 0
	Done:
		for {
			select {
			case <-control:
				break Done

			case res := <-results:
				received++
				fmt.Printf("%s\n", res)
			}
		}
		log.Tracef("receive results exiting")
	}()

	wg.Wait()
	control <- 0
	close(paths)

	return nil
}

func searchWorker(wg *sync.WaitGroup, w int, opts *Options, paths chan string, results chan string) {
	wid := fmt.Sprintf("worker%d", w)
	log.Tracef("%s started", wid)
Dance:
	for {
		select {
		case path, ok := <-paths:
			if !ok {
				break Dance
			}
			if pathMatch(path, opts.Args, opts.CaseSensitive) {
				results <- path
			}
			wg.Done()
		}
	}
	log.Tracef("%s exiting", wid)
}

func pathMatch(path string, args []string, casesensitive bool) bool {
	idx := 0
	offset := 0
	for _, arg := range args {

		if casesensitive {
			idx = strings.Index(strings.ToLower(path[offset:]), strings.ToLower(arg))
		} else {
			idx = strings.Index(path[offset:], arg)
		}

		if idx < 0 {
			return false
		}
		offset = idx
	}
	return true
}
