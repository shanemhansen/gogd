Gogd
----

Gogd is a set of bindings to the libgd library. Gogd deeply integrates
go and libgd via IOCtx objects, which allow libgd to unterstand golang
io.Readers and io.Writers.

Supported formats include:

* png
* jpeg
* gif
* webp
* WBMP
* bmp

All gd apis are supported including:

* resizing
* get/set pixel
* affine transforms
* line drawing
* clipping
* true color images
* text drawing
* Animated gifs
* Gaussian/Edge Detection/Scatter, etc.
