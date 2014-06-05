package	main

import (
	"log"
	"github.com/mrsaints/gonx"
)

func main() {
	nxFile := gonx.New("../Data/Character.nx")
	root := nxFile.Root()
	log.Print(root.ChildByID(26))
	log.Print(root.Child("Cap"))
	log.Print(nxFile.Resolve("Cap", ""))
}