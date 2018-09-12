package makernotes

import (
	"bytes"
	"encoding/binary"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

const (
	// Apple-specific fields
	Apple_RunTime            exif.FieldName = "Apple.RunTime"
	Apple_AccelerationVector exif.FieldName = "Apple.AccelerationVector"
	Apple_HDRImageType       exif.FieldName = "Apple.HDRImageType"
	Apple_BurstUUID          exif.FieldName = "Apple.BurstUUID"
	Apple_ContentIdentifier  exif.FieldName = "Apple.ContentIdentifier"
	Apple_ImageUniqueID      exif.FieldName = "Apple.ImageUniqueID"
)

// Source: https://www.sno.phy.queensu.ca/~phil/exiftool/TagNames/Apple.html
var makerNoteAppleFields = map[uint16]exif.FieldName{
	0x0003: Apple_RunTime, // undefined type but it is a binary serialized CMTime structure
	0x0008: Apple_AccelerationVector,
	0x000a: Apple_HDRImageType,
	0x000b: Apple_BurstUUID,
	0x0011: Apple_ContentIdentifier,
	0x0015: Apple_ImageUniqueID,
}

var Apple = &apple{}

type apple struct{}

// Parse decodes Apple makernote data found in x and adds it to x
func (_ *apple) Parse(x *exif.Exif) error {
	m, err := x.Get(exif.MakerNote)
	if err != nil {
		return nil
	} else if bytes.Compare(m.Val[:10], []byte("Apple iOS\000")) != 0 {
		return nil
	}

	// Apple makenotes is a self contained IFD and offsets are relative
	// to the start of the of the maker note.  Apple does not write a special
	// tiff marker to the metadata
	buf := bytes.NewReader(m.Val)
	buf.Seek(14, 0) // skip header, endian marker

	appleDir, _, err := tiff.DecodeDir(buf, binary.BigEndian)
	if err != nil {
		return err
	}

	x.LoadTags(appleDir, makerNoteAppleFields, false)
	return nil
}
