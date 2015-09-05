package nx

import (
	"testing"
)

func getNode(i uint, parse bool) (*Node, error) {
	nxf, _ := NewFile(TestFile, parse)
	return NewNode(nxf, i)
}

func getChildren(i uint) (*Children, error) {
	nd, _ := getNode(i, true)
	_ = nd.Parse()
	return nd.Children()
}

func TestNewNode(t *testing.T) {
	nxf, _ := NewFile(TestFile, true)
	nd, err := NewNode(nxf, 0)
	if err != nil {
		t.Fatalf("NewNode returned unexpected error: %+v", err)
	}
	if got, want := nd.f, nxf; got != want {
		t.Errorf("Node file is %+v, want %+v", got, want)
	}
	if got, want := nd.Id, uint(0); got != want {
		t.Errorf("Node id is %+v, want %+v", got, want)
	}
}

func TestNewNode_invalidIndex(t *testing.T) {
	nxf, _ := NewFile(TestFile, true)
	_, err := NewNode(nxf, uint(nxf.header.NodeCount))
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrNodeIndex {
		t.Errorf("Expected a node index error, got %+v", err)
	}
}

func TestNodeParse(t *testing.T) {
	nd, _ := getNode(1, true)
	if err := nd.Parse(); err != nil {
		t.Fatalf("Parse returned unexpected error: %+v", err)
	}
	if got, want := nd.Name, "Character"; got != want {
		t.Errorf("Node name is %+v, want %+v", got, want)
	}
	if got, want := nd.ChildId, uint32(19); got != want {
		t.Errorf("Node first child id is %+v, want %+v", got, want)
	}
	if got, want := nd.Count, uint16(0); got != want {
		t.Errorf("Node children count is %+v, want %+v", got, want)
	}
	if got, want := nd.Type, uint16(0); got != want {
		t.Errorf("Node type is %+v, want %+v", got, want)
	}
	if got, want := nd.Data, (EmptyNode{}); got != want {
		t.Errorf("Node data is %+v, want %+v", got, want)
	}
}

func TestNodeParse_repeatedCall(t *testing.T) {
	nd, _ := getNode(1, true)
	_ = nd.Parse()
	err := nd.Parse()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrNodeParsed {
		t.Errorf("Expected a node parsed error, got %+v", err)
	}
}

func TestNodeParse_badInitialisation(t *testing.T) {
	nd := new(Node)
	err := nd.Parse()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrNodeFile {
		t.Errorf("Expected a node parsed error, got %+v", err)
	}
}

func TestChildren(t *testing.T) {
	if _, err := getChildren(0); err != nil {
		t.Fatalf("Children returned unexpected error: %+v", err)
	}
}

func TestChildren_noParse(t *testing.T) {
	nd, _ := getNode(0, true)
	_, err := nd.Children()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrNodeNoChild {
		t.Errorf("Expected a no children node error, got %+v", err)
	}
}

func TestChildrenGet(t *testing.T) {
	c, _ := getChildren(0)
	got, err := c.Get("Reactor")
	if err != nil {
		t.Fatalf("Child returned unexpected error: %+v", err)
	}
	want := c.Nodes[c.Indexes["Reactor"]]
	if got != want {
		t.Errorf("Child returned %+v, want %+v", got, want)
	}
}

func TestChildrenGet_invalidIndex(t *testing.T) {
	c, _ := getChildren(0)
	_, err := c.Get("Invalid Name")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrNodeIndex {
		t.Errorf("Expected a node index error, got %+v", err)
	}
}

func TestChildrenGetById(t *testing.T) {
	c, _ := getChildren(0)
	got, err := c.GetById(0)
	if err != nil {
		t.Fatalf("ChildById returned unexpected error: %+v", err)
	}
	want := c.Nodes[0]
	if got != want {
		t.Errorf("ChildById returned %+v, want %+v", got, want)
	}
}

func TestChildrenGetById_invalidIndex(t *testing.T) {
	c, _ := getChildren(0)
	_, err := c.GetById(c.Total)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrNodeIndex {
		t.Errorf("Expected a node index error, got %+v", err)
	}
}
