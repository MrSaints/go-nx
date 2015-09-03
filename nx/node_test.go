package nx

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	h := &Node{f: f}
	if got, want := NewNode(f).f, h.f; got != want {
		t.Errorf("NewNode.f is %+v, want %+v", got, want)
	}
}

func TestNodeParse(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	if err := n.Parse(1); err != nil {
		t.Fatalf("n.Parse returned unexpected error: %v", err)
	}
	if got, want := n.Name, "Character"; got != want {
		t.Errorf("n.Parse n.Name is %v, want %v", got, want)
	}
	if got, want := n.ChildID, uint32(19); got != want {
		t.Errorf("n.Parse n.ChildID is %v, want %v", got, want)
	}
	if got, want := n.Count, uint16(0); got != want {
		t.Errorf("n.Parse n.Count is %v, want %v", got, want)
	}
	if got, want := n.Type, uint16(0); got != want {
		t.Errorf("n.Parse n.Type is %v, want %v", got, want)
	}
	if got, want := n.Data, interface{}(nil); got != want {
		t.Errorf("n.Parse n.Data is %v, want %v", got, want)
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
	if n := NewNode(f); n.Parse(uint(n.f.Header.nodeCount)+1) == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestNodeChildren(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	_ = n.Parse(0)
	if _, err := n.Children(); err != nil {
		t.Fatalf("n.Children returned unexpected error: %v", err)
	}
}

func TestNodeChildren_noParse(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	if _, err := n.Children(); err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestChildren(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	_ = n.Parse(0)
	c, _ := n.Children()
	got, _ := c.Get(0)
	want := c.Nodes[0]
	if got != want {
		t.Errorf("n.Children returned %v, want %v", got, want)
	}
}

func TestChildrenGet_wrongIndex(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	n := NewNode(f)
	_ = n.Parse(0)
	c, _ := n.Children()
	if _, err := c.Get(c.Total + 1); err == nil {
		t.Errorf("Expected error to be returned")
	}
}
