package nx

import (
	"errors"
)

type Children struct {
	Indexes map[string]uint
	Nodes   []*Node
	Total   uint
}

type Node struct {
	f       *File
	Name    string
	ChildID uint32
	Count   uint16
	Type    uint16
	Data    interface{}
}

type LongNode struct {
	Value int64
}

type FloatNode struct {
	Value float64
}

type StringNode struct {
	Value uint32
}

type VectorNode struct {
	X, Y int32
}

type BitmapNode struct {
	ID     uint32
	Width  uint16
	Height uint16
}

type AudioNode struct {
	ID     uint32
	Length uint32
}

func NewNode(f *File) *Node {
	return &Node{f: f}
}

func (n *Node) Parse(i uint) error {
	if n.Name != "" {
		err := errors.New("Cannot unmarshal an initialised node. It may be corrupted.")
		return err
	}
	if i >= uint(n.f.Header.nodeCount) {
		err := errors.New("The node index provided does not exists. It exceeds the total number of nodes.")
		return err
	}

	offset := n.f.Header.nodeOffset + uint64(i)*20
	buffer := n.f.raw

	stringID := readU32(buffer[offset:])
	n.Name = n.f.GetString(uint(stringID))
	offset += 4
	n.ChildID = readU32(buffer[offset:])
	offset += 4
	n.Count = readU16(buffer[offset:])
	offset += 2
	n.Type = readU16(buffer[offset:])
	offset += 2

	switch n.Type {
	case 1: // Int64
		n.Data = LongNode{read64(buffer[offset:])}
	case 2: // Double
		n.Data = FloatNode{readFloat64(buffer[offset:])}
	case 3: // NX_STRING (StringID)
		n.Data = StringNode{readU32(buffer[offset:])}
	case 4: // NX_VECTOR (X and Y)
		n.Data = VectorNode{
			read32(buffer[offset:]),
			read32(buffer[offset+4:]),
		}
	case 5: // NX_BITMAP (BitmapID, W, H)
		n.Data = BitmapNode{
			readU32(buffer[offset:]),
			readU16(buffer[offset+4:]), readU16(buffer[offset+6:]),
		}
	case 6: // NX_AUDIO
		n.Data = AudioNode{
			readU32(buffer[offset:]),
			readU32(buffer[offset+4:]),
		}
	}
	return nil
}

func (n *Node) Children() (*Children, error) {
	totalNodes := uint(n.Count)
	if totalNodes == 0 {
		err := errors.New("This node does not have any children or it is not yet parsed.")
		return nil, err
	}

	c := new(Children)
	c.Indexes = make(map[string]uint, totalNodes)
	c.Nodes = make([]*Node, totalNodes)
	c.Total = totalNodes

	for i := uint(0); i < totalNodes; i++ {
		cN := NewNode(n.f)
		err := cN.Parse(uint(n.ChildID) + i)
		if err != nil {
			return nil, err
		}
		c.Indexes[cN.Name] = i
		c.Nodes[i] = cN
	}
	return c, nil
}

func (c *Children) Get(i uint) (*Node, error) {
	if i >= c.Total {
		err := errors.New("The child node index provided does not exist. It exceeds the total number of child nodes.")
		return nil, err
	}
	return c.Nodes[i], nil
}

func (c *Children) GetByName(n string) (*Node, error) {
	if i, ok := c.Indexes[n]; c.Nodes[i] != nil && ok {
		return c.Nodes[i], nil
	}
	err := errors.New("The child node name provided does not exist.")
	return nil, err
}

func (n *Node) Child(i uint) (*Node, error) {
	c, err := n.Children()
	if err != nil {
		return nil, err
	}
	return c.Get(i)
}

func (n *Node) ChildByName(id string) (*Node, error) {
	c, err := n.Children()
	if err != nil {
		return nil, err
	}
	return c.GetByName(id)
}
