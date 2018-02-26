package main

import (
	libsass "github.com/wellington/go-libsass"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you must specify a scss file")
		return
	}
	r, err := os.Open(&os.Args[1])
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
