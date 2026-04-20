package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry == nil {
			continue
		}
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("err %v", err)
		return nil
	}

	fileInfos := make([]os.FileInfo, len(entries))
	fmt.Printf("\n entries %v\n", entries)
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		fileInfos = append(fileInfos, info)
	}

	return fileInfos
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB \n", nfiles, nbytes)
}

func main() {
	flag.Parse()
	fileSizes := make(chan int64)

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fmt.Printf("%v\n", roots)
	var wg sync.WaitGroup

	for _, root := range roots {
		wg.Add(1)
		go func(rt string) {
			defer wg.Done()
			walkDir(rt, fileSizes)
		}(root)

	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for fileSize := range fileSizes {
		nfiles++
		nbytes += fileSize
	}

	printDiskUsage(nfiles, nbytes)
}
