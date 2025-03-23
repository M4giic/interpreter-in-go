package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Jedreks programming langugae!\n", user.Username)
	fmt.Printf("Rob ta co chce ta\n")
	repl.Start(os.Stdin, os.Stdout)
}
