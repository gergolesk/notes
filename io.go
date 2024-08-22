package main

import (
	"bufio"
	"fmt"
	"os"
)

// Loading or creating collection
func LoadOrCreateCollection(collectionName string) []string {
	//return 'file' object for reading the contents, if file doesnt exist,
	//err will capture the error
	file, err := os.Open(collectionName)
	//check if the error is because the file doesnt exist. if file doesnt exist yet,
	//then an emtpy collection is created (but not saved yet)
	if os.IsNotExist(err) {
		return []string{}
		//error exists, file also exists, print error and exit program
	} else if err != nil {
		fmt.Println("Error loading collection:", err)
		os.Exit(1)
	}
	//ensure file properly closed, whether function completed successfully or not
	defer file.Close()
	//declare notes variable (a slice of strings), bufio scanner reads text from 'file'
	var notes []string
	scanner := bufio.NewScanner(file)
	//iterate over each line in the file, scanner.Scan returns the scanned line of text
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}
	// same case for errors as before
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading collection", err)
		os.Exit(1)
	}
	return notes //returns all the lines we scanned form the file in a notes slice
}

// Saving
func SaveCollection(notes []string, collectionName string) {
	// Create file or display error if err value is not nil, incase of error exit the program
	file, err := os.Create(collectionName)
	if err != nil {
		fmt.Println(colorRed+"Error saving collection:", err, colorReset)
		os.Exit(1)
	}
	// defer makes sure the file gets closed, even if an error occurs later in function
	defer file.Close()
	//iterate over each element in notes slice
	for _, note := range notes {
		_, err := file.WriteString(note + "\n")
		//check for errors, if err value is not nil then display error
		if err != nil {
			fmt.Println(colorRed, "Error writing to file:", err, colorReset)
			os.Exit(1)
		}
	}

	//fmt.Println("Collection saved.")
}

func removeCollection(collectionName string) int {
	// Return codes:
	// 0 - success
	// 1 - canceled by user

	// Display a warning message to the user about the irreversibility of deleting a collection
	fmt.Println(colorYellow + "Warning! A deleted collection cannot be restored! Are you sure? (y/n)" + colorReset)

	// Variable to store the user's decision
	decision := ""

	// Loop until the user provides a valid input (either "y"/"yes" or "n"/"no")
	for {
		// Read user input
		fmt.Scan(&decision)

		// If the user confirms (input is "y" or "yes"), remove the collection and return 0 (success)
		if decision == "y" || decision == "yes" {
			os.Remove(collectionName)
			return 0
		} else if decision == "n" || decision == "no" {
			// If the user cancels (input is "n" or "no"), return 1 (canceled)
			return 1
		}
		// Loop will repeat if input is invalid
	}
}
