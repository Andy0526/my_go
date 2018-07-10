package main

import (
	"flag"
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages.")

func dirEnts(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirEnts(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nFiles, nBytes int64) {
	fmt.Printf("%d files %.1f GB\n", nFiles, float64(nBytes)/1e9)
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nFiles, nBytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nFiles++
			nBytes += size
		case <-tick:
			printDiskUsage(nFiles, nBytes)
		}
	}
	printDiskUsage(nFiles, nBytes)

}
