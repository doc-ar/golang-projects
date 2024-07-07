package main

import "fmt"

func startAlternateScreenBuffer() {
	fmt.Print("\033[?1049h") // Enable alternate screen buffer
	fmt.Print("\033[2J")     // Clear entire screen
	fmt.Print("\033[0;0H")   // Move cursor to (0, 0)
}

func restoreOriginalScreenBuffer() {
	defer fmt.Print("\033[?1049l") // Restore normal screen buffer
}

func main() {
	var loop bool = true
	startAlternateScreenBuffer()
	fmt.Printf("Welcome to my app!")

	for loop {
		var exit string = "no"

		fmt.Printf("\nEnter yes to exit: ")
		fmt.Scanf("%s", &exit)

		if exit == "yes" {
			loop = false
		}
	}
	restoreOriginalScreenBuffer()
}
