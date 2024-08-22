package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	// Colors
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
)

func printHelp() {
	fmt.Println("Usage: ./notestool [COLLECTION_NAME]")
	//fmt.Println("Manage your notes collection using this tool.")
	fmt.Println("The tool opens [COLLECTION_NAME] or create new one if it does not exist")
}

func clearConsole() {
	//fmt.Print("\033[H\033[2J")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func greeting() {
	fmt.Println("*** Welcome to the notes tool! ***")
}

// Get time in format [Last modified: hh:mm:ss dd.mm.yyyy]
func getTimeStamp() string {
	currentTime := time.Now()
	// formatting time and date

	result := fmt.Sprintf(" [Last modified: %v]", currentTime.Format("15:04:05 02.01.2006"))
	return result
}
