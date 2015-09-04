package nx

import (
	"encoding/binary"
	"math"
)

func readU16(b []byte) uint16 { // 2
	return binary.LittleEndian.Uint16(b)
}

func readU32(b []byte) uint32 { // 4
	return binary.LittleEndian.Uint32(b)
}

func read32(b []byte) int32 {
	return int32(readU32(b))
}

func readU64(b []byte) uint64 { // 8
	return binary.LittleEndian.Uint64(b)
}

func read64(b []byte) int64 {
	return int64(readU64(b))
}

func readFloat64(b []byte) float64 { // IEEE 754 double-precision
	return math.Float64frombits(readU64(b))
}
