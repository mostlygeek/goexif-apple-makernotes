package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mostlygeek/goexif-apple-makernotes/makernotes"
	"github.com/rwcarlsen/goexif/exif"
)

func main() {
	// add the apple parser
	exif.RegisterParsers(makernotes.Apple)

	flag.Parse()
	filename := flag.Arg(0)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("File error: ", err.Error())
		return
	}

	x, err := exif.Decode(f)
	if err != nil {
		fmt.Println("Decode error: ", err.Error())
		return
	}

	notefields := []exif.FieldName{
		makernotes.Apple_RunTime,
		makernotes.Apple_AccelerationVector,
		makernotes.Apple_HDRImageType,
		makernotes.Apple_BurstUUID,
		makernotes.Apple_ContentIdentifier,
		makernotes.Apple_ImageUniqueID,
	}

	for _, fn := range notefields {
		t, err := x.Get(fn)
		if err != nil {
			fmt.Println(" --- Get Error: ", err.Error())
		} else if fn == makernotes.Apple_RunTime {
			fmt.Printf("%s : (binary plist %d bytes)\n", fn, len(t.Val))
		} else {
			j, _ := t.MarshalJSON()
			fmt.Printf("%s : %s\n", fn, j)
		}
	}

}
