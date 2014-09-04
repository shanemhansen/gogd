package gogd

/*
 #include <gd.h>
 #cgo LDFLAGS: -lgd
extern int gogd_get_c(struct gdIOCtx *);
extern int gogd_get_buf(struct gdIOCtx *, const void*, int);
extern int gogd_put_c(struct gdIOCtx *, int);
extern int gogd_put_buf(struct gdIOCtx *, const void*, int);
*/
import "C"
import "unsafe"
import "io"

type IOCtx struct {
	GetC   unsafe.Pointer
	GetBuf unsafe.Pointer
	PutC   unsafe.Pointer
	PutBuf unsafe.Pointer
	Seek   unsafe.Pointer
	Free   unsafe.Pointer
	Data   unsafe.Pointer
}

// GetContext makes a gdio context
func getContext(g gdio) *C.gdIOCtx {
	return &C.gdIOCtx{
		getC:   (*[0]byte)(C.gogd_get_c),
		getBuf: (*[0]byte)(C.gogd_get_buf),
		putC:   (*[0]byte)(C.gogd_put_c),
		putBuf: (*[0]byte)(C.gogd_put_buf),
		data:   unsafe.Pointer(&g),
	}
}

func ImageCreateFromPngCtx(g io.Reader) C.gdImagePtr {
	ctx := getContext(g)

	return C.gdImageCreateFromPngCtx(ctx)
}
func ImageCreateTrueColor(sx, sy int) C.gdImagePtr {
	return C.gdImageCreateTrueColor(C.int(sx), C.int(sy))
}

func ImageCopyResampled(newimage, oldimage C.gdImagePtr, dstX,
	dstY, srcX, srcY, dstW, dstH,
	srcW, srcH int) {
	C.gdImageCopyResampled(newimage, oldimage, C.int(dstX), C.int(dstY), C.int(srcX), C.int(srcY), C.int(dstW),
		C.int(dstH), C.int(srcW), C.int(srcH))
}
func ImageCopyResized(newimage, oldimage C.gdImagePtr, dstX,
	dstY, srcX, srcY, dstW, dstH,
	srcW, srcH int) {
	C.gdImageCopyResized(newimage, oldimage, C.int(dstX), C.int(dstY), C.int(srcX), C.int(srcY), C.int(dstW),
		C.int(dstH), C.int(srcW), C.int(srcH))
}
func ImagePngCtx(img C.gdImagePtr, out io.Writer) {
	ctx := getContext(out)
	C.gdImagePngCtx(img, ctx)
}
func ImageDestroy(img C.gdImagePtr) {
	C.gdImageDestroy(img)
}
