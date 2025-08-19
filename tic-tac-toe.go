package main

import (
	"fmt"
)


func displayGrid(numbers [9]string) {
	for i := 0; i < len(numbers); i++ {
		if (i + 1)%3 == 0 {
			fmt.Println(numbers[i])
		} else {
			if i != 0 {
				if numbers[i - 1] == "❌" || numbers[i - 1] == "⭕" {
					fmt.Print(numbers[i])
				} else {
					fmt.Print(numbers[i])
					fmt.Print(" ")
				}
			} else {
				fmt.Print(numbers[i])
				fmt.Print(" ")
			}
		}
	}
}

func ticTacToeGame() {
	currentPlayer := "❌"
		// Setup
		numbers := [9]string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣"}
		displayGrid(numbers)
		var turnNum int
		for turnNum < 9 {
			if currentPlayer == "❌" {
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
				}
				
			}			
		}
}