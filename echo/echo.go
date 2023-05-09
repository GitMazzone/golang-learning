package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	RangedEcho()
	rangedEchoSecs := time.Since(start).Milliseconds()
	fmt.Printf("Runtime of RangedEcho: %.5dms", rangedEchoSecs)
}
