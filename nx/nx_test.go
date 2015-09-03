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
