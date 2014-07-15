package gonx

import (
    "errors"
)

type Header struct {
    Magic           string
    NodeCount       uint32
    NodeOffset      uint64
    StringCount     uint32
    StringOffset    uint64
    BitmapCount     uint32
    BitmapOffset    uint64
    AudioCount      uint32
    AudioOFfset     uint64
}

func (NX *NXFile) ParseHeader() (header Header) {
    header = Header{ Magic: string(NX.Raw[0:4]) }

    if header.Magic != "PKG4" {
        err := errors.New(NX.FileName + " is not a PKG4 NX file.")
        pError(err)
    }

    header.NodeCount = ReadU32(NX.Raw[4:8])
    header.NodeOffset = ReadU64(NX.Raw[8:16])
    header.StringCount = ReadU32(NX.Raw[16:20])
    header.StringOffset = ReadU64(NX.Raw[20:28])
    header.BitmapCount = ReadU32(NX.Raw[28:32])
    header.BitmapOffset = ReadU64(NX.Raw[32:40])
    header.AudioCount = ReadU32(NX.Raw[40:44])
    header.AudioOFfset = ReadU64(NX.Raw[44:52])
    return
}