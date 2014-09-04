package gogd

// #include <gd_io.h>
// #cgo LDFLAGS: -lgd
import "C"
import "io"
import "log"
import "reflect"
import "unsafe"
import "os"

type gdio interface {
	io.Reader
	io.Writer
	io.Seeker
}

//export gogd_get_c
func gogd_get_c(ctx *C.gdIOCtx) int {
	gdio := (*(*gdio)(ctx.data))
	buf := make([]byte, 1)
	_, err := gdio.Read(buf)
	if err != nil {
		log.Println(err)
	}
	return int(buf[0])
}

//export gogd_get_buf
func gogd_get_buf(ctx *C.gdIOCtx, cbuf unsafe.Pointer, l C.int) int {
	gdio := (*(*gdio)(ctx.data))
	buf := GoSliceFromCString((*C.char)(cbuf), int(l))
	n, err := gdio.Read(buf)
	if err != nil {
		log.Println(err)
		return 0
	}
	return n
}

//export gogd_put_buf
func gogd_put_buf(ctx *C.gdIOCtx, cbuf unsafe.Pointer, l C.int) int {
	gdio := (*(*gdio)(ctx.data))
	buf := GoSliceFromCString((*C.char)(cbuf), int(l))
	n, err := gdio.Write(buf)
	if err != nil {
		log.Println(err)
		return 0
	}
	return n
}

//export gogd_put_c
func gogd_put_c(ctx *C.gdIOCtx, c int) {
	gdio := (*(*gdio)(ctx.data))
	buf := []byte{byte(c)}
	_, err := gdio.Write(buf)
	if err != nil {
		log.Println(err)
	}
}

//export gogd_seek
func gogd_seek(ctx *C.gdIOCtx, c int) int {
	gdio := (*(*gdio)(ctx.data))
	n, err := gdio.Seek(int64(c), 0)
	if err != nil {
		log.Println(err)
	}
	return int(n)
}

//export gogd_tell
func gogd_tell(ctx *C.gdIOCtx) int {
	gdio := (*(*gdio)(ctx.data))
	n, err := gdio.Seek(int64(os.SEEK_CUR), 0)
	if err != nil {
		log.Println(err)
	}
	return int(n)
}

// GoSliceFromCString provides a zero copy interface for returning a go slice backed by a c array.
func GoSliceFromCString(cArray *C.char, size int) []byte {
	//See http://code.google.com/p/go-wiki/wiki/cgo
	//It turns out it's really easy to
	//make a string from a *C.char and vise versa.
	//not so easy to write to a c array.
	var sliceHeader reflect.SliceHeader
	sliceHeader.Cap = size
	sliceHeader.Len = size
	sliceHeader.Data = uintptr(unsafe.Pointer(cArray))
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}
