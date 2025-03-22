package main

import (
	"fmt"
	"os"
	"os/user"
	"umbuLang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s. This is the umbuLang Programming Language.\n", user.Username)
	fmt.Print("Type in any command.\n")
	repl.Start(os.Stdin, os.Stdout)
}
