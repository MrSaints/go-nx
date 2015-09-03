package nx

import (
	"testing"
)

func TestRoot(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	_, err := f.Root()
	if err != nil {
		t.Fatalf("Root returned unexpected error: %v", err)
	}
}
