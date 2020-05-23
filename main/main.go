package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/", hello)
	_ = http.ListenAndServe(":80", nil)
}
