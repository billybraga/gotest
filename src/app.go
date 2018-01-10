package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	//"github.com/pkg/profile"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100; i++ {
		file, err := ioutil.ReadFile(fmt.Sprintf("./data/list%d.json", (i % 3) + 1))
		if err != nil {
			panic(err)
			return
		}

		w.Write(file)
	}
}

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}