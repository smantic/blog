package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type HelloHandler struct {
	data string
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, h.data)
}

func main() {

	var data string = "Hello, World!\n"

	if len(os.Args) == 2 {
		data = os.Args[1]
	}

	h := &HelloHandler{
		data: data,
	}

	s := http.NewServeMux()
	s.Handle("/foo/bar", h)

	err := http.ListenAndServe(":8000", s)
	if err != nil {
		log.Println(err)
	}
}
