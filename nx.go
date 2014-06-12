package gonx

import (
    "io/ioutil"
    "strings"
)

type NXFile struct {
    FileName        string
    Raw             []byte
    Header          Header
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

func (NX *NXFile) Resolve(path string, separator string) *Node {
    if separator == "" {
        separator = "/"
    }

    if path == separator {
        return NX.Root()
    }

    nodes := strings.Split(path, separator)
    cursor := NX.Root()

    for i := 0; i < len(nodes); i++ {
        if cursor == nil {
            return nil
        }
        cursor = cursor.Child(nodes[i])
    }

    return cursor
}