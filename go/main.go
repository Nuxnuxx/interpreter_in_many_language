package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Nuxnuxx Programming language!\n", user.Username)

	fmt.Printf("Type some command\n")

	repl.Start(os.Stdin, os.Stdout)
}
