package	main

import (
	"github.com/mrsaints/gonx"
)

func main() {
	nxFile := gonx.Open("../Data/Character.nx")
	nxFile.Parse()
}