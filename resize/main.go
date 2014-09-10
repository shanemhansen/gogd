package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"fmt"
	"github.com/shanemhansen/gogd"
	"strings"
)

var addr = flag.String("addr", ":9090", "address to listen on")

func handle(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	src := params.Get("url")
	if src == "" {
		http.Error(w, "url param is required", http.StatusBadRequest)
		return
	}
	resizeratio, err := strconv.ParseFloat(params.Get("ratio"), 64)
	if err != nil {
		http.Error(w, "bad ratio, please give a number between 0 and 1", http.StatusBadRequest)
		return
	}
	resp, err := http.Get(src)
	if err != nil {
		http.Error(w, "unable to get url", http.StatusBadRequest)
		return
	}
	var imageio gogd.ImageWriter
	switch strings.ToLower(resp.Header.Get("Content-Type")) {
	case "image/png":
		imageio = new(gogd.PngIO)
	case "image/jpeg":
		imageio = &gogd.JpegIO{Quality: 90}
	case "image/gif":
		imageio = new(gogd.GifIO)
	case "image/webp":
		imageio = &gogd.WebpIO{Quantization: 90}
	default:
		http.Error(w, fmt.Sprintf("unsupported image type %s", resp.Header.Get("Content-Type")), http.StatusBadRequest)
		return
	}
	// load image
	srcimg := imageio.Decode(resp.Body)
	if !srcimg.Valid() {
		http.Error(w, "Unknown error decoding image", http.StatusBadRequest)
	}
	defer srcimg.Destroy()
	width, height := srcimg.Size()
	newWidth := int(resizeratio * float64(width))
	newHeight := int(resizeratio * float64(height))
	// resample
	dstimg := gogd.ImageCreateTrueColor(newWidth, newHeight)
	defer dstimg.Destroy()
	srcimg.CopyResampled(dstimg, 0, 0, 0, 0, newWidth, newHeight, width, height)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	// optionally add some text
	msg := params.Get("msg")
	if msg != "" {
		font := gogd.FontGetSmall()
		dstimg.String(font, newWidth/10, newHeight/10, msg, dstimg.ColorAllocate(0, 0, 0))
	}

	imageio.Encode(dstimg, w)
}
func main() {
	flag.Parse()
	http.HandleFunc("/", handle)
	log.Println("listening on ", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
