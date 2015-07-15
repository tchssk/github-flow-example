package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func helloWorld(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func helloGolang(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, golang!")
}

func main() {
	goji.Get("/", helloWorld)
	goji.Get("/golang", helloGolang)
	goji.Serve()
}
