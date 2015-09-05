package nx

import (
	"errors"
)

var (
	ErrStringIndex = errors.New("the string index provided exceeds the total number of string entries")
)

func (nxf *File) GetString(i uint) (string, error) {
	if i >= uint(nxf.header.StringCount) {
		return "", ErrStringIndex
	}

	tableOffset := nxf.header.stringOffset + uint64(i)*8
	stringOffset := readU64(nxf.raw[tableOffset:])
	length := readU16(nxf.raw[stringOffset:])
	stringOffset = stringOffset + 2
	return string(nxf.raw[stringOffset : stringOffset+uint64(length)]), nil
}
