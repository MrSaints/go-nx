package gonx

import (
    //"log"
)

type Children struct {
    Indexes     map[string]int
    Nodes       []*Node
}

type Node struct {
    *NXFile
    StringID    uint32
    ChildID     uint32
    Count       uint16
    Type        uint16
    Data        interface{}
    *Children
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

func (node *Node) ParseNode(index int) {
    if node.StringID != 0 || index >= int(node.NXFile.Header.NodeCount) {
        return
    }

    offset := node.NXFile.Header.NodeOffset + uint64(index) * 20
    buffer := node.NXFile.Raw

    stringID := ReadU32(buffer[offset:])
    node.StringID = stringID
    offset += 4
    node.ChildID = ReadU32(buffer[offset:])
    offset += 4
    node.Count = ReadU16(buffer[offset:])
    offset += 2
    node.Type = ReadU16(buffer[offset:])
    offset += 2

    switch node.Type {
        case 1: // Int64
            node.Data = LongNode{Read64(buffer[offset:])}
        case 2: // Double
            node.Data = FloatNode{ReadFloat64(buffer[offset:])}
        case 3: // NX_STRING (StringID)
            node.Data = StringNode{ReadU32(buffer[offset:])}
        case 4: // NX_VECTOR (X and Y)
            node.Data = VectorNode{Read32(buffer[offset:]), Read32(buffer[offset + 4:])}
        case 5: // NX_BITMAP (BitmapID, W, H)
            node.Data = BitmapNode{ReadU32(buffer[offset:]), ReadU16(buffer[offset + 4:]), ReadU16(buffer[offset + 6:])}
        case 6: // NX_AUDIO
            node.Data = AudioNode{ReadU32(buffer[offset:]), ReadU32(buffer[offset + 4:])}
    }
}

func (node *Node) ParseChildren() {
    if node.Count == 0 || node.Children != nil {
        return
    }

    children := new(Children)
    children.Indexes = make(map[string]int)

    for i := 0; i < int(node.Count); i++ {
        childNode := new(Node)
        childNode.NXFile = node.NXFile
        childNode.ParseNode(int(node.ChildID) + i)

        children.Indexes[childNode.Name()] = int(childNode.StringID)
        children.Nodes = append(children.Nodes, childNode)
    }

    node.Children = children
}

func (node *Node) Name() string {
    return node.NXFile.String(int(node.StringID))
}

func (node *Node) ChildByID(index int) *Node {
    if index < 0 || index >= int(node.Count) {
        return nil
    }

    node.ParseChildren()
    return node.Nodes[index]
}

func (node *Node) Child(index string) *Node {
    node.ParseChildren()
    if value, ok := node.Indexes[index]; node.Children != nil && ok {
        return node.Nodes[value]
    }
    return nil
}