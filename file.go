package gonx

import (
	"os"
    "github.com/edsrzf/mmap-go"
)

type NXFile struct {
	Name			string
	Map				mmap.MMap
	Header			Header
}

type Header struct {
	Magic			string
	NodeCount		uint32
	NodeOffset		uint64
	StringCount		uint32
	StringOffset	uint64
	BitmapCount		uint32
	BitmapOffset	uint64
	AudioCount		uint32
	AudioOFfset		uint64
}

func Open(fileName string) (nxFile *NXFile) {
	file, err := os.Open(fileName)
	pError(err)

	buffer, err := mmap.Map(file, mmap.RDONLY, 0)
	pError(err)

	nxFile = new(NXFile)
	nxFile.Name = fileName
	nxFile.Map = buffer
	return
}