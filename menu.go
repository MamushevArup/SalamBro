package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s []string
var menuMap = make(map[string]int, len(s))

func displayMenu() {
	if len(s) == 0 {
		fmt.Println("\nYou don't have menu yet!")
		return
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v. %v\n", i+1, s[i])
		menuMap[s[i]] = i + 1
	}
}
func changeMenu() []string {
	fmt.Println("\nHere you can do the following activities:\n0. Go Back.\n1. Add categories.\n2. Remove categories.")
	var choice int
	fmt.Print("\nYour choice: ")
	fmt.Scan(&choice)
	if choice == 1 {
		fmt.Print("\nWrite the names of categories separated by comma (c1, c2, ...);\n>")
		scanner := bufio.NewReader(os.Stdin)
		read, _ := scanner.ReadString('.')
		read = strings.TrimLeft(read, "\n")
		s1 := strings.Split(read[:len(read)-1], ", ")
		for i := 0; i < len(s1); i++ {
			s = append(s, s1[i])
		}
	} else if choice == 2 {
		fmt.Print("\nChoose what category you want remove using their number and separated by comma (c1, c2, ...)\n> ")
		scanner := bufio.NewReader(os.Stdin)
		line, _ := scanner.ReadString('.')
		line = strings.TrimLeft(line, "\n")
		rmvStr := strings.Split(line[:len(line)-1], ",")
		rmvInt := make([]int, len(rmvStr))
		for i, str := range rmvStr {
			rmvInt[i], _ = strconv.Atoi(str)
		}
		for i := 0; i < len(rmvInt); i++ {
			for j := 0; j <= len(s); j++ {
				if rmvInt[i] < len(s) {
					if rmvInt[i] == j {
						temp := s[j-1]
						s[j-1] = s[len(s)-1]
						s[len(s)-1] = temp
						s = s[:len(s)-1]
						delete(menuMap, temp)
					}
				}
			}
		}
	}
	return s
}
