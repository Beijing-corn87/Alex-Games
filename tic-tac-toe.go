package main

import (
	"fmt"
	"time"
)

func checkIfWon(numbers [9]string) bool {
	var winningCombinations [8][3]int = [8][3]int{
		{1, 2, 3}, // Top row
		{4, 5, 6}, // Middle row
		{7, 8, 9}, // Bottom row

		{1, 4, 7}, // Left column
		{2, 5, 8}, // Middle column
		{3, 6, 9}, // Right column

		{1, 5, 9}, // Diagonal from top-left to bottom-right
		{3, 5, 7}, // Diagonal from top-right to bottom-left
	}
	for combo := 0; combo < len(winningCombinations); combo++ {
		matches := [3]bool{false, false, false}
		for combix := 0; combix < 3; combix++ {
			if numbers[winningCombinations[combo][combix]] == "❌" {
				matches[combix] = true
			}
		}
		if matches[0] && matches[1] && matches[2] {
			return true
		} else {
			matches[0], matches[1], matches[2] = false, false, false
		}
	}
	return false
}

func displayGrid(numbers [9]string) {
	for i := 0; i < len(numbers); i++ {
		if (i+1)%3 == 0 {
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
		} else if numbers[usrPlace-1] == "❌" || numbers[usrPlace-1] == "⭕" {
			clearConsole()
			displayGrid(numbers)
			fmt.Println("Can't place there! Try again: ")
			time.Sleep(1000)
		} else {
			numbers[usrPlace-1] = currentPlayer
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
	checkIfWon(numbers)
	fmt.Println(hasWon, winner)
}
