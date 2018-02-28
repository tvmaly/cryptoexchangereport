package main

import (
	"fmt"
	libsass "github.com/wellington/go-libsass"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("you must specify a scss file")
		return
	}

	err := os.Chdir(filepath.Dir(os.Args[1]))

	if err != nil {
		log.Fatalf("cannot change to directory where scss is: %s\n", err)
		return
	}

	r, err := os.Open(filepath.Base(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	comp, err := libsass.New(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}

	if err := comp.Run(); err != nil {
		log.Fatal(err)
	}
}
