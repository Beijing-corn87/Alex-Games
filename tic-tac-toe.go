package main

import (
	"fmt"
)


func displayGrid(numbers [9]string) {
	for i := 0; i < len(numbers); i++ {
		if (i + 1)%3 == 0 {
			fmt.Println(numbers[i])
		} else {
			fmt.Print(numbers[i])
		}
	}
}

// ENTRYPOINT
func ticTacToeGame() {
	// Setup
	currentPlayer := "❌"
	hasWon := false
	var winner string
	
	//acuals game
	numbers := [9]string{"1 ", " 2 ", " 3 ", "4 ", " 5 ", " 6 ", "7 ", " 8 ", " 9 "}
	displayGrid(numbers)
	var turnNum int
	for turnNum < 9 {
		var usrPlace int
		fmt.Print(currentPlayer, "'s turn: ")
		_, err := fmt.Scan(&usrPlace)
		if err != nil {
			fmt.Println("Error reading input:", err)
		} else {
			numbers[usrPlace - 1] = currentPlayer
			clearConsole()
			displayGrid(numbers)
			turnNum++
			if currentPlayer == "❌" {
				currentPlayer = "⭕"
			} else {
				currentPlayer = "❌"
			}
		}			
	}
}
