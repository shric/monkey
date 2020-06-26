package main

import (
	"fmt"
	"os/user"

	"github.com/shric/monkey/object"

	"github.com/shric/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	env := object.NewEnvironment()
	repl.Start(env)
}
