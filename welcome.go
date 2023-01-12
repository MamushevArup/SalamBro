package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func changeWelcome(s *string) *string {
	fmt.Print("\n", "Please, type the message that clients will see when they enter the menu:\n",
		"Tip: use underscore (_) as a splitter\n")
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	// Read a line from the user.
	text, _ := reader.ReadString('!')
	text = strings.ReplaceAll(text, "_ ", "\n")
	text = text[:len(text)-1] + "\n"
	*s = text
	return s
}
func displayWelcome(s string) {
	fmt.Print(s)
}
