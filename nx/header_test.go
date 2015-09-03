package nx

import (
	"testing"
)

func TestNewHeader(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	h := &Header{f: f}
	if got, want := NewHeader(f).f, h.f; got != want {
		t.Errorf("NewHeader.f is %+v, want %+v", got, want)
	}
}

func TestHeaderParse(t *testing.T) {
	f, _ := NewFile(TEST_FILE)
	if got, want := f.Header.magic, PKG4; got != want {
		t.Errorf("h.Parse h.magic is %v, want %v", got, want)
	}
	if got, want := f.Header.nodeCount, uint32(440); got != want {
		t.Errorf("h.Parse h.nodeCount is %v, want %v", got, want)
	}
	if got, want := f.Header.nodeOffset, uint64(64); got != want {
		t.Errorf("h.Parse h.nodeOffset is %v, want %v", got, want)
	}
	if got, want := f.Header.stringCount, uint32(231); got != want {
		t.Errorf("h.Parse h.stringCount is %v, want %v", got, want)
	}
	if got, want := f.Header.stringOffset, uint64(8880); got != want {
		t.Errorf("h.Parse h.stringOffset is %v, want %v", got, want)
	}
	if got, want := f.Header.bitmapCount, uint32(0); got != want {
		t.Errorf("h.Parse h.bitmapCount is %v, want %v", got, want)
	}
	if got, want := f.Header.bitmapOffset, uint64(0); got != want {
		t.Errorf("h.Parse h.bitmapOffset is %v, want %v", got, want)
	}
	if got, want := f.Header.audioCount, uint32(0); got != want {
		t.Errorf("h.Parse h.audioCount is %v, want %v", got, want)
	}
	if got, want := f.Header.audioOffset, uint64(0); got != want {
		t.Errorf("h.Parse Header.audioOffset is %v, want %v", got, want)
	}
}
