package nx

import (
	"errors"
)

const (
	PKG4 = "PKG4"
)

type Header struct {
	f            *File
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

func NewHeader(f *File) *Header {
	return &Header{f: f}
}

func (h *Header) Parse() error {
	v := string(h.f.raw[0:4])
	if v != PKG4 {
		err := errors.New(h.f.fn + " is not a PKG4 NX file.")
		return err
	}
	h.magic = v
	h.nodeCount = readU32(h.f.raw[4:8])
	h.nodeOffset = readU64(h.f.raw[8:16])
	h.stringCount = readU32(h.f.raw[16:20])
	h.stringOffset = readU64(h.f.raw[20:28])
	h.bitmapCount = readU32(h.f.raw[28:32])
	h.bitmapOffset = readU64(h.f.raw[32:40])
	h.audioCount = readU32(h.f.raw[40:44])
	h.audioOffset = readU64(h.f.raw[44:52])
	return nil
}
