package gonx

import (
    "math"
)

func pError(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadU16(b []byte) uint16 { // 2
    return uint16(b[0]) | uint16(b[1])<<8
}

func ReadU32(b []byte) uint32 { // 4
    return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func Read32(b []byte) int32 {
    return int32(ReadU32(b))
}

func ReadU64(b []byte) uint64 { // 8
    return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
            uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func Read64(b []byte) int64 {
    return int64(ReadU64(b))
}

func ReadFloat64(b []byte) float64 { // IEEE 754 double-precision
    u64 := ReadU64(b)
    return math.Float64frombits(u64)
}