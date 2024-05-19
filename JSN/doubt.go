package JSN

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// no one of this func, wrote some usefull things inside this func
// like getting author info, commit msg
func dummy() {
	// Check Other Files for helper functions and boiler plate
	// You copy and use them in your own code

	// Code to get author info and commit message (the easy way for now)

	// environment variable are defined in shell
	// set environment in your shell by
	// Example: export AUTHOR_NAME=your_name
	// Execute echo $AUTHOR_NAME to check if the variable has been set in the shell
	// Author info from environment variable
	name := os.Getenv("AUTHOR_NAME")
	email := os.Getenv("AUTHOR_EMAIL")

	// commit message from stdin
	// This is different from just Scanning input
	// Scanning Input is quite limited and breaks with new line character
	// ReadAll will read os.Stdin till EOF
	// so you can cat FILE.txt | ./X commit to pipeline any kind of commit msg to your vcs
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		// panic will end the program
		panic(err)
	}
	// typeCasting []byte to string
	m := string(stdin)
	// Using this code you can pipeline Commit message in binary
	// Example: echo "First Commit" | ./X commit

	fmt.Println(name, email, m)

	// whatever your root commit ID is
	var commitID string

	// This two lines will print output of your commit command like git
	fmt.Printf("[(root-commit) %s] ", commitID)
	fmt.Printf("%s\n", strings.Split(m, "\n")[0])
}
