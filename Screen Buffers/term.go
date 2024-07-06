package main

import (
	"fmt"
)

func startAlternateScreenBuffer() {
	fmt.Print("\033[?1049h") // Enable alternate screen buffer
	fmt.Print("\033[2J")     // Clear entire screen
	fmt.Print("\033[0;0H")   // Move cursor to (0, 0)
}

func restoreOriginalScreenBuffer() {
	defer fmt.Print("\033[?1049l") // Restore normal screen buffer
}

func main() {
	var buf bool = true
	startAlternateScreenBuffer()

	for buf {
		var exit string = "no"

		fmt.Println("Welcome to my app!")
		fmt.Print("Enter yes to exit: ")
		fmt.Scanf("%s", &exit)

		if exit == "yes" {
			buf = false
			restoreOriginalScreenBuffer()
		}
	}
}
