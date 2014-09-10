package gogd

/*
 #include <gd.h>
 #cgo LDFLAGS: -lgd
*/
import "C"
import "unsafe"
import "io"

type Image struct {
	ptr C.gdImagePtr
}

func (i *Image) Valid() bool {
	return i != nil && i.ptr != nil
}

type Point struct {
	P C.gdPoint
}
type PointF struct {
	P C.gdPointF
}
type ImageWriter interface {
	Decode(io.Reader) Image
	Encode(Image, io.Writer)
}

func newImage(d C.gdImagePtr) Image {
	return Image{ptr: d}
}

type PngIO struct{}

func (*PngIO) Decode(w io.Reader) Image {
	return newImage(C.gdImageCreateFromPngCtx(toGdioCtx(getContext(w))))
}
func (*PngIO) Encode(i Image, r io.Writer) {
	c := getContext(r)
	C.gdImagePngCtx(i.ptr, toGdioCtx(c))
}

type GifIO struct{}

func (*GifIO) Decode(w io.Reader) Image {
	return newImage(C.gdImageCreateFromGifCtx(toGdioCtx(getContext(w))))
}
func (*GifIO) Encode(i Image, r io.Writer) {
	c := getContext(r)
	C.gdImageGifCtx(i.ptr, toGdioCtx(c))
}

type WBMPIO struct {
	ForeGround int
}

func (*WBMPIO) Decode(w io.Reader) Image {
	return newImage(C.gdImageCreateFromWBMPCtx(toGdioCtx(getContext(w))))
}
func (w *WBMPIO) Encode(i Image, r io.Writer) {
	c := getContext(r)
	C.gdImageWBMPCtx(i.ptr, C.int(w.ForeGround), toGdioCtx(c))
}

type JpegIO struct {
	Quality int
}

func (*JpegIO) Decode(w io.Reader) Image {
	return newImage(C.gdImageCreateFromJpegCtx(toGdioCtx(getContext(w))))
}
func (w *JpegIO) Encode(i Image, r io.Writer) {
	c := getContext(r)
	C.gdImageJpegCtx(i.ptr, toGdioCtx(c), C.int(w.Quality))
}

type WebpIO struct {
	Quantization int
}

func (*WebpIO) Decode(w io.Reader) Image {
	return newImage(C.gdImageCreateFromWebpCtx(toGdioCtx(getContext(w))))
}
func (w *WebpIO) Encode(i Image, r io.Writer) {
	c := getContext(r)
	C.gdImageWebpCtx(i.ptr, toGdioCtx(c), C.int(w.Quantization))
}

type TiffIO struct{}

func (*TiffIO) Decode(w io.Reader) Image {
	return newImage(C.gdImageCreateFromTiffCtx(toGdioCtx(getContext(w))))
}
func (*TiffIO) Encode(i Image, r io.Writer) {
	c := getContext(r)
	C.gdImageTiffCtx(i.ptr, toGdioCtx(c))
}

type BmpIO struct {
	Compression int
}

func (*BmpIO) Decode(w io.Reader) Image {
	return newImage(C.gdImageCreateFromBmpCtx(toGdioCtx(getContext(w))))
}
func (w *BmpIO) Encode(i Image, r io.Writer) {
	c := getContext(r)
	C.gdImageBmpCtx(i.ptr, toGdioCtx(c), C.int(w.Compression))
}

// ImageCreate makes a 256 color image
func ImageCreate(sx, sy int) Image {
	return newImage(C.gdImageCreate(C.int(sx), C.int(sy)))
}

// ImageCreateTrueColor creates a truecolor image
func ImageCreateTrueColor(sx, sy int) Image {
	return newImage(C.gdImageCreateTrueColor(C.int(sx), C.int(sy)))
}

func (i Image) SetPixel(x, y, color int) {
	C.gdImageSetPixel(i.ptr, C.int(x), C.int(y), C.int(color))
}
func (i Image) Pixel(x, y int) int {
	return int(C.gdImageGetPixel(i.ptr, C.int(x), C.int(y)))
}
func (i Image) TrueColorPixel(x, y int) int {
	return int(C.gdImageGetTrueColorPixel(i.ptr, C.int(x), C.int(y)))
}
func (i Image) AABlend() {
	C.gdImageAABlend(i.ptr)
}
func (i Image) Line(x1, y1, x2, y2, color int) {
	C.gdImageLine(i.ptr, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(color))
}
func (i Image) DashedLine(x1, y1, x2, y2, color int) {
	C.gdImageDashedLine(i.ptr, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(color))
}

func (i Image) Rectangle(x1, y1, x2, y2, color int) {
	C.gdImageRectangle(i.ptr, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(color))
}
func (i Image) FilledRectangle(x1, y1, x2, y2, color int) {
	C.gdImageFilledRectangle(i.ptr, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(color))
}
func (i Image) SetClip(x1, y1, x2, y2 int) {
	C.gdImageSetClip(i.ptr, C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}
func (i Image) GetClip() (x1, x2, y1, y2 int) {
	var x1p, y1p, x2p, y2p C.int
	C.gdImageGetClip(i.ptr, &x1p, &y1p, &x2p, &y2p)
	x1 = int(x1p)
	y1 = int(y1p)
	x2 = int(x2p)
	y2 = int(y2p)
	return
}
func (i Image) SetResolution(resx, resy uint) {
	C.gdImageSetResolution(i.ptr, C.uint(resx), C.uint(resy))
}
func (i Image) BoundsSafe(x, y int) int {
	return int(C.gdImageBoundsSafe(i.ptr, C.int(x), C.int(y)))
}

type Font C.gdFontPtr

func (i Image) Char(f Font, x, y, c, color int) {
	C.gdImageChar(i.ptr, f, C.int(x), C.int(y), C.int(c), C.int(color))
}
func (i Image) CharUp(f Font, x, y, c, color int) {
	C.gdImageCharUp(i.ptr, f, C.int(x), C.int(y), C.int(c), C.int(color))
}
func (i Image) String(f Font, x, y int, s string, color int) {
	str := (*C.uchar)(unsafe.Pointer(cStringFromGoString(s)))
	C.gdImageString(i.ptr, f, C.int(x), C.int(y), str, C.int(color))
}
func (i Image) StringUp(f Font, x, y int, s string, color int) {
	str := (*C.uchar)(unsafe.Pointer(cStringFromGoString(s)))
	C.gdImageStringUp(i.ptr, f, C.int(x), C.int(y), str, C.int(color))
}

func FontCacheSetup() {
	C.gdFontCacheSetup()
}
func FontCacheShutdown() {
	C.gdFontCacheShutdown()
}

func FTUseFontConfig(flag int) int {
	return int(C.gdFTUseFontConfig(C.int(flag)))
}

func (i Image) Polygon(p Point, n, c int) {
	C.gdImagePolygon(i.ptr, &p.P, C.int(n), C.int(c))
}
func (i Image) OpenPolygon(p Point, n, c int) {
	C.gdImageOpenPolygon(i.ptr, &p.P, C.int(n), C.int(c))
}
func (i Image) ColorAllocate(r, g, b int) int {
	return int(C.gdImageColorAllocate(i.ptr, C.int(r), C.int(g), C.int(b)))
}
func (i Image) ColorAllocateAlpha(r, g, b, a int) int {
	return int(C.gdImageColorAllocateAlpha(i.ptr, C.int(r), C.int(g), C.int(b), C.int(a)))
}
func (i Image) ColorClosest(r, g, b int) int {
	return int(C.gdImageColorClosest(i.ptr, C.int(r), C.int(g), C.int(b)))
}
func (i Image) ColorClosestAlpha(r, g, b, a int) int {
	return int(C.gdImageColorClosestAlpha(i.ptr, C.int(r), C.int(g), C.int(b), C.int(a)))
}
func (i Image) ColorClosestHWB(r, g, b int) int {
	return int(C.gdImageColorClosestHWB(i.ptr, C.int(r), C.int(g), C.int(b)))
}
func (i Image) ColorExact(r, g, b int) int {
	return int(C.gdImageColorExact(i.ptr, C.int(r), C.int(g), C.int(b)))
}
func (i Image) ColorExactAlpha(r, g, b, a int) int {
	return int(C.gdImageColorExactAlpha(i.ptr, C.int(r), C.int(g), C.int(b), C.int(a)))
}
func (i Image) ColorResolve(r, g, b int) int {
	return int(C.gdImageColorResolve(i.ptr, C.int(r), C.int(g), C.int(b)))
}
func (i Image) ColorResolveAlpha(r, g, b, a int) int {
	return int(C.gdImageColorResolveAlpha(i.ptr, C.int(r), C.int(g), C.int(b), C.int(a)))
}

func (i Image) ColorDeallocate(color int) {
	C.gdImageColorDeallocate(i.ptr, C.int(color))
}
func (i Image) CreatePaletteFromTrueColor(ditherFlag, colorsWanted int) Image {
	return newImage(C.gdImageCreatePaletteFromTrueColor(i.ptr, C.int(ditherFlag), C.int(colorsWanted)))
}

func (i Image) TrueColorToPalette(ditherFlag, colorsWanted int) int {
	return int(C.gdImageTrueColorToPalette(i.ptr, C.int(ditherFlag), C.int(colorsWanted)))
}
func (i Image) PaletteToTrueColor() int {
	return int(C.gdImagePaletteToTrueColor(i.ptr))
}

func (i Image) ColorMatch(i2 Image) int {
	return int(C.gdImageColorMatch(i.ptr, i2.ptr))
}
func (i Image) TrueColorToPaletteSetMethod(method, speed int) int {
	return int(C.gdImageTrueColorToPaletteSetMethod(i.ptr, C.int(method), C.int(speed)))
}
func (i Image) TrueColorToPaletteSetQuality(min, max int) {
	C.gdImageTrueColorToPaletteSetQuality(i.ptr, C.int(min), C.int(max))
}
func (i Image) ColorTransparent(color int) {
	C.gdImageColorTransparent(i.ptr, C.int(color))
}
func (i Image) PaletteCopy(dst Image) {
	C.gdImagePaletteCopy(i.ptr, dst.ptr)
}
func (i Image) ColorReplace(src, dst int) int {
	return int(C.gdImageColorReplace(i.ptr, C.int(src), C.int(dst)))
}
func (i Image) ColorReplaceThreshold(src, dst int, threshold float64) int {
	return int(C.gdImageColorReplaceThreshold(i.ptr, C.int(src), C.int(dst), C.float(threshold)))
}
func (i Image) ColorReplaceArray(src, dst []int) int {
	return int(C.gdImageColorReplaceArray(i.ptr, C.int(len(src)),
		(*C.int)(unsafe.Pointer(&src[0])),
		(*C.int)(unsafe.Pointer(&dst[0])),
	))
}
func free(p unsafe.Pointer) {
	C.gdFree(p)
}

func (i Image) GifAnimBegin(out *IOCtx, GlobalCM, loops int) {
	C.gdImageGifAnimBeginCtx(i.ptr, toGdioCtx(out), C.int(GlobalCM), C.int(loops))
}
func (i Image) GifAnimAdd(out *IOCtx, LocalCM, LeftOfs, TopOfs, Delay, Disposal int, previm Image) {
	C.gdImageGifAnimAddCtx(i.ptr, toGdioCtx(out), C.int(LocalCM), C.int(LeftOfs), C.int(TopOfs), C.int(Delay), C.int(Disposal), previm.ptr)
}
func GifAnimEnd(out *IOCtx) {
	C.gdImageGifAnimEndCtx(toGdioCtx(out))
}

func (i Image) FilledArc(cx, cy, w, h, s, e, color, style int) {
	C.gdImageFilledArc(i.ptr, C.int(cx), C.int(cy), C.int(w), C.int(h),
		C.int(s), C.int(e), C.int(color), C.int(style))
}
func (i Image) Arc(cx, cy, w, h, s, e, color int) {
	C.gdImageArc(i.ptr, C.int(cx), C.int(cy), C.int(w), C.int(h),
		C.int(s), C.int(e), C.int(color))
}
func (i Image) Ellipse(cx, cy, w, h, color int) {
	C.gdImageEllipse(i.ptr, C.int(cx), C.int(cy), C.int(w), C.int(h), C.int(color))
}
func (i Image) FilledEllipse(cx, cy, w, h, color int) {
	C.gdImageFilledEllipse(i.ptr, C.int(cx), C.int(cy), C.int(w), C.int(h), C.int(color))
}
func (i Image) FillToBorder(x, y, border, color int) {
	C.gdImageFillToBorder(i.ptr, C.int(x), C.int(y), C.int(border), C.int(color))
}
func (i Image) Fill(x, y, color int) {
	C.gdImageFill(i.ptr, C.int(x), C.int(y), C.int(color))
}
func (i Image) Copy(dst Image, dstx, dsty, srcx, srcy, w, h int) {
	C.gdImageCopy(dst.ptr, i.ptr, C.int(dstx), C.int(dsty),
		C.int(srcx), C.int(srcy), C.int(w), C.int(h))
}
func (i Image) CopyMerge(dst Image, dstx, dsty, srcx, srcy, w, h, pct int) {
	C.gdImageCopyMerge(dst.ptr, i.ptr, C.int(dstx), C.int(dsty),
		C.int(srcx), C.int(srcy), C.int(w), C.int(h), C.int(pct))
}
func (i Image) CopyMergeGray(dst Image, dstx, dsty, srcx, srcy, w, h, pct int) {
	C.gdImageCopyMergeGray(dst.ptr, i.ptr, C.int(dstx), C.int(dsty),
		C.int(srcx), C.int(srcy), C.int(w), C.int(h), C.int(pct))
}
func (i Image) CopyResized(dst Image, dstx, dsty, srcx, srcy, dstw, dsth, srcw, srch int) {
	C.gdImageCopyResized(dst.ptr, i.ptr, C.int(dstx), C.int(dsty), C.int(srcx),
		C.int(srcy), C.int(dstw), C.int(dsth), C.int(srcw), C.int(srch))
}
func (i Image) CopyResampled(dst Image, dstx, dsty, srcx, srcy, dstw, dsth, srcw, srch int) {
	C.gdImageCopyResampled(dst.ptr, i.ptr, C.int(dstx), C.int(dsty), C.int(srcx),
		C.int(srcy), C.int(dstw), C.int(dsth), C.int(srcw), C.int(srch))
}

func (i Image) CopyRotated(dst Image, dstx, dsty float64, srcx, srcy, srcw, srch, angle int) {
	C.gdImageCopyRotated(dst.ptr, i.ptr, C.double(dstx), C.double(dsty), C.int(srcx), C.int(srcy), C.int(srcw), C.int(srch), C.int(angle))
}

func (i Image) Clone() Image {
	return newImage(C.gdImageClone(i.ptr))
}
func (i Image) Destroy() {
	C.gdImageDestroy(i.ptr)
}

func (i Image) SetBrush(brush Image) {
	C.gdImageSetBrush(i.ptr, brush.ptr)
}
func (i Image) SetTile(tile Image) {
	C.gdImageSetTile(i.ptr, tile.ptr)
}
func (i Image) SetAntiAliased(c int) {
	C.gdImageSetAntiAliased(i.ptr, C.int(c))
}
func (i Image) SetAntiAliasedDontBlend(c, dontblend int) {
	C.gdImageSetAntiAliasedDontBlend(i.ptr, C.int(c), C.int(dontblend))
}
func (i Image) SetThickness(thickness int) {
	C.gdImageSetThickness(i.ptr, C.int(thickness))
}
func (i Image) Interlace(interlaceArg int) {
	C.gdImageInterlace(i.ptr, C.int(interlaceArg))
}
func (i Image) AlphaBlending(alphaBlendingArg int) {
	C.gdImageAlphaBlending(i.ptr, C.int(alphaBlendingArg))
}
func (i Image) SaveAlpha(saveAlphaArg int) {
	C.gdImageSaveAlpha(i.ptr, C.int(saveAlphaArg))
}

func (i Image) NeuQuant(maxColor, sampleFactor int) Image {
	return newImage(C.gdImageNeuQuant(i.ptr, C.int(maxColor), C.int(sampleFactor)))
}
func (i Image) Pixelate(blockSize int, mode uint) int {
	return int(C.gdImagePixelate(i.ptr, C.int(blockSize), C.uint(mode)))
}
func (i Image) Scatter(sub, plus int) int {
	return int(C.gdImageScatter(i.ptr, C.int(sub), C.int(plus)))
}
func (i Image) ScatterColor(sub, plus int, colors []int, numcolors uint) int {
	return int(C.gdImageScatterColor(i.ptr, C.int(sub), C.int(plus),
		(*C.int)(unsafe.Pointer(&colors[0])),
		C.uint(len(colors))))
}
func (i Image) Smooth(weight float32) int {
	return int(C.gdImageSmooth(i.ptr, C.float(weight)))
}
func (i Image) MeanRemoval() int {
	return int(C.gdImageMeanRemoval(i.ptr))
}
func (i Image) Emboss() int {
	return int(C.gdImageEmboss(i.ptr))
}
func (i Image) GaussianBlur() int {
	return int(C.gdImageGaussianBlur(i.ptr))
}
func (i Image) EdgeDetectQuick() int {
	return int(C.gdImageEdgeDetectQuick(i.ptr))
}
func (i Image) SelectiveBlur() int {
	return int(C.gdImageSelectiveBlur(i.ptr))
}
func (i Image) Convolution(filter [3][3]float32, filterdiv, offset float32) int {
	return int(C.gdImageConvolution(i.ptr,
		(*[3]C.float)(unsafe.Pointer(&filter[0][0])),
		C.float(filterdiv), C.float(offset)))
}
func (i Image) Color(red, green, blue, alpha int) int {
	return int(C.gdImageColor(i.ptr, C.int(red), C.int(green), C.int(blue), C.int(alpha)))
}
func (i Image) Contrast(contrast float64) int {
	return int(C.gdImageContrast(i.ptr, C.double(contrast)))
}
func (i Image) Brightness(brightness int) int {
	return int(C.gdImageBrightness(i.ptr, C.int(brightness)))
}
func (i Image) GrayScale() int {
	return int(C.gdImageGrayScale(i.ptr))
}
func (i Image) Negate() int {
	return int(C.gdImageNegate(i.ptr))
}
func (i Image) TrueColor() bool {
	return i.ptr.trueColor != 0
}
func (i Image) Size() (int, int) {
	return int(i.ptr.sx), int(i.ptr.sy)
}
func (i Image) ColorsTotal() int {
	return int(i.ptr.colorsTotal)
}
func (i Image) Compare(i2 Image) int {
	return int(C.gdImageCompare(i.ptr, i2.ptr))
}
func (i Image) FlipHorizontal() {
	C.gdImageFlipHorizontal(i.ptr)
}
func (i Image) FlipVertical() {
	C.gdImageFlipVertical(i.ptr)
}
func (i Image) FlipBoth() {
	C.gdImageFlipBoth(i.ptr)
}

type Rect struct {
	r C.gdRect
}

func (i Image) Crop(crop *Rect) Image {
	return newImage(C.gdImageCrop(i.ptr, &crop.r))
}
func (i Image) CropAuto(mode uint) Image {
	return newImage(C.gdImageCropAuto(i.ptr, C.uint(mode)))
}
func (i Image) CropThreshold(color uint, threshold float32) Image {
	return newImage(C.gdImageCropThreshold(i.ptr, C.uint(color), C.float(threshold)))
}
func (i Image) SetInterpolationMethod(id C.gdInterpolationMethod) int {
	return int(C.gdImageSetInterpolationMethod(i.ptr, id))
}
func (i Image) RotateInterpolated(angle float32, bgcolor int) Image {
	return newImage(C.gdImageRotateInterpolated(i.ptr, C.float(angle), C.int(bgcolor)))
}

func AffineApplyToPoint(dst PointF, src PointF, affine [6]float64) int {
	return int(C.gdAffineApplyToPointF(&dst.P, &src.P, (*C.double)(unsafe.Pointer(&affine[0]))))
}
func AffineInvert(dst, src [6]float64) int {
	return int(C.gdAffineInvert((*C.double)(unsafe.Pointer(&dst[0])), (*C.double)(unsafe.Pointer(&src[0]))))
}
func AffineFlip(dst, src [6]float64, fliph, flipv int) int {
	return int(C.gdAffineFlip((*C.double)(unsafe.Pointer(&dst[0])), (*C.double)(unsafe.Pointer(&src[0])), C.int(fliph), C.int(flipv)))
}
func AffineConcat(dst, m1, m2 [6]float64) int {
	return int(C.gdAffineConcat((*C.double)(unsafe.Pointer(&dst[0])),
		(*C.double)(unsafe.Pointer(&m1[0])),
		(*C.double)(unsafe.Pointer(&m2[0]))))
}
func AffineIdentity(dst [6]float64) int {
	return int(C.gdAffineIdentity((*C.double)(unsafe.Pointer(&dst[0]))))
}
func AffineScale(dst [6]float64, scalex, scaley float64) int {
	return int(C.gdAffineScale((*C.double)(unsafe.Pointer(&dst[0])), C.double(scalex), C.double(scaley)))
}
func AffineRotate(dst [6]float64, angle float64) int {
	return int(C.gdAffineRotate((*C.double)(unsafe.Pointer(&dst[0])), C.double(angle)))
}
func AffineShearHorizontal(dst [6]float64, angle float64) int {
	return int(C.gdAffineShearHorizontal((*C.double)(unsafe.Pointer(&dst[0])), C.double(angle)))
}
func AffineShearVertical(dst [6]float64, angle float64) int {
	return int(C.gdAffineShearVertical((*C.double)(unsafe.Pointer(&dst[0])), C.double(angle)))
}
func AffineTranslate(dst [6]float64, offsetx, offsety float64) int {
	return int(C.gdAffineTranslate((*C.double)(unsafe.Pointer(&dst[0])), C.double(offsetx), C.double(offsety)))
}
func AffineExpanstion(dst [6]float64) int {
	return int(C.gdAffineExpansion((*C.double)(unsafe.Pointer(&dst[0]))))
}
func AffineRectilinear(dst [6]float64) int {
	return int(C.gdAffineRectilinear((*C.double)(unsafe.Pointer(&dst[0]))))
}
func AffineEqual(dst, src [6]float64) int {
	return int(C.gdAffineEqual((*C.double)(unsafe.Pointer(&dst[0])), (*C.double)(unsafe.Pointer(&src[0]))))
}
func TransformAffineGetImage(dst, src Image, srcArea Rect, affine [6]float64) int {
	return int(C.gdTransformAffineGetImage(&dst.ptr, src.ptr, &srcArea.r, (*C.double)(unsafe.Pointer(&affine))))
}
func TransformAffineCopy(dst Image, dstx, dsty int, src Image, region Rect, affine [6]float64) int {
	return int(C.gdTransformAffineCopy(dst.ptr, C.int(dstx), C.int(dsty), src.ptr, &region.r,
		(*C.double)(unsafe.Pointer(&affine[0]))))
}
func TransformAffineBoundingBox(src Rect, affine [6]float64, bbox Rect) int {
	return int(C.gdTransformAffineBoundingBox(&src.r, (*C.double)(unsafe.Pointer(&affine[0])), &bbox.r))
}
