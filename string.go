package gonx

func (NX *NXFile) String(index int) string {
    tableOffset := NX.Header.StringOffset + uint64(index) * 8
    stringOffset := ReadU64(NX.Raw[tableOffset:])
    length := ReadU16(NX.Raw[stringOffset:])
    return string(NX.Raw[stringOffset+2:stringOffset+2+uint64(length)])
}