package gonx

import (
    "log"
    "os"
    "github.com/edsrzf/mmap-go"
)

type NXFile struct {
    FileName        string
    Raw             mmap.MMap
    Header          *Header
    //Indexes         map[string]int
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