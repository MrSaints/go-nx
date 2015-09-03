package main

import (
	"github.com/mrsaints/go-nx"
	"log"
)

func main() {
	nxFile := gonx.New("../../Data/Character.nx")
	root := nxFile.Root()
	log.Print(root.ChildByID(26))
	log.Print(root.Child("Cap"))
	log.Print(nxFile.Resolve("Cap", ""))
}
