package main

import (
	"fmt"
	"os"
	"os/user"
	repl "github.com/RootControl/Interpreter/internal/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Hello %s! This is the Codex programming language!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
