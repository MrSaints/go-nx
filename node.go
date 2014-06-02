package gonx

import (
    //"log"
)

type Node struct {
    *NXFile
    StringID    uint32
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

func (NX *NXFile) Node(index int) (node *Node) {
    node = new(Node)
    node.NXFile = NX
    offset := NX.Header.NodeOffset + uint64(index) * 20

    stringID := ReadU32(NX.Raw[offset:])
    node.StringID = stringID
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

    //NX.Indexes = make(map[string]int)
    //NX.Indexes[NX.GetString(int(stringID))] = int(stringID)

    return
}

func (NX *NXFile) Root() *Node {
    return NX.GetNode(0)
}

func (node *Node) Name() string {
    return node.NXFile.GetString(int(node.StringID))
}