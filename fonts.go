package gogd

/*
 #include <gdfonts.h>
 #cgo LDFLAGS: -lgd
*/
import "C"

func FontGetSmall() Font {
	return Font(C.gdFontGetSmall())
}
