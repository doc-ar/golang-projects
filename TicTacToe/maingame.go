package main

import (
	"fmt"
)

var gridValues = [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}

func requestSelection() {
}

func render() {
	const seperator string = "- ⚬ - ⚬ -"
	for i := 0; i < 9; i = i + 3 {
		// Print row
		fmt.Printf("%s | %s | %s\n", gridValues[i], gridValues[i+1], gridValues[i+2])
		// Print seperator and exclude it in last iteration
		if i < 6 {
			fmt.Println(seperator)
		}
	}
}

func checkWin(player string) bool {
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

func main() {
	fmt.Println(checkWin("X"))
	fmt.Println(checkWin("O"))
}
