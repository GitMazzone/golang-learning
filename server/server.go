package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", showCount)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	count++
	mutex.Unlock()

	fmt.Fprintf(writer, "URL.pathypoo = %q\n", req.URL.Path)
}

func showCount(writer http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	fmt.Fprintf(writer, "Count: %d\n", count)
	mutex.Unlock()
}
