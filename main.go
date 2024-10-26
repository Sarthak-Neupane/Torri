package main

import (
	"fmt"
	"os"
	"os/user"
	"torri/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Yoooooo %s! This is the new torri language!\n",
		user.Username)
	fmt.Printf("Just type anything\n")
	repl.Start(os.Stdin, os.Stdout)
}
