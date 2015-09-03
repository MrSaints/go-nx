package nx

import (
	"testing"
)

func TestNewHeader(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	h := NewHeader(f)
	th := Header{f: f}
	if got, want := *h, th; got != want {
		t.Errorf("NewHeader is %+v, want %+v", got, want)
	}
}

func TestHeaderParse(t *testing.T) {
	tf, err := loadTestFile()
	if err != nil {
		t.Fatalf("Unable to load test file: %v", err)
	}

	nxf := &File{raw: tf}
	h := NewHeader(nxf)
	if h.Parse() != nil {
		t.Fatalf("Parse returned unexpected error: %v", err)
	}

	if got, want := h.magic, PKG4; got != want {
		t.Errorf("Magic / version is %v, want %v", got, want)
	}
	if got, want := h.nodeCount, uint32(440); got != want {
		t.Errorf("Node count is %v, want %v", got, want)
	}
	if got, want := h.nodeOffset, uint64(64); got != want {
		t.Errorf("Node offset is %v, want %v", got, want)
	}
	if got, want := h.stringCount, uint32(231); got != want {
		t.Errorf("String count is %v, want %v", got, want)
	}
	if got, want := h.stringOffset, uint64(8880); got != want {
		t.Errorf("String offset is %v, want %v", got, want)
	}
	if got, want := h.bitmapCount, uint32(0); got != want {
		t.Errorf("Bitmap count is %v, want %v", got, want)
	}
	if got, want := h.bitmapOffset, uint64(0); got != want {
		t.Errorf("Bitmap offset is %v, want %v", got, want)
	}
	if got, want := h.audioCount, uint32(0); got != want {
		t.Errorf("Audio count is %v, want %v", got, want)
	}
	if got, want := h.audioOffset, uint64(0); got != want {
		t.Errorf("Audio offset is %v, want %v", got, want)
	}
}

func TestHeaderParse_noFile(t *testing.T) {
	h := NewHeader(nil)
	err := h.Parse()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}
