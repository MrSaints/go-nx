package main

import (
	"github.com/mrsaints/go-nx/nx"
	"log"
)

func main() {
	f, err := nx.NewFile("../../data/Character.nx")
	if err != nil {
		log.Fatalln(err)
	}

	root, err := f.Root()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v", root)
}
