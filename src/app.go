package main

import (
	"net/http"
	"html/template"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/nytimes/gziphandler"
)

type Item struct {
	Name string
	Description string
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("views/home.html")
	if e != nil {
		w.Write([]byte(e.Error()))
		return
	}
	if t.Tree == nil || t.Tree.Root == nil {
		w.Write([]byte("Template null"))
		return
	}
	l := len(t.Tree.Root.Nodes)
	if l <= 0 {
		w.Write([]byte("Empty template"))
		return
	}

	var data [100][]Item

	for i := 0; i < 100; i++ {
		file, err := ioutil.ReadFile(fmt.Sprintf("./data/list%d.json", (i % 3) + 1))
		if err != nil {
			panic(err)
			return
		}

		json.Unmarshal(file, &data[i])
	}

	e = t.Execute(w, data)
	if e != nil {
		w.Write([]byte(e.Error()))
		return
	}
}

func main() {
	http.Handle("/", gziphandler.GzipHandler(http.HandlerFunc(handler)))
	http.ListenAndServe(":8081", nil)
}