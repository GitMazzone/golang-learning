package main

import (
	"fmt"
	"os"
)

func main() {
	modName := os.Args[1]
	dirName := "../" + modName

	err := os.Mkdir(dirName, 0777)

	if err != nil {
		fmt.Fprintf(os.Stderr, "mkdir: %s\n", err)
		os.Exit(1)
	}

	file, err := os.Create(dirName + "/" + modName + ".go")

	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Create: %s\n", err)
		os.Exit(1)
	}

	file.WriteString(`package main

import (
	"fmt"
)

func main() {
	fmt.Println("New module")
}
	`)

	defer file.Close()
}
