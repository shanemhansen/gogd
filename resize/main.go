package main

import (
	"flag"
	"github.com/davecheney/profile"
	"github.com/shanemhansen/gogd"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":9090", "address to listen on")
var newimg = gogd.ImageCreateTrueColor(64, 64)

func handle(w http.ResponseWriter, r *http.Request) {
	out, err := os.Open("bot.png")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	orig := gogd.ImageCreateFromPngCtx(out)
	gogd.ImageCopyResized(newimg, orig, 0, 0, 0, 0, 64, 64, 128, 128)
	gogd.ImagePngCtx(newimg, w)
	gogd.ImageDestroy(orig)
}
func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	http.HandleFunc("/", handle)
	log.Println("listening on ", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
