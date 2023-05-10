package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func SequentialFetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		res, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch error: %s", err)
			os.Exit(1)
		}

		io.Copy(os.Stdout, res.Body)

		// body, err := ioutil.ReadAll(res.Body)
		// res.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch error reading: %s: %s", url, err)
			os.Exit(1)
		}

		// fmt.Printf("%s", body)
		fmt.Printf("Response when fetching URL %s: %s\n", url, res.Status)
	}
}
