package nx

import (
	"testing"
)

func TestRoot(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	if _, err := f.Root(); err != nil {
		t.Fatalf("Root returned unexpected error: %v", err)
	}
}

func BenchmarkRoot(b *testing.B) {
	f, _ := NewFile(TEST_FILE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Root()
	}
}
