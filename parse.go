package gonx

import (
	"errors"
	"log"
)

func (this *NXFile) Parse() {

	this.Header = Header{}
	this.Header.Magic = string(this.Map[0:4])

	if this.Header.Magic != "PKG4" {
		err := errors.New(this.Name + " is not a PKG4 NX file.")
		pError(err)
	}

	this.Header.NodeCount = ReadUint(this.Map[4:8])
	this.Header.NodeOffset = ReadDouble(this.Map[8:16])
	this.Header.StringCount = ReadUint(this.Map[16:20])
	this.Header.StringOffset = ReadDouble(this.Map[20:28])
	this.Header.BitmapCount = ReadUint(this.Map[28:32])
	this.Header.BitmapOffset = ReadDouble(this.Map[32:40])
	this.Header.AudioCount = ReadUint(this.Map[40:44])
	this.Header.AudioOFfset = ReadDouble(this.Map[44:52])

	log.Print(this.Header.NodeOffset)

	//magic := []byte(this.Header.Magic)
	//log.Print(ReadUint(magic))

	log.Print("Success.")
}