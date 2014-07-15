package gonx

import (
    "testing"
)

func (node *Node) GetChildren() (nodes []*Node) {
    if node.Count > 0 {
        node.ParseChildren()
        return node.Children.Nodes
    }
    return
}

func Recurse(parentNode *Node) {
    for _, childNode := range parentNode.GetChildren() {
        Recurse(childNode)
    }
}

func BenchmarkRecurse(b *testing.B) {
    nxFile := New(TEST_FILE)
    root := nxFile.Root()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Recurse(root)
    }
}

func BenchmarkParseChildren(b *testing.B) {
    nxFile := New(TEST_FILE)
    root := nxFile.Root()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        root.ParseChildren()
    }
}

func BenchmarkChildByID(b *testing.B) {
    nxFile := New(TEST_FILE)
    root := nxFile.Root()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        root.ChildByID(26)
    }
}

func BenchmarkChild(b *testing.B) {
    nxFile := New(TEST_FILE)
    root := nxFile.Root()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        root.Child("Cap")
    }
}