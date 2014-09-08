package gogd

/*
 #include <gd.h>
 #include <gd_color_map.h>
 #cgo LDFLAGS: -lgd
extern int gogd_get_c(struct gdIOCtx *);
extern int gogd_get_buf(struct gdIOCtx *, const void*, int);
extern int gogd_put_c(struct gdIOCtx *, int);
extern int gogd_put_buf(struct gdIOCtx *, const void*, int);
*/
import "C"
import "unsafe"

// GetContext makes a gdio context
func getContext(g gdio) *IOCtx {
	return (*IOCtx)(unsafe.Pointer(&C.gdIOCtx{
		getC:   (*[0]byte)(C.gogd_get_c),
		getBuf: (*[0]byte)(C.gogd_get_buf),
		putC:   (*[0]byte)(C.gogd_put_c),
		putBuf: (*[0]byte)(C.gogd_put_buf),
		data:   unsafe.Pointer(&g),
	}))
}

func toGdioCtx(i *IOCtx) *C.gdIOCtx {
	return (*C.gdIOCtx)(unsafe.Pointer(i))
}
func fromGdioCtx(i *C.gdIOCtx) *IOCtx {
	return (*IOCtx)(unsafe.Pointer(i))
}

// IOCbPtr is an unsafe pointer to a callback function
type IOCbPtr unsafe.Pointer
type IOCtx struct {
	getC    IOCbPtr
	getBuf  IOCbPtr
	putC    IOCbPtr
	putBuf  IOCbPtr
	seek    IOCbPtr
	tell    IOCbPtr
	gd_free IOCbPtr
	data    unsafe.Pointer
}
