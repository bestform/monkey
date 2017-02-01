package main

import (
    "os/user"
    "fmt"
    "github.com/bestform/monkey/repl"
    "os"
)

func main() {
    activeUser, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hello %s! This is the Monkey programming language!\n", activeUser.Username)
    fmt.Println("Feel free to type in commands")
    repl.Start(os.Stdin, os.Stdout)
}


