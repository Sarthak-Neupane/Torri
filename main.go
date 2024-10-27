package main

import (
	"fmt"
    "os"
    "os/user"
    "torri/repl"
    // "io"
)

func main() {
	// Get current user for personalized greeting
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    
    if len(os.Args) > 1 {
        // If a command is given, check if it's "run" and execute file
        command := os.Args[1]
        if command == "run" && len(os.Args) == 3 {
            fileName := os.Args[2]
            runFile(fileName)
            return
        } else {
            fmt.Println("Usage: torri run <file>")
            os.Exit(1)
        }
    }

    // Start REPL if no arguments are given
    fmt.Printf("Yoooooo %s! This is the new Torri language!\n", user.Username)
    fmt.Println("Just type anything")
    repl.Start(os.Stdin, os.Stdout)
}

func runFile(fileName string) {
    // Read the file
    content, err := os.ReadFile(fileName)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }

    repl.Execute(string(content), os.Stdout)
}
