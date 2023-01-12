package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var res = make(map[string]int)

func displayPrice() {
	if len(s) == 0 {
		fmt.Println("You don't have any menu")
		return
	}
	fmt.Println("\nChoose the category for which you want to see its subcategories prices:")
	displayMenu()
	var choice int
	fmt.Print("\n> ")
	fmt.Scan(&choice)
	if len(hashMap[s[choice-1]]) == 0 {
		fmt.Println("\nYou don't have any subcategories")
		return
	}
	for i := 0; i < len(hashMap[s[choice-1]]); i++ {
		fmt.Printf("%d. %v - %d kzt.\n", i+1, hashMap[s[choice-1]][i], res[hashMap[s[choice-1]][i]])
	}
}
func changePrice() {
	fmt.Println("\nChoose a category for which you want to change prices of subcategories: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d. %v\n", i+1, s[i])
	}
	var men int
	fmt.Print("\n> ")
	fmt.Scan(&men)
	if len(hashMap[s[men-1]]) == 0 {
		fmt.Println("You don't have any subcategories")
		return
	}
	fmt.Println("Choose subcategories for which you want to change prices")
	for i := 0; i < len(hashMap[s[men-1]]); i++ {
		fmt.Printf("%d. %v - %d kzt.\n", i+1, hashMap[s[men-1]][i], res[hashMap[s[men-1]][i]])
	}
	fmt.Print("Tip: write the indexes separated by comm (c1, c2, ...) \n>")
	scanner := bufio.NewReader(os.Stdin)
	line, _ := scanner.ReadString('.')
	line = strings.TrimLeft(line, "\n")
	rmvStr := strings.Split(line[:len(line)-1], ",")
	rmvInt := make([]int, len(rmvStr))
	for i, str := range rmvStr {
		rmvInt[i], _ = strconv.Atoi(str)
	}
	for i := 0; i < len(rmvInt); i++ {
		if rmvInt[i]-1 > len(hashMap[s[men-1]]) {
			rmvInt = append(rmvInt[:i], rmvInt[i+1:]...)
		}
	}
	if rmvInt[len(rmvInt)-1] > len(hashMap[s[men-1]]) {
		rmvInt = rmvInt[:len(rmvInt)-1]
	}
	var price int
	for i := 0; i < len(rmvInt); i++ {
		fmt.Printf("Set price for %v : ", hashMap[s[men-1]][rmvInt[i]-1])
		fmt.Scan(&price)
		res[hashMap[s[men-1]][rmvInt[i]-1]] = price
	}
}
