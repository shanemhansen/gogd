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
import "errors"

type Image interface {
	img() C.gdImagePtr
	Write(io.Writer) error
	Read(io.Reader) error
	FromBuffer([]byte) error
	ToBuffer() ([]byte, error)
}
type GdImage struct {
	i C.gdImagePtr
}

func (img *GdImage) Png() *Png {
	return &Png{GdImage: img}
}
func (img *GdImage) Jpeg() *Jpeg {
	return &Jpeg{GdImage: img}
}
func (img *GdImage) Gif() *Gif {
	return &Gif{GdImage: img}
}
func (img *GdImage) Destroy() {
	C.gdImageDestroy(img.i)
}
func (img *GdImage) img() C.gdImagePtr {
	return img.i
}
func (img *GdImage) CopyResampled(dst Image, dstx, dsty, srcx,
	srcy, dstw, dsth, srcw, srch int) {
	C.gdImageCopyResampled(dst.img(), img.i, C.int(dstx),
		C.int(dsty), C.int(srcx), C.int(srcy),
		C.int(dstw),
		C.int(dsth), C.int(srcw), C.int(srch))
}

type IOCtx struct {
	GetC   unsafe.Pointer
	GetBuf unsafe.Pointer
	PutC   unsafe.Pointer
	PutBuf unsafe.Pointer
	Seek   unsafe.Pointer
	Free   unsafe.Pointer
	Data   unsafe.Pointer
}

func NewPng() *Png {
	return &Png{GdImage: new(GdImage)}
}

type Gif struct {
	*GdImage
}
type Png struct {
	*GdImage
}
type Jpeg struct {
	*GdImage
	quality int
}
type Webp struct {
	i C.gdImagePtr
}

func (i *Gif) Read(r io.Reader) error {
	ctx := getContext(r)
	i.i = C.gdImageCreateFromGifCtx(ctx)
	return nil
}
func (i *Gif) Write(w io.Writer) error {
	ctx := getContext(w)
	C.gdImagePngCtx(i.i, ctx)
	return nil
}
func (i *Png) Read(r io.Reader) error {
	ctx := getContext(r)
	i.i = C.gdImageCreateFromPngCtx(ctx)
	return nil
}
func (i *Png) Write(w io.Writer) error {
	ctx := getContext(w)
	C.gdImagePngCtx(i.i, ctx)
	return nil
}
func (i *Png) FromBuffer(buf []byte) error {
	i.i = C.gdImageCreateFromPngPtr(C.int(len(buf)), unsafe.Pointer(&buf[0]))
	return nil
}
func (i *Png) ToBuffer() ([]byte, error) {
	// for now copy
	var size C.int
	data := C.gdImagePngPtr(i.i, &size)
	if size <= 0 { // don't know if this signals an error
		return nil, errors.New("couldn't turn image into buffer")
	}
	src := GoSliceFromCString((*C.char)(data), int(size))
	dst := make([]byte, size)
	copy(dst, src)
	C.gdFree(data)
	return dst, nil
}
func (i *Jpeg) Read(r io.Reader) error {
	ctx := getContext(r)
	i.i = C.gdImageCreateFromJpegCtx(ctx)
	return nil
}
func (i *Jpeg) Write(w io.Writer) error {
	ctx := getContext(w)
	C.gdImageJpegCtx(i.i, ctx, C.int(i.quality))
	return nil
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

func ImageCreate(sx, sy int) *GdImage {
	return &GdImage{i: C.gdImageCreate(C.int(sx), C.int(sy))}
}
func ImageCreateTrueColor(sx, sy int) *GdImage {
	return &GdImage{i: C.gdImageCreateTrueColor(C.int(sx), C.int(sy))}
}
