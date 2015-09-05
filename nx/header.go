package nx

import (
	"errors"
)

const (
	PKG4 = "PKG4"
)

var (
	ErrFileInvalid = errors.New("this file is not an NX (PKG4) file or it has not been initialised properly with NewFile")
)

type Header struct {
	Magic        string
	NodeCount    uint32
	nodeOffset   uint64
	StringCount  uint32
	stringOffset uint64
	BitmapCount  uint32
	bitmapOffset uint64
	AudioCount   uint32
	audioOffset  uint64
}

func (nxf *File) Header() (Header, error) {
	if nxh := nxf.header; nxh != (Header{}) {
		return nxh, nil
	}
	if len(nxf.raw) < 4 {
		return Header{}, ErrFileInvalid
	}
	v := string(nxf.raw[0:4])
	if v != PKG4 {
		return Header{}, ErrFileInvalid
	}

	nxh := Header{}
	nxh.Magic = v
	nxh.NodeCount = readU32(nxf.raw[4:8])
	nxh.nodeOffset = readU64(nxf.raw[8:16])
	nxh.StringCount = readU32(nxf.raw[16:20])
	nxh.stringOffset = readU64(nxf.raw[20:28])
	nxh.BitmapCount = readU32(nxf.raw[28:32])
	nxh.bitmapOffset = readU64(nxf.raw[32:40])
	nxh.AudioCount = readU32(nxf.raw[40:44])
	nxh.audioOffset = readU64(nxf.raw[44:52])
	return nxh, nil
}
