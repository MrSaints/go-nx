package gonx

import (
    "errors"
    //"log"
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

func (NX *NXFile) ParseHeader() {
    NX.Header = Header{}
    NX.Header.Magic = string(NX.Raw[0:4])

    if NX.Header.Magic != "PKG4" {
        err := errors.New(NX.Name + " is not a PKG4 NX file.")
        pError(err)
    }

    NX.Header.NodeCount = ReadU32(NX.Raw[4:8])
    NX.Header.NodeOffset = ReadU64(NX.Raw[8:16])
    NX.Header.StringCount = ReadU32(NX.Raw[16:20])
    NX.Header.StringOffset = ReadU64(NX.Raw[20:28])
    NX.Header.BitmapCount = ReadU32(NX.Raw[28:32])
    NX.Header.BitmapOffset = ReadU64(NX.Raw[32:40])
    NX.Header.AudioCount = ReadU32(NX.Raw[40:44])
    NX.Header.AudioOFfset = ReadU64(NX.Raw[44:52])
}

type Node struct {
    Name        string
    ChildID     uint32
    Count       uint16
    Type        uint16
    Data        interface{}
}

type LongNode struct {
    Value   int64
}

type FloatNode struct {
    Value   float64
}

type StringNode struct {
    Value   uint32
}

type VectorNode struct {
    X, Y    int32
}

type BitmapNode struct {
    ID      uint32
    Width   uint16
    Height  uint16
}

type AudioNode struct {
    ID      uint32
    Length  uint32
}

func (NX *NXFile) GetNode(index int) (node Node) {
    node = Node{}
    offset := NX.Header.NodeOffset + uint64(index) * 20

    stringID := int(ReadU32(NX.Raw[offset:]))
    node.Name = NX.GetString(stringID)
    offset += 4
    node.ChildID = ReadU32(NX.Raw[offset:])
    offset += 4
    node.Count = ReadU16(NX.Raw[offset:])
    offset += 2
    node.Type = ReadU16(NX.Raw[offset:])
    offset += 2

    switch node.Type {
        case 1: // Int64
            node.Data = LongNode{Read64(NX.Raw[offset:])}
        case 2: // Double
            node.Data = FloatNode{ReadFloat64(NX.Raw[offset:])}
        case 3: // NX_STRING (StringID)
            node.Data = StringNode{ReadU32(NX.Raw[offset:])}
        case 4: // NX_VECTOR (X and Y)
            node.Data = VectorNode{Read32(NX.Raw[offset:]), Read32(NX.Raw[offset + 4:])}
        case 5: // NX_BITMAP (BitmapID, W, H)
            node.Data = BitmapNode{ReadU32(NX.Raw[offset:]), ReadU16(NX.Raw[offset + 4:]), ReadU16(NX.Raw[offset + 6:])}
        case 6: // NX_AUDIO
            node.Data = AudioNode{ReadU32(NX.Raw[offset:]), ReadU32(NX.Raw[offset + 4:])}
    }
    return
}

func (NX *NXFile) GetString(index int) string {
    tableOffset := NX.Header.StringOffset + uint64(index) * 8
    stringOffset := ReadU64(NX.Raw[tableOffset:])
    length := ReadU16(NX.Raw[stringOffset:])
    return string(NX.Raw[stringOffset+2:stringOffset+2+uint64(length)])
}