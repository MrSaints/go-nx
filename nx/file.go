package nx

import (
	"github.com/edsrzf/mmap-go"
	"os"
)

type File struct {
	Fn     string
	header Header
	raw    mmap.MMap
}

func NewFile(fn string, p bool) (*File, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		return nil, err
	}

	nxf := new(File)
	nxf.Fn = fn
	nxf.raw = buf

	if p {
		err = nxf.Parse()
	}

	return nxf, err
}

func (nxf *File) Parse() error {
	nxh, err := nxf.Header()
	if err != nil {
		return err
	}
	nxf.header = nxh
	return nil
}
