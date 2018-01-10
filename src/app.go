package main

import (
	"net/http"
	//"io/ioutil"
	"fmt"
	//"github.com/pkg/profile"
	"io"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100; i++ {
		fileReader, err := os.Open(fmt.Sprintf("./data/list%d.json", (i % 3) + 1))
		if err != nil {
			panic(err)
			return
		}

		_, err = io.Copy(w, fileReader)
		if err != nil {
			panic(err)
			return
		}
	}
}

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}