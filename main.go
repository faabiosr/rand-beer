// Package main
package main

import (
	"fmt"
	"os"

	fake "github.com/brianvoe/gofakeit/v6"
)

const unknown = "unknown"

// variables are expected to be set at build time.
var (
	version = unknown
	osArch  = unknown
)

func main() {
	const nArgs = 2
	if len(os.Args) < nArgs {
		fmt.Println(fake.BeerName())
		return
	}

	if os.Args[1] == "version" {
		fmt.Printf("rand-beer version %s %s\n", version, osArch)
		return
	}

	os.Exit(1)
}
