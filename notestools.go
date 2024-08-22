package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		printHelp()
		return
	}

	collectionName := os.Args[1]

	if collectionName == "help" {
		printHelp()
		return
	}

	notes := LoadOrCreateCollection(collectionName)
	clearConsole()
	greeting()
	//Greeting

	//fmt.Println("Welcome to the notes tool!")
	PrintMenu()
	for {
		//displayNotes()

		//fmt.Print("Choose an option: ")
		//printMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1: //show
			clearConsole()
			greeting()
			displayNotes(notes)
			PrintMenu()
		case 2: //add
			e := addNote(&notes, collectionName)
			//printMenu()
			clearConsole()
			greeting()
			switch e {
			case 0:
				fmt.Println(colorGreen + "\nThe notes have been added" + colorReset)
			case 1:
				fmt.Println(colorRed + "Empty input. The notes have not been added." + colorReset)

			}

			PrintMenu()
		case 3: //remove note
			//e := removeNote(&notes, collectionName)
			clearConsole()
			greeting()
			displayNotes(notes)
			e := removeNote(&notes, collectionName)
			clearConsole()
			greeting()
			switch e {
			case 0:
				fmt.Println(colorGreen + "\nThe notes have been removed" + colorReset)
			case 1:
				fmt.Println(colorGreen + "\nNo notes to remove." + colorReset)
			case 2:
				fmt.Println(colorGreen + "\nNote removal canceled." + colorReset)
			}

			PrintMenu()
		case 4:
			clearConsole()
			greeting()
			displayNotes(notes)
			e := editNote(&notes, collectionName)
			clearConsole()
			greeting()
			switch e {
			case 0:
				fmt.Println(colorGreen + "\nThe notes have been edited" + colorReset)
			case 1:
				fmt.Println(colorGreen + "\nNo notes to edit." + colorReset)
			case 2:
				fmt.Println(colorGreen + "\nNote editing canceled." + colorReset)
			default:
				fmt.Println(colorRed + "\nUnexpected error." + colorReset)
			}
			PrintMenu()
		case 5: //remove collection
			e := removeCollection(collectionName)
			clearConsole()
			greeting()
			switch e {
			case 0:
				fmt.Println(colorGreen + "\nThe collection has been removed" + colorReset)
				fmt.Println("Press Enter to exit...")
				fmt.Scanln()
				clearConsole()
				return
			case 1:
				fmt.Println(colorGreen + "\nCollection removal canceled." + colorReset)
				PrintMenu()
			}
		case 6: //exit
			SaveCollection(notes, collectionName)
			clearConsole()
			//fmt.Println("Exiting...")
			return
		default:
			fmt.Println(colorRed + "Invalid option. Please choose again." + colorReset)
		}
	}

}
