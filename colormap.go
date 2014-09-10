package gogd

/*
 #include <gd_color_map.h>
 #cgo LDFLAGS: -lgd
*/
import "C"
import "unsafe"

type ColorMapEntry struct {
	ColorName string
	Red       int
	Green     int
	Blue      int
}
type ColorMap struct {
	colorMap C.gdColorMap
}

func ColorMapLookup(colorMap ColorMap, colorName string) *ColorMapEntry {
	var r, g, b C.int
	ret := C.gdColorMapLookup(colorMap.colorMap, cStringFromGoString(colorName),
		&r, &g, &b)
	if ret == 0 {
		return nil
	}
	return &ColorMapEntry{Red: int(r), Green: int(g), Blue: int(b)}
}

// cStringFromGoString returns a null terminated string
func cStringFromGoString(data string) *C.char {
	slice := make([]byte, len(data)+1)
	for i, c := range []byte(data) {
		slice[i] = c
	}
	// last element of slice will be null
	return (*C.char)(unsafe.Pointer(&slice[0]))
}
