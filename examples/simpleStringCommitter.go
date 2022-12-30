package main

import (
	"fmt"
	"os"

	"github.com/muryoutaisuu/gitter"
)

func main() {
	if len(os.Args) != 8 {
		fmt.Printf("Error: Need 7 arguments: <git repo url> <token> <filename> <file content> <commit message> <commit author name> <commit author email>\n")
		return
	}
	var url = os.Args[1]
	var token = os.Args[2]
	var filename = os.Args[3]
	var content = os.Args[4]
	var message = os.Args[5]
	var name = os.Args[6]
	var email = os.Args[7]

	cs := gitter.CreateCommitSignature(name, email)
	c, err := gitter.New(url, token, cs)
	if err != nil {
		fmt.Printf("Got an Error initiating Client, error was: %s\n", err)
		return
	}

	err = c.CommitNewFile(filename, message,
		[]byte(fmt.Sprintf(content)))
	if err != nil {
		fmt.Printf("Got an Error commit new file, error was: %s\n", err)
		return
	}
}
