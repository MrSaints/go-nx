package main

import (
	"github.com/mrsaints/go-nx/nx"
	"log"
)

func main() {
	nxf, err := nx.NewFile("../../data/Base.nx", true)
	failOnError(err)

	nxh, err := nxf.Header()
	failOnError(err)
	log.Printf("File header: %+v", nxh)

	root, err := nxf.Root()
	failOnError(err)
	log.Printf("Root: %+v", root)
	c, err := root.Children()
	failOnError(err)
	log.Printf("Root's children: %+v", c)

	//nd, err := nx.NewNode(nxf, 0)
	//failOnError(err)
	//err = nd.Parse()
	//failOnError(err)
	//log.Printf("Node: %+v", nd)
	//c, err := nd.Children()
	//failOnError(err)
	//log.Printf("Node's children: %+v", c)

	// Blocking operation
	printChildren(c)

	log.Printf("Total nodes: %+v", nxh.NodeCount)
}

func printChildren(c *nx.Children) {
	for i, nd := range c.Nodes {
		log.Printf("%+v: %+v | %+v | %+v", i, nd.Id, nd.Name, nd.Type)

		if nd.Count > 0 {
			c2, err := nd.Children()
			failOnError(err)
			printChildren(c2)
		}
	}
}

func failOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
