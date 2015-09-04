package nx

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	h := Node{f: f}
	if got, want := *n, h; got != want {
		t.Errorf("NewNode is %+v, want %+v", got, want)
	}
}

func TestNodeParse(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	if err := n.Parse(1); err != nil {
		t.Fatalf("Parse returned unexpected error: %+v", err)
	}
	if got, want := n.Name, "Character"; got != want {
		t.Errorf("Node name is %v, want %v", got, want)
	}
	if got, want := n.ChildID, uint32(19); got != want {
		t.Errorf("Node first child ID is %v, want %v", got, want)
	}
	if got, want := n.Count, uint16(0); got != want {
		t.Errorf("Node children count is %v, want %v", got, want)
	}
	if got, want := n.Type, uint16(0); got != want {
		t.Errorf("Node type is %v, want %v", got, want)
	}
	if got, want := n.Data, interface{}(nil); got != want {
		t.Errorf("Node data is %v, want %v", got, want)
	}
}

func TestNodeParse_uninitialised(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	n.Name = "Initialised Name"
	if n.Parse(0) == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestNodeParse_wrongIndex(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	if n := NewNode(f); n.Parse(uint(n.f.header.nodeCount)) == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestNodeChildren(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	_ = n.Parse(0)
	if _, err := n.Children(); err != nil {
		t.Fatalf("Children returned unexpected error: %+v", err)
	}
}

func TestNodeChildren_noParse(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	if _, err := n.Children(); err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestGet(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	_ = n.Parse(0)
	c, _ := n.Children()
	got, _ := c.Get(0)
	want := c.Nodes[0]
	if got != want {
		t.Errorf("Get returned %+v, want %+v", got, want)
	}
}

func TestGet_wrongIndex(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	_ = n.Parse(0)
	c, _ := n.Children()
	if _, err := c.Get(c.Total); err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestGetByName(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	_ = n.Parse(0)
	c, _ := n.Children()
	cN, err := c.GetByName("Reactor")
	if err != nil {
		t.Fatalf("GetByName returned unexpected error: %+v", err)
	}
	if got, want := cN.Name, "Reactor"; got != want {
		t.Errorf("Node name is %v, want %v", got, want)
	}
}

func TestGetByName_wrongName(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	_ = n.Parse(0)
	c, _ := n.Children()
	if _, err := c.GetByName("Invalid Name"); err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestChild(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	_ = n.Parse(0)
	cN, err := n.Child(10)
	if err != nil {
		t.Fatalf("Child returned unexpected error: %+v", err)
	}
	c, _ := n.Children()
	tcN, _ := c.Get(10)
	if got, want := *cN, *tcN; got != want {
		t.Errorf("Child returned %+v, want %+v", got, want)
	}
}

func TestChild_noParse(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	if _, err := n.Child(10); err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestChildByName(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	_ = n.Parse(0)
	cN, err := n.ChildByName("Quest")
	if err != nil {
		t.Fatalf("ChildByName returned unexpected error: %+v", err)
	}
	c, _ := n.Children()
	tcN, _ := c.GetByName("Quest")
	if got, want := *cN, *tcN; got != want {
		t.Errorf("ChildByName returned %+v, want %+v", got, want)
	}
}

func TestChildByName_noParse(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	if _, err := n.ChildByName("Quest"); err == nil {
		t.Errorf("Expected error to be returned")
	}
}
