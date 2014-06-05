package gonx

import (
	"testing"
)

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