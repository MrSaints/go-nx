package nx

import (
	"github.com/edsrzf/mmap-go"
	"os"
	"testing"
)

const TEST_FILE = "../data/Base.nx"

func loadTestFile() (mmap.MMap, error) {
	f, err := os.Open(TEST_FILE)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return mmap.Map(f, mmap.RDONLY, 0)
}

func TestNewFile(t *testing.T) {
	f, err := NewFile(TEST_FILE)
	if err != nil {
		t.Fatalf("NewFile returned unexpected error: %v", err)
	}
	if got, want := f.fn, TEST_FILE; got != want {
		t.Errorf("NewFile.fn is %v, want %v", got, want)
	}
}

func TestNewFile_noFile(t *testing.T) {
	_, err := NewFile("")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*os.PathError); !ok {
		t.Errorf("Expected a *os.PathError, got %+v", err)
	}
}
