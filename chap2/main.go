package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kassy11/monkey-interpreter/chap2/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the monkey Programming Language\n", user.Username)
	fmt.Printf("Feel free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
