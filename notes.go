package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Displaying notes
func displayNotes(notes []string) {
	// Print the header with cyan color
	fmt.Printf(colorCyan + "\nNotes:\n" + colorReset)

	// Check if there are no notes in the collection
	if len(notes) == 0 {
		// Print a message in yellow color if the collection is empty
		fmt.Println(colorYellow + "No notes in the collection." + colorReset)
		return
	}

	// Iterate over the notes slice
	for i, note := range notes {
		// Print each note with its index (starting from 1), and format it with leading zeros
		// Each note is printed with the default terminal color (colorReset)
		fmt.Printf("%v%03d - %s\n", colorReset, i+1, note)
	}
}

// Adding notes
func addNote(notes *[]string, collectionName string) int {
	//Return codes
	// 0 success
	// 1 empty input, not added
	// Prompt the user to enter the note text
	fmt.Println("\nEnter the note text (or leave blank to cancel):")

	// Read the user's input from the standard input (stdin)
	reader := bufio.NewReader(os.Stdin)
	note, _ := reader.ReadString('\n') // Read the input until a newline character
	note = strings.TrimSpace(note)     // Remove any leading and trailing whitespace from the input

	// Check if the input is empty
	if note == "" {
		// Print an error message in red if the input is empty
		//fmt.Println(colorRed + "Empty input. The note will not be added" + colorReset)
		return 1 // Return to the menu without adding the note
	}

	// Append a timestamp to the note
	note += getTimeStamp()

	// Add the note with the timestamp to the slice of notes
	*notes = append(*notes, note)

	// Save the updated collection of notes to the file
	SaveCollection(*notes, collectionName)
	return 0
}

// Removing notes
func removeNote(notes *[]string, collectionName string) int {
	// Return codes:
	// 0 - Success
	// 1 - No notes to remove
	// 2 - Cancellation by the user

	// Check if there are any notes to remove
	if len(*notes) == 0 {
		// Return code 1 if there are no notes
		//fmt.Println("No notes to remove.")
		return 1
	}

	var index int
	for {
		// Prompt the user to enter the note number to remove or 0 to cancel
		fmt.Print("\nEnter the note number to remove (or 0 to cancel):\n")

		// Read the user's input as an integer
		fmt.Scanln(&index)

		// If the user enters 0, cancel the removal process
		if index == 0 {
			// Return code 2 to indicate the user canceled the removal
			//fmt.Println("Note removal canceled. Returning to the menu.")
			return 2
		}

		// Validate if the entered index is within the valid range of notes
		if index < 1 || index > len(*notes) {
			// Print an error message in red if the note number is invalid
			fmt.Println(colorRed + "Invalid note number." + colorReset)
		} else {
			// Break the loop when a valid note number is entered
			break
		}
	}

	// Remove the selected note by slicing the notes slice
	*notes = append((*notes)[:index-1], (*notes)[index:]...)

	// Save the updated notes collection to the file
	SaveCollection(*notes, collectionName)

	// Return code 0 to indicate the note was successfully removed
	//fmt.Println("Note removed successfully!")
	return 0
}

// Edit
func editNote(notes *[]string, collectionName string) int {
	// Return codes:
	// 0 - success
	// 1 - no notes available to edit
	// 2 - operation canceled by the user

	// If the notes slice is empty, return 1 indicating there are no notes to edit
	if len(*notes) == 0 {
		return 1
	}

	// Prompt the user to enter the ID of the note they want to edit or 0 to cancel
	fmt.Print(colorYellow + "\nEnter the ID of the note to edit (or 0 to cancel): " + colorReset)
	var id int

	// Loop until the user provides a valid ID
	for {
		fmt.Scanln(&id)

		// Check again if notes are still available
		if len(*notes) == 0 {
			return 1
		}

		// If the user enters 0, cancel the operation and return 2
		if id == 0 {
			return 2
		}

		// If the ID is invalid (greater than the number of notes), notify the user
		if id > len(*notes) {
			fmt.Println(colorRed + "Note with the given ID not found." + colorReset)
		} else {
			// If a valid ID is provided, exit the loop
			break
		}
	}

	// Prompt the user to enter the new text for the note or leave blank to cancel
	fmt.Print(colorYellow + "Enter the new text for the note (or leave blank to cancel): \n" + colorReset)
	scanner := bufio.NewReader(os.Stdin)

	// Read the new text input by the user
	newText, _ := scanner.ReadString('\n')
	newText = strings.TrimSpace(newText) // Remove any surrounding whitespace

	// If the user leaves the input blank, cancel the operation and return 2
	if newText == "" {
		return 2
	}

	// Append a timestamp to the new note text
	newText += getTimeStamp()

	// Update the note in the list with the new text
	(*notes)[id-1] = newText

	// Save the updated collection to the file
	SaveCollection(*notes, collectionName)

	// Return 0 to indicate success
	return 0
}
