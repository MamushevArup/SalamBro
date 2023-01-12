package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var hashMap = make(map[string][]string)

func displaySubMenu() {
	fmt.Println("\nChoose the category for which you want to see its subcategories:")
	displayMenu()
	var choice int
	fmt.Print("\n> ")
	fmt.Scan(&choice)
	// choice = index +1 s = array with the menu
	if hashMap[s[choice-1]] == nil {
		fmt.Println("You don't have any subcategories")
		return
	}
	for i := 0; i < len(hashMap[s[choice-1]]); i++ {
		fmt.Printf("%v. %v\n", i+1, hashMap[s[choice-1]][i])
	}

}
func changeSubMenu() {
	fmt.Println("\nHere you can do the following activities:\n0. Go Back.\n1. Add categories.")
	var choice int
	fmt.Print("\nYour choice: ")
	fmt.Scan(&choice)
	if choice == 1 {
		fmt.Print("\nChoose a category for which you want to add subcategories\n")
		displayMenu()
		var sub int
		fmt.Print("\n> ")
		fmt.Scan(&sub)
		fmt.Print("\nWrite the names of subcategories separated by comma (c1, c2, ...)\n> ")
		scanner := bufio.NewReader(os.Stdin)
		read, _ := scanner.ReadString('.')
		read = strings.TrimLeft(read, "\n")
		ss := strings.Split(read[:len(read)-1], ", ")
		for i := 0; i < len(s); i++ {
			if sub-1 == i {
				hashMap[s[i]] = ss
			}
		}
	}
}
