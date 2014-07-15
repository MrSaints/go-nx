package gonx

import (
    "os"
    "strings"
    "github.com/edsrzf/mmap-go"
)

type NXFile struct {
    FileName        string
    Raw             mmap.MMap
    Header          Header
}

func New(fileName string) (NX *NXFile) {
    file, err := os.Open(fileName)
    pError(err)

    buffer, err := mmap.Map(file, mmap.RDONLY, 0)
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