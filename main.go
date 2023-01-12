package main

import (
	"fmt"
)

func main() {
	welcome := "Welcome to the Admin Panel of 'Salam Bro' Fast Food Station\nHere you can do the following activities\n"
	entireMenu := "\n0. Exit\n1. Change Welcome message\n2. Change Menu\n3. Change SubMenu\n4. Change Price\n5. Display Welcome message\n6. Display Menu\n7. Display SubMenu\n8. Display Prices\n\n"
	var choice int
	fmt.Print(welcome, entireMenu)
	fmt.Print("Choose a activity: ")
	fmt.Scan(&choice)
	for choice != 0 {
		if choice == 1 {
			changeWelcome(&welcome)
		} else if choice == 5 {
			displayWelcome(welcome)
		} else if choice == 6 {
			displayMenu()
		} else if choice == 2 {
			changeMenu()
		} else if choice == 7 {
			displaySubMenu()
		} else if choice == 3 {
			changeSubMenu()
		} else if choice == 8 {
			displayPrice()
		} else if choice == 4 {
			changePrice()
		}

		fmt.Print(entireMenu)
		fmt.Print("Choose a activity: ")
		fmt.Scan(&choice)
	}
}
