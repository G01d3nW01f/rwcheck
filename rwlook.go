package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	red    = "\033[31m"
	blue   = "\033[34m"
	yellow = "\033[33m"
	reset  = "\033[0m"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run filelist.go <directory>")
		os.Exit(1)
	}

	root := os.Args[1]

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Ignore access denied errors
			if os.IsPermission(err) {
				return nil
			}
			return err
		}
		if !info.IsDir() && info.Mode().IsRegular() {
			perm := info.Mode().Perm()
			switch {
			case perm&0600 == 0600:
				fmt.Printf("%s%s%s [readable/writable]\n", red, path, reset)
			case perm&0400 == 0400:
				fmt.Printf("%s%s%s [readable]\n", blue, path, reset)
			case perm&0200 == 0200:
				fmt.Printf("%s%s%s [writable]\n", yellow, path, reset)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
