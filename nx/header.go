package nx

import (
	"errors"
)

const (
	PKG4 = "PKG4"
)

var (
	ErrFileUninitialised = errors.New("this file has not been initialised properly with NewFile")
	ErrFileInvalid       = errors.New("this file is not an NX (PKG4) file")
)

type Header struct {
	magic        string
	nodeCount    uint32
	nodeOffset   uint64
	stringCount  uint32
	stringOffset uint64
	bitmapCount  uint32
	bitmapOffset uint64
	audioCount   uint32
	audioOffset  uint64
}

func (nxf *File) Header() (Header, error) {
	if nxh := nxf.header; nxh != nil {
		return nxh, nil
	}
	if len(nxf.raw) < 4 {
		return nil, ErrFileUninitialised
	}
	v := string(nxf.raw[0:4])
	if v != PKG4 {
		return nil, ErrFileInvalid
	}
	nxh := new(Header)
	nxh.magic = v
	nxh.nodeCount = readU32(nxf.raw[4:8])
	nxh.nodeOffset = readU64(nxf.raw[8:16])
	nxh.stringCount = readU32(nxf.raw[16:20])
	nxh.stringOffset = readU64(nxf.raw[20:28])
	nxh.bitmapCount = readU32(nxf.raw[28:32])
	nxh.bitmapOffset = readU64(nxf.raw[32:40])
	nxh.audioCount = readU32(nxf.raw[40:44])
	nxh.audioOffset = readU64(nxf.raw[44:52])
	return nxh, nil
}
