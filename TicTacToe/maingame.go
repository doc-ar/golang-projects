package main

import "fmt"

const SEPERATOR string = "- ⚬ - ⚬ -"

var (
	isWinOrDraw bool = false
	gridValues       = [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
)

// Function to use alternate screen buffer
func startAlternateScreenBuffer() {
	fmt.Print("\033[?1049h")
}

// Function to clear terminal screen
func clearScreen() {
	fmt.Print("\033[2J")   // Clear entire screen
	fmt.Print("\033[0;0H") // Move cursor to (0, 0)
}

// Function to restore original screen buffer
func restoreOriginalScreenBuffer() {
	defer fmt.Print("\033[?1049l")
}

func resetGrid() {
	for i := 0; i < 9; i++ {
		gridValues[i] = " "
	}
}

func render() {
	for i := 0; i < 9; i = i + 3 {
		// Print row
		fmt.Printf("%s | %s | %s\n", gridValues[i], gridValues[i+1], gridValues[i+2])
		// Print seperator and exclude it in last iteration
		if i < 6 {
			fmt.Println(SEPERATOR)
		}
	}
}

func checkDraw() bool {
	var draw bool = true
	for i := 0; i < 9; i++ {
		if gridValues[i] == " " {
			draw = false
			break
		}
	}
	return draw
}

func checkWin(player string) bool {
	// The win conditions are indexes corresponding to the grid values
	winConditions := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	// Check win conditions against grid values
	for i := 0; i < 8; i++ {
		if gridValues[winConditions[i][0]] == player &&
			gridValues[winConditions[i][1]] == player &&
			gridValues[winConditions[i][2]] == player {
			return true
		}
	}
	return false
}

func gameIntro() {
	var continueGame string = "n"

	fmt.Printf("Welcome to Tic Tac Toe\n\n")
	fmt.Printf("The game will be played in this order\n\n")
	fmt.Printf("1 | 2 | 3\n%s\n4 | 5 | 6\n%s\n7 | 8 | 9\n\n", SEPERATOR, SEPERATOR)

	for continueGame != "y" {
		fmt.Printf("Press y to continue: ")
		fmt.Scanf("%s", &continueGame)
	}
}

func playGame() {
	clearScreen()
	for !isWinOrDraw {
		if !isWinOrDraw {
			playTurn("X")
		}
		if !isWinOrDraw {
			playTurn("O")
		}
	}
}

func playTurn(player string) {
	var loop bool = true

	clearScreen()
	render()
	fmt.Printf("\nPlayer %s's Turn", player)

	for loop {
		var position int
		fmt.Print("\nEnter position: ")
		fmt.Scanf("%d", &position)

		if position < 1 || position > 9 {
			fmt.Println("\nInvalid position. Try again!!")
		} else if gridValues[position-1] != " " {
			fmt.Println("\nInvalid position. Try again!!")
		} else {
			gridValues[position-1] = player
			loop = false
			if checkDraw() {
				isWinOrDraw = true
				drawDialogue()
			} else if checkWin(player) {
				isWinOrDraw = true
				winDialogue(player)
			}
		}
	}
}

func winDialogue(player string) {
	clearScreen()
	fmt.Printf("\033[1m========================\033[0m\n")
	fmt.Printf("\033[1m==== Player %s Wins =====\033[0m\n", player)
	fmt.Printf("\033[1m========================\033[0m\n")
}

func drawDialogue() {
	clearScreen()
	fmt.Printf("\033[1m==============================\033[0m\n")
	fmt.Printf("\033[1m==== The game was a draw =====\033[0m\n")
	fmt.Printf("\033[1m==============================\033[0m\n")
}

func main() {
	startAlternateScreenBuffer()

	var reset bool = true
	gameIntro()

	for reset {
		var play string = "y"
		if play == "y" {
			playGame()
		}
		fmt.Printf("\nPress \"y\" to play again. Press anything else to exit.")
		fmt.Printf("\n\n=====> ")
		fmt.Scanf("%s", &play)

		if play == "y" {
			isWinOrDraw = false
			resetGrid()
		} else {
			reset = false
		}
	}

	restoreOriginalScreenBuffer()
}
