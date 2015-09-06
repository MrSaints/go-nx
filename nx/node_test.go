package nx

import (
	"fmt"
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
	if got, want := nd.ID, uint(0); got != want {
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

func TestNode_Parse(t *testing.T) {
	nd, _ := getNode(1, true)
	if err := nd.Parse(); err != nil {
		t.Fatalf("Parse returned unexpected error: %+v", err)
	}
	if got, want := nd.Name, "Character"; got != want {
		t.Errorf("Node name is %+v, want %+v", got, want)
	}
	if got, want := nd.ChildID, uint32(19); got != want {
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

func TestNode_Parse_repeatedCall(t *testing.T) {
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

func TestNode_Parse_badInitialisation(t *testing.T) {
	nd := new(Node)
	err := nd.Parse()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrNodeFile {
		t.Errorf("Expected a node parsed error, got %+v", err)
	}
}

func TestNode_GetData(t *testing.T) {
	nd, _ := getNode(1, true)
	_ = nd.Parse()

	if got, ok := nd.GetData(0).(EmptyNode); !ok {
		t.Errorf("GetData returned %+v, want EmptyNode{}", got)
	}
	nd.Type = 1
	if got, ok := nd.GetData(0).(LongNode); !ok {
		t.Errorf("GetData returned %+v, want LongNode{}", got)
	}
	nd.Type = 2
	if got, ok := nd.GetData(0).(FloatNode); !ok {
		t.Errorf("GetData returned %+v, want FloatNode{}", got)
	}
	nd.Type = 3
	if got, ok := nd.GetData(0).(StringNode); !ok {
		t.Errorf("GetData returned %+v, want StringNode{}", got)
	}
	nd.Type = 4
	if got, ok := nd.GetData(0).(VectorNode); !ok {
		t.Errorf("GetData returned %+v, want VectorNode{}", got)
	}
	nd.Type = 5
	if got, ok := nd.GetData(0).(BitmapNode); !ok {
		t.Errorf("GetData returned %+v, want BitmapNode{}", got)
	}
	nd.Type = 6
	if got, ok := nd.GetData(0).(AudioNode); !ok {
		t.Errorf("GetData returned %+v, want AudioNode{}", got)
	}
	nd.Type = 7
	if got := nd.GetData(0); got != nil {
		t.Errorf("GetData returned %+v, want nil", got)
	}
}

func TestNode_Children(t *testing.T) {
	if _, err := getChildren(0); err != nil {
		t.Fatalf("Children returned unexpected error: %+v", err)
	}
}

func TestNode_Children_noParse(t *testing.T) {
	nd, _ := getNode(0, true)
	_, err := nd.Children()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrNodeNoChild {
		t.Errorf("Expected a no children node error, got %+v", err)
	}
}

func TestNode_Children_repeatedCall(t *testing.T) {
	nd, _ := getNode(0, true)
	_ = nd.Parse()
	got, _ := nd.Children()
	want, _ := nd.Children()
	if got != want {
		t.Errorf("Children returned %p, want %p", got, want)
	}
}

func TestChildren_Get(t *testing.T) {
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

func TestChildren_Get_invalidIndex(t *testing.T) {
	c, _ := getChildren(0)
	_, err := c.Get("Invalid Name")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrNodeIndex {
		t.Errorf("Expected a node index error, got %+v", err)
	}
}

func TestChildren_GetByID(t *testing.T) {
	c, _ := getChildren(0)
	got, err := c.GetByID(0)
	if err != nil {
		t.Fatalf("ChildById returned unexpected error: %+v", err)
	}
	want := c.Nodes[0]
	if got != want {
		t.Errorf("ChildById returned %+v, want %+v", got, want)
	}
}

func TestChildren_GetByID_invalidIndex(t *testing.T) {
	c, _ := getChildren(0)
	_, err := c.GetByID(c.Total)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrNodeIndex {
		t.Errorf("Expected a node index error, got %+v", err)
	}
}

func ExampleNode_Parse() {
	nxf, _ := NewFile(TestFile, true)
	nd, _ := NewNode(nxf, 12)
	_ = nd.Parse()
	fmt.Printf("Node name: %s", nd.Name)
	// Output: Node name: Sound
}
