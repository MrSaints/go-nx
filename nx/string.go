package nx

func (nxf *File) GetString(i uint) string {
	tableOffset := nxf.header.stringOffset + uint64(i)*8
	stringOffset := readU64(nxf.raw[tableOffset:])
	length := readU16(nxf.raw[stringOffset:])
	stringOffset = stringOffset + 2
	return string(nxf.raw[stringOffset : stringOffset+uint64(length)])
}
