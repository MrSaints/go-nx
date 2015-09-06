package nx

import (
	"fmt"
	"testing"
)

func TestFile_Root(t *testing.T) {
	nxf, _ := NewFile(TestFile, true)
	if _, err := nxf.Root(); err != nil {
		t.Fatalf("Root returned unexpected error: %+v", err)
	}
}

func TestFile_Root_badInitialisation(t *testing.T) {
	nxf := new(File)
	if _, err := nxf.Root(); err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func ExampleFile_Root() {
	nxf, _ := NewFile(TestFile, true)
	root, _ := nxf.Root()
	fmt.Printf("Total children: %v", root.Count)
	// Output: Total children: 18
}

func BenchmarkFile_Root(b *testing.B) {
	nxf, _ := NewFile(TestFile, true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nxf.Root()
	}
}
