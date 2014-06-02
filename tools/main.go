package	main

import (
	"log"
	"github.com/mrsaints/gonx"
)

func main() {
	nxFile := gonx.New("../Data/Character.nx")
	log.Print(nxFile.Node(1).Name())
}