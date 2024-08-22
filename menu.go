package main

import "fmt"

func PrintMenu() {
	fmt.Println(colorCyan + "\nSelect operation:")
	fmt.Println("1. Show notes")
	fmt.Println("2. Add a note")
	fmt.Println("3. Delete a note")
	fmt.Println("4. Edit a note")
	fmt.Println("5. Delete the collection")
	fmt.Println("6. Exit" + colorReset)
}
