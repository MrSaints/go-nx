package nx

import (
//"strings"
)

func (nxf *File) Root() (*Node, error) {
	rnd := NewNode(nxf)
	err := rnd.Parse(0)
	return rnd, err
}
