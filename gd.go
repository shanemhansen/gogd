package gogd

/*
 #include <gd.h>
 #cgo LDFLAGS: -lgd
*/
import "C"
import "unsafe"
import "errors"

func (i Image) SquareToCircle(radius int) Image {
	return newImage(C.gdImageSquareToCircle(i.ptr, C.int(radius)))
}
func (i Image) StringFTCircle(cx, cy int, radius, textRadius, fillPortion float64, font string, points float64, top string, bottom string, fgcolor int) error {
	err := C.gdImageStringFTCircle(i.ptr, C.int(cx), C.int(cy), C.double(radius),
		C.double(textRadius), C.double(fillPortion), CStringFromGoString(font),
		C.double(points), CStringFromGoString(top), CStringFromGoString(bottom),
		C.int(fgcolor))
	if err == nil {
		return nil
	}
	// null terminated string.
	c := err
	count := 0
	for *c != 0 {
		c = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + 8))
		count++
	}
	msg := GoSliceFromCString(err, count)
	// make a copy of msg
	return errors.New(string(msg))
}
func (i Image) Sharpen(pct int) {
	C.gdImageSharpen(i.ptr, C.int(pct))
}
