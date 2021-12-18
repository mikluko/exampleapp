package main

import (
	"fmt"
	"os"

	"github.com/mikluko/examplelib"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("args required")
		os.Exit(1)
	}
	for i := range args {
		fmt.Println(args[i], examplelib.Reverse(args[i]))
	}
}
