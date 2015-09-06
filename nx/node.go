package nx

import (
	"errors"
)

var (
	ErrNodeIndex   = errors.New("the node index provided does not exist / it exceeds the total number of nodes")
	ErrNodeParsed  = errors.New("this node has already been parsed")
	ErrNodeFile    = errors.New("this node has not been initialised properly with NewNode")
	ErrNodeNoChild = errors.New("this node does not have any children or it is not yet parsed")
)

type Children struct {
	Indexes map[string]uint
	Nodes   []*Node
	Total   uint
}

type Node struct {
	f       *File
	c       *Children
	ID      uint
	Name    string
	ChildID uint32
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
	return &Node{f: nxf, ID: i}, nil
}

func (nd *Node) Parse() error {
	if nd.Name != "" {
		return ErrNodeParsed
	}

	if nd.f == nil {
		return ErrNodeFile
	}

	offset := nd.f.header.nodeOffset + uint64(nd.ID)*20

	sid := readU32(nd.f.raw[offset:])
	name, err := nd.f.GetString(uint(sid))
	if err != nil {
		return err
	}
	nd.Name = name
	offset += 4
	nd.ChildID = readU32(nd.f.raw[offset:])
	offset += 4
	nd.Count = readU16(nd.f.raw[offset:])
	offset += 2
	nd.Type = readU16(nd.f.raw[offset:])
	offset += 2
	nd.Data = nd.GetData(offset)

	return nil
}

func (nd *Node) GetData(offset uint64) interface{} {
	// TODO: Reader interface
	switch nd.Type {
	case 0:
		return EmptyNode{}
	case 1: // Int64
		return LongNode{read64(nd.f.raw[offset:])}
	case 2: // Double
		return FloatNode{readFloat64(nd.f.raw[offset:])}
	case 3: // NX_STRING (String Id)
		return StringNode{readU32(nd.f.raw[offset:])}
	case 4: // NX_VECTOR (X and Y)
		return VectorNode{
			read32(nd.f.raw[offset:]),
			read32(nd.f.raw[offset+4:]),
		}
	case 5: // NX_BITMAP (BitmapID, W, H)
		return BitmapNode{
			readU32(nd.f.raw[offset:]),
			readU16(nd.f.raw[offset+4:]), readU16(nd.f.raw[offset+6:]),
		}
	case 6: // NX_AUDIO
		return AudioNode{
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
	if nd.c != nil {
		return nd.c, nil
	}

	nd.c = new(Children)
	nd.c.Indexes = make(map[string]uint, totalNodes)
	nd.c.Nodes = make([]*Node, totalNodes)
	nd.c.Total = totalNodes

	for i := uint(0); i < totalNodes; i++ {
		cnd, err := NewNode(nd.f, uint(nd.ChildID)+i)
		if err != nil {
			return nil, err
		}
		err = cnd.Parse()
		if err != nil {
			return nil, err
		}
		nd.c.Indexes[cnd.Name] = i
		nd.c.Nodes[i] = cnd
	}
	return nd.c, nil
}

func (c *Children) Get(n string) (*Node, error) {
	if i, ok := c.Indexes[n]; c.Nodes[i] != nil && ok {
		return c.Nodes[i], nil
	}
	return nil, ErrNodeIndex
}

func (c *Children) GetByID(i uint) (*Node, error) {
	if i >= c.Total {
		return nil, ErrNodeIndex
	}
	return c.Nodes[i], nil
}
