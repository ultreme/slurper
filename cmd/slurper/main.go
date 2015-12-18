package main

import (
	"fmt"
	"os"
	"github.com/ultreme/slurper"
)

// main is the entrypoint
func main() {
	s := slurper.NewSlurper()
	inputs := os.Args[1:]

	for i := range(inputs) {
		fmt.Printf(s.Slurp(inputs[i]))
	}

	if len(inputs) != 0 {
		fmt.Printf("\n")
	}
}
