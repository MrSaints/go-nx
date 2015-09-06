package nx

import (
	"fmt"
	"github.com/edsrzf/mmap-go"
	"os"
	"reflect"
	"testing"
)

const TestFile = "../data/Base.nx"

func loadTestFile() (mmap.MMap, error) {
	f, err := os.Open(TestFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return mmap.Map(f, mmap.RDONLY, 0)
}

func TestNewFile(t *testing.T) {
	tf, err := loadTestFile()
	if err != nil {
		t.Fatalf("Unable to load test file: %+v", err)
	}

	nxf, err := NewFile(TestFile, false)
	if err != nil {
		t.Fatalf("NewFile returned unexpected error: %+v", err)
	}
	if got, want := nxf.Fn, TestFile; got != want {
		t.Errorf("File name is %+v, want %+v", got, want)
	}
	if got, want := nxf.raw, tf; !reflect.DeepEqual(got, want) {
		t.Errorf("File buffer is %+v, want %+v", got, want)
	}
	if got, want := nxf.header, (Header{}); got != want {
		t.Errorf("File header is %+v, want %+v", got, want)
	}
}

func TestNewFile_noFile(t *testing.T) {
	_, err := NewFile("", false)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*os.PathError); !ok {
		t.Errorf("Expected an OS path error, got %+v", err)
	}
}

func TestFile_Parse(t *testing.T) {
	nxf, _ := NewFile(TestFile, false)
	if err := nxf.Parse(); err != nil {
		t.Fatalf("Parse returned unexpected error: %+v", err)
	}
	nxh, _ := nxf.Header()
	if got, want := nxh, nxf.header; !reflect.DeepEqual(got, want) {
		t.Errorf("Header returned %+v, want %+v", got, want)
	}
}

func TestFile_Parse_badInitialisation(t *testing.T) {
	nxf := new(File)
	if err := nxf.Parse(); err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func ExampleNewFile() {
	nxf, _ := NewFile(TestFile, false)
	fmt.Printf("File name: '%s'", nxf.Fn)
	// Output: File name: '../data/Base.nx'
}

func BenchmarkNewFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewFile(TestFile, false)
	}
}
