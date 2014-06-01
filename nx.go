package gonx

import (
    "log"
    "os"
    "github.com/edsrzf/mmap-go"
)

type NXFile struct {
    Name            string
    Raw             mmap.MMap
    Header          Header
}

func New(fileName string) (NX *NXFile) {
    log.Print("Opening...")

    file, err := os.Open(fileName)
    pError(err)

    buffer, err := mmap.Map(file, mmap.RDONLY, 0)
    pError(err)

    NX = new(NXFile)
    NX.Name = fileName
    NX.Raw = buffer

    log.Print("Parsing header...")

    NX.ParseHeader()

    //log.Print(NX.Header.NodeOffset)

    //magic := []byte(NX.Header.Magic)
    //log.Print(ReadUint(magic))

    log.Print("Success.")
    return
}