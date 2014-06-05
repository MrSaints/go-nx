package gonx

import (
	"testing"
)

const TEST_FILE = "./Data/Base.nx"

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(TEST_FILE)
	}
}

func BenchmarkRoot(b *testing.B) {
	nxFile := New(TEST_FILE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nxFile.Root()
	}
}

func BenchmarkResolve(b *testing.B) {
	nxFile := New(TEST_FILE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nxFile.Resolve("Cap", "")
	}
}