package gonx

import (
	"testing"
)

func BenchmarkRoot(b *testing.B) {
	nxFile := New("./Data/Base.nx")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nxFile.Root()
	}
}

func BenchmarkChildren(b *testing.B) {
	nxFile := New("./Data/Base.nx")
	root := nxFile.Root()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root.Children()
	}
}
