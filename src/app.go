package main

import (
	"github.com/valyala/fasthttp"
	//"io/ioutil"
	"fmt"
	//"github.com/pkg/profile"
	"io"
	"os"
)

func handler(ctx *fasthttp.RequestCtx) {
	defer ctx.SetConnectionClose()

	for i := 0; i < 100; i++ {
		fileReader, err := os.Open(fmt.Sprintf("./data/list%d.json", (i % 3) + 1))
		if err != nil {
			panic(err)
			return
		}
		defer fileReader.Close()

		_, err = io.Copy(ctx.Response.BodyWriter(), fileReader)
		if err != nil {
			panic(err)
			return
		}
	}
}

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()
	fasthttp.ListenAndServe(":8081", handler)
}