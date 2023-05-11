package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "URL.pathypoo = %q\n", req.URL.Path)
}
