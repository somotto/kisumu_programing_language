package main

import (
	"fmt"
	"kisumu/repl"
	"os"
)

func main() {
    fmt.Println("Welcome to Kisumu Programming Language!")
    repl.Start(os.Stdin, os.Stdout)
}
