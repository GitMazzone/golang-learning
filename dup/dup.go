package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	happensIn := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, happensIn)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file in dup: %v\n", err)
				continue
			}

			countLines(file, counts, happensIn)
			file.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, happensIn[line])
		}
	}
}

func countLines(file *os.File, counts map[string]int, happensIn map[string]string) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		line := input.Text()

		counts[line]++

		if !strings.Contains(happensIn[line], file.Name()) {
			happensIn[line] += file.Name() + " "
		}
	}

	// Ideally handle possible read errors from input.Error()
}
