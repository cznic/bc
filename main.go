package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cznic/httpfs"
	"github.com/cznic/virtual"
)

func main() {
	fs := httpfs.NewFileSystem(assets, time.Now())
	f, err := fs.Open("/bc")
	if err != nil {
		panic(err)
	}

	var bin virtual.Binary
	if _, err := bin.ReadFrom(f); err != nil {
		panic(err)
	}

	exitCode, err := virtual.Exec(&bin, os.Args, os.Stdin, os.Stdout, os.Stderr, 0, 8<<20, "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if exitCode == 0 {
			exitCode = 1
		}
	}
	os.Exit(exitCode)
}
