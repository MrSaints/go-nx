package nx

import (
	"errors"
)

var (
	ErrNodeIndex   = errors.New("the node index provided does not exist / it exceeds the total number of nodes")
	ErrNodeParsed  = errors.New("this node has already been parsed")
	ErrNodeNoChild = errors.New("this node does not have any children or it is not yet parsed")
)

type Children struct {
	Indexes map[string]uint
	Nodes   []*Node
	Total   uint
}

type Node struct {
	f       *File
	Id      uint
	Name    string
	ChildId uint32
	Count   uint16
	Type    uint16
	Data    interface{}
}

type DataNode interface {
	Value() interface{}
}

type EmptyNode struct {
}

type LongNode struct {
	d int64
}

type FloatNode struct {
	d float64
}

type StringNode struct {
	d uint32
}

type VectorNode struct {
	x, y int32
}

type BitmapNode struct {
	id uint32
	w  uint16
	h  uint16
}

type AudioNode struct {
	id     uint32
	length uint32
}

func NewNode(nxf *File, i uint) (*Node, error) {
	if i >= uint(nxf.header.NodeCount) {
		return nil, ErrNodeIndex
	}
	return &Node{f: nxf, Id: i}, nil
}

func (nd *Node) Parse() error {
	// TODO: Check for empty nd.f
	if nd.Name != "" {
		return ErrNodeParsed
	}

	offset := nd.f.header.nodeOffset + uint64(nd.Id)*20

	sid := readU32(nd.f.raw[offset:])
	nd.Name = nd.f.GetString(uint(sid))
	offset += 4
	nd.ChildId = readU32(nd.f.raw[offset:])
	offset += 4
	nd.Count = readU16(nd.f.raw[offset:])
	offset += 2
	nd.Type = readU16(nd.f.raw[offset:])
	offset += 2

	// WIP / TODO
	switch nd.Type {
	case 0:
		nd.Data = EmptyNode{}
	case 1: // Int64
		nd.Data = LongNode{read64(nd.f.raw[offset:])}
	case 2: // Double
		nd.Data = FloatNode{readFloat64(nd.f.raw[offset:])}
	case 3: // NX_STRING (String Id)
		nd.Data = StringNode{readU32(nd.f.raw[offset:])}
	case 4: // NX_VECTOR (X and Y)
		nd.Data = VectorNode{
			read32(nd.f.raw[offset:]),
			read32(nd.f.raw[offset+4:]),
		}
	case 5: // NX_BITMAP (BitmapID, W, H)
		nd.Data = BitmapNode{
			readU32(nd.f.raw[offset:]),
			readU16(nd.f.raw[offset+4:]), readU16(nd.f.raw[offset+6:]),
		}
	case 6: // NX_AUDIO
		nd.Data = AudioNode{
			readU32(nd.f.raw[offset:]),
			readU32(nd.f.raw[offset+4:]),
		}
	}
	return nil
}

func (nd *Node) Children() (*Children, error) {
	totalNodes := uint(nd.Count)
	if totalNodes == 0 {
		return nil, ErrNodeNoChild
	}

	c := new(Children)
	c.Indexes = make(map[string]uint, totalNodes)
	c.Nodes = make([]*Node, totalNodes)
	c.Total = totalNodes

	for i := uint(0); i < totalNodes; i++ {
		cnd, err := NewNode(nd.f, uint(nd.ChildId)+i)
		if err != nil {
			return nil, err
		}
		err = cnd.Parse()
		if err != nil {
			return nil, err
		}
		c.Indexes[cnd.Name] = i
		c.Nodes[i] = cnd
	}
	return c, nil
}

func (c *Children) Get(n string) (*Node, error) {
	if i, ok := c.Indexes[n]; c.Nodes[i] != nil && ok {
		return c.Nodes[i], nil
	}
	return nil, ErrNodeIndex
}

func (c *Children) GetById(i uint) (*Node, error) {
	if i >= c.Total {
		return nil, ErrNodeIndex
	}
	return c.Nodes[i], nil
}
