package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kaputi/sindar/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Sindar programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")

	// repl.StartParser(os.Stdin, os.Stdout)
	// repl.StartLexer(os.Stdin, os.Stdout)
	repl.Start(os.Stdin, os.Stdout)
}
