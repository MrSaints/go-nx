package nx

import (
	"github.com/edsrzf/mmap-go"
	"os"
)

type File struct {
	Header *Header
	fn     string
	raw    mmap.MMap
}

func NewFile(fn string) (*File, error) {
	f, err := os.Open(fn)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	buffer, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		return nil, err
	}

	nxf := new(File)
	nxf.fn = fn
	nxf.raw = buffer

	nxf.Header = NewHeader(nxf)
	err = nxf.Header.Parse()

	return nxf, err
}

func (nx *File) GetString(index uint) string {
	tableOffset := nx.Header.stringOffset + uint64(index)*8
	stringOffset := readU64(nx.raw[tableOffset:])
	length := readU16(nx.raw[stringOffset:])
	return string(nx.raw[stringOffset+2 : stringOffset+2+uint64(length)])
}
