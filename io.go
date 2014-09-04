package gogd

/*
 #include <gd_io.h>
 #include <gd.h>
 #cgo LDFLAGS: -lgd
extern int gogd_get_c(struct gdIOCtx *);
extern int gogd_get_buf(struct gdIOCtx *, const void*, int);
extern int gogd_put_c(struct gdIOCtx *, int);
extern int gogd_put_buf(struct gdIOCtx *, const void*, int);
*/
import "C"
import "unsafe"

// GetContext makes a gdio context
func GetContext(g gdio) *C.gdIOCtx {
	return &C.gdIOCtx{
		getC:   (*[0]byte)(C.gogd_get_c),
		getBuf: (*[0]byte)(C.gogd_get_buf),
		putC:   (*[0]byte)(C.gogd_put_c),
		putBuf: (*[0]byte)(C.gogd_put_buf),
		data:   unsafe.Pointer(&g),
	}
}

func gdImageCreateFromPngCtx(g gdio) C.gdImagePtr {
	return nil
}
