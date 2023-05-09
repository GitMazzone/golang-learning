package main

import (
	"fmt"
	"os"
)

func RangedEcho() {
	fmt.Printf("Invoked by: %v \n", os.Args[0])

	for i, arg := range os.Args[1:] {
		fmt.Printf("Arg #%v: %v\n", i+1, arg)
	}
}
