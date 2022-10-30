package main

import (
	"fmt"
	"os"

	fake "github.com/brianvoe/gofakeit/v6"
)

var version = ""
var osArch = ""

func main() {
	if len(os.Args) < 2 {
		fmt.Println(fake.BeerName())
		return
	}

	if os.Args[1] == "version" {
		fmt.Printf("rand-beer version %s %s\n", version, osArch)
		return
	}

	os.Exit(1)
}
