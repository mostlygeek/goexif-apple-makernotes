# goexif-apple-makernotes

This repo adds support for Apple iOS makernote parsing to [rwcarlsen/goexif](https://github.com/rwcarlsen/goexif) until [PR#62](https://github.com/rwcarlsen/goexif/pull/62) lands.

## Usage

See extract.go as an example of usage.  

```go
func main() {
	// add the apple parser
	exif.RegisterParsers(makernotes.Apple)

    // ...

	x, err := exif.Decode(f)

    // ...

    tag, err := x.Get(makernotes.Apple_BurstUUID)

    // ...
}
```

Run it on the testdata/ files to see examples of apple makernote data.

```
$ go run extract.go testdata/apple_burstuuid.jpg
Apple.RunTime : (binary plist 104 bytes)
Apple.AccelerationVector : ["-5691/6191","-677/22055","-1538/3835"]
 --- Get Error:  exif: tag "Apple.HDRImageType" is not present
Apple.BurstUUID : "81F26A91-7A4F-4968-AA5A-9189D3681D91"
 --- Get Error:  exif: tag "Apple.ContentIdentifier" is not present
 --- Get Error:  exif: tag "Apple.ImageUniqueID" is not present
```

and

```
$ go run extract.go testdata/apple_contentid.jpg
Apple.RunTime : (binary plist 104 bytes)
Apple.AccelerationVector : ["-1340/16057","-812/905","-1205/2874"]
 --- Get Error:  exif: tag "Apple.HDRImageType" is not present
 --- Get Error:  exif: tag "Apple.BurstUUID" is not present
Apple.ContentIdentifier : "D4122CE4-5DA9-4EBA-9475-01579D01ADCE"
 --- Get Error:  exif: tag "Apple.ImageUniqueID" is not present
```

