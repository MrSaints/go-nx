package gonx

import (
    "io/ioutil"
    //"log"
    //"os"
)

type NXFile struct {
    FileName        string
    Raw             []byte
    Header          *Header
}

func New(fileName string) (NX *NXFile) {
    buffer, err := ioutil.ReadFile(fileName)
    pError(err)

    NX = new(NXFile)
    NX.FileName = fileName
    NX.Raw = buffer

    NX.Header = NX.ParseHeader()
    return
}

func (NX *NXFile) Root() (node *Node) {
    node = new(Node)
    node.NXFile = NX
    node.ParseNode(0)
    return
}