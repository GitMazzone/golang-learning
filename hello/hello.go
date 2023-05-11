// Package consists of all files in same dir
package main

// Find packages at https://pkg.go.dev/
import (
	"fmt"
	"log"

	"golang-learning/greetings"
)

// Run with go run .
func main() {
	log.SetPrefix("Error @ golang-learning/greetings: ")
	log.SetFlags(0)

	message, error := greetings.Hello("Mike")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
