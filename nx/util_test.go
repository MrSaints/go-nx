package nx

import (
	"testing"
)

const (
	TestString = "Hello World!"
)

func TestDo_readU16(t *testing.T) {
	if got, want := readU16([]byte(TestString)), uint16(25928); got != want {
		t.Errorf("readU16 returned %+v, want %+v", got, want)
	}
}

func TestDo_readU32(t *testing.T) {
	if got, want := readU32([]byte(TestString)), uint32(1819043144); got != want {
		t.Errorf("readU32 returned %+v, want %+v", got, want)
	}
}

func TestDo_read32(t *testing.T) {
	if got, want := read32([]byte(TestString)), int32(1819043144); got != want {
		t.Errorf("read32 returned %+v, want %+v", got, want)
	}
}

func TestDo_readU64(t *testing.T) {
	if got, want := readU64([]byte(TestString)), uint64(8022916924116329800); got != want {
		t.Errorf("readU64 returned %+v, want %+v", got, want)
	}
}

func TestDo_read64(t *testing.T) {
	if got, want := read64([]byte(TestString)), int64(8022916924116329800); got != want {
		t.Errorf("read64 returned %+v, want %+v", got, want)
	}
}

func TestDo_readFloat64(t *testing.T) {
	if got, want := readFloat64([]byte(TestString)), float64(2.1914441197069634e+228); got != want {
		t.Errorf("readFloat64 returned %+v, want %+v", got, want)
	}
}
