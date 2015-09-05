package nx

import (
	"testing"
)

func TestGetString(t *testing.T) {
	nxf, _ := NewFile(TestFile, true)
	s, err := nxf.GetString(6)
	if err != nil {
		t.Fatalf("GetString returned unexpected error: %+v", err)
	}
	if got, want := s, "Map"; got != want {
		t.Errorf("GetString returned %+v, want %+v", got, want)
	}
}

func TestGetString_invalidIndex(t *testing.T) {
	nxf, _ := NewFile(TestFile, true)
	_, err := nxf.GetString(uint(nxf.header.StringCount))
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrStringIndex {
		t.Errorf("Expected a string index error, got %+v", err)
	}
}
