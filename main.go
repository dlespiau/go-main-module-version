package main

import (
	"fmt"
	"os"
)

func main() {
	if err := Print(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
