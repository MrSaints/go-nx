package nx

import (
//"strings"
)

func (nxf *File) Root() (*Node, error) {
	rnd, err := NewNode(nxf, 0)
	if err != nil {
		return nil, err
	}
	err = rnd.Parse()
	return rnd, err
}
