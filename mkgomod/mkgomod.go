package main

import (
	"fmt"
	"os"
)

func main() {
	modName := "../" + os.Args[1]

	err := os.Mkdir(modName, 0777)

	if err != nil {
		fmt.Fprintf(os.Stderr, "mkdir: %s\n", err)
		os.Exit(1)
	}
}
