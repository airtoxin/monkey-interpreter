package main

import (
	"fmt"
	"os"

	"github.com/hackebrot/turtle"
)

func main() {
	name := "wolf"
	emoji, ok := turtle.Emojis[name]

	if !ok {
		fmt.Fprintf(os.Stderr, "no emoji found for name: %v\n", name)
		os.Exit(1)
	}

	fmt.Printf("Name: %q\n", emoji.Name)
	fmt.Printf("Char: %s\n", emoji.Char)
	fmt.Printf("Category: %q\n", emoji.Category)
	fmt.Printf("Keywords: %q\n", emoji.Keywords)
}
