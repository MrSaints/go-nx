package nx

import (
//"strings"
)

func (f *File) Root() (*Node, error) {
	rN := NewNode(f)
	err := rN.Parse(0)
	return rN, err
}
