package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

func main() {
	filesizes := make(chan <- int64)
	walkDir(".", filesizes)
}

func walkDir(dir string, fileSizes chan <- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir,entry.Name())
			walkDir(subdir, fileSizes)
		}else {
			fmt.Printf("%s\n", entry.Name())
			fileSizes <- entry.Size()
		}
	}
} // dirents returns the entries of directory

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}