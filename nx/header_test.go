package nx

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFile_Header(t *testing.T) {
	nxf, _ := NewFile(TestFile, false)
	nxh, err := nxf.Header()
	if err != nil {
		t.Fatalf("Header returned unexpected error: %+v", err)
	}

	if got, want := nxh.Magic, PKG4; got != want {
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

func TestFile_Header_repeatedCall(t *testing.T) {
	nxf, _ := NewFile(TestFile, true)
	nxf.header.Magic = TestString
	nxh, _ := nxf.Header()
	if got, want := nxh.Magic, nxf.header.Magic; !reflect.DeepEqual(got, want) {
		t.Errorf("Magic / version is %+v, want %+v", got, want)
	}
}

func TestFile_Header_badInitialisation(t *testing.T) {
	nxf := new(File)
	_, err := nxf.Header()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrFileInvalid {
		t.Errorf("Expected an invalid file error, got %+v", err)
	}
}

func TestFile_Header_invalidFile(t *testing.T) {
	nxf := &File{raw: []byte(TestString)}
	_, err := nxf.Header()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err != ErrFileInvalid {
		t.Errorf("Expected an invalid file error, got %+v", err)
	}
}

func ExampleFile_Header() {
	nxf, _ := NewFile(TestFile, false)
	nxh, _ := nxf.Header()
	fmt.Printf("Magic / version: %s", nxh.Magic)
	// Output: Magic / version: PKG4
}
