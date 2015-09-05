package nx

import (
	"testing"
)

func TestHeader(t *testing.T) {
	nxf, _ := NewFile(TestFile, false)
	nxh, err := nxf.Header()
	if err != nil {
		t.Fatalf("Header returned unexpected error: %+v", err)
	}

	if got, want := nxh.magic, PKG4; got != want {
		t.Errorf("Magic / version is %+v, want %+v", got, want)
	}
	if got, want := nxh.NodeCount, uint32(440); got != want {
		t.Errorf("Node count is %+v, want %+v", got, want)
	}
	if got, want := nxh.nodeOffset, uint64(64); got != want {
		t.Errorf("Node offset is %+v, want %+v", got, want)
	}
	if got, want := nxh.StringCount, uint32(231); got != want {
		t.Errorf("String count is %+v, want %+v", got, want)
	}
	if got, want := nxh.stringOffset, uint64(8880); got != want {
		t.Errorf("String offset is %+v, want %+v", got, want)
	}
	if got, want := nxh.BitmapCount, uint32(0); got != want {
		t.Errorf("Bitmap count is %+v, want %+v", got, want)
	}
	if got, want := nxh.bitmapOffset, uint64(0); got != want {
		t.Errorf("Bitmap offset is %+v, want %+v", got, want)
	}
	if got, want := nxh.AudioCount, uint32(0); got != want {
		t.Errorf("Audio count is %+v, want %+v", got, want)
	}
	if got, want := nxh.audioOffset, uint64(0); got != want {
		t.Errorf("Audio offset is %+v, want %+v", got, want)
	}
}

func TestHeader_repeatedCall(t *testing.T) {
	nxf, _ := NewFile(TestFile, true)
	got, _ := nxf.Header()
	want := nxf.header
	if got != want {
		t.Errorf("Header returned %+v, want %+v", got, want)
	}
}

func TestHeader_badInitialisation(t *testing.T) {
	nxf := new(File)
	_, err := nxf.Header()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrFileInvalid {
		t.Errorf("Expected an invalid file error, got %+v", err)
	}
}

func TestHeader_invalidFile(t *testing.T) {
	nxf := &File{raw: []byte("PKG1")}
	_, err := nxf.Header()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrFileInvalid {
		t.Errorf("Expected an invalid file error, got %+v", err)
	}
}
