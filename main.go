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
	fmt.Printf("Hello %s! Welcome to Monkey v0.1.0\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
