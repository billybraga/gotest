package main

import (
	"github.com/valyala/fasthttp"
	//"io/ioutil"
	"fmt"
	//"github.com/pkg/profile"
	//"io"
	//"os"
	"io/ioutil"
)

type FileManager struct {
	files [3][]byte
}

func NewFileManager() *FileManager {
	f := new(FileManager)

	for i := 0; i < 3; i++ {
		content, err := ioutil.ReadFile(fmt.Sprintf("./data/list%d.json", (i%3)+1))
		if err != nil {
			panic(err)
			return f
		}
		f.files[i] = content
	}

	return f
}

type Handler struct {
	fileManager *FileManager
}

func NewHandler() *Handler {
	h := new(Handler)
	h.fileManager = NewFileManager()
	return h
}

func (h *Handler) handle(ctx *fasthttp.RequestCtx) {
	defer ctx.SetConnectionClose()

	for i := 0; i < 100; i++ {
		ctx.Response.AppendBody(h.fileManager.files[i % 3])
	}
}

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()
	handler := NewHandler()
	fasthttp.ListenAndServe(":8081", handler.handle)
}