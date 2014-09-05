package main

import (
	"flag"
	"github.com/davecheney/profile"
	"github.com/shanemhansen/gogd"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":9090", "address to listen on")
var newimg = gogd.ImageCreateTrueColor(64, 64).Png()

func handle1(w http.ResponseWriter, r *http.Request) {
	in, err := os.Open("bot.png")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	img := gogd.NewPng()
	if err := img.Read(in); err != nil {
		log.Println(err)
		return
	}
	defer img.Destroy()
	img.CopyResampled(newimg, 0, 0, 0, 0, 64, 64, 128, 128)
	if err := newimg.Write(w); err != nil {
		log.Println(err)
		return
	}
}
func handle2(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadFile("bot.png")
	if err != nil {
		log.Println(err)
		return
	}
	img := gogd.NewPng()
	if err := img.FromBuffer(buf); err != nil {
		log.Println(err)
		return
	}
	defer img.Destroy()
	img.CopyResampled(newimg, 0, 0, 0, 0, 64, 64, 128, 128)
	bufout, err := newimg.ToBuffer()
	if err != nil {
		log.Println(err)
	}
	w.Write(bufout)
}
func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	http.HandleFunc("/1", handle1)
	http.HandleFunc("/2", handle2)
	log.Println("listening on ", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
