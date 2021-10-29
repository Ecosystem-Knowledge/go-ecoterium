package main

import (
	"fmt"
	"os"

	"github.com/Ecosystem-Knowledge/go-ecoterium/cmd/opera/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
