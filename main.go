package main

import (
	"fmt"
	"math/rand"
)

func displayGrid(numbers [9]string) {
	for i := 0; i < len(numbers); i++ {
		if (i+1)%3 == 0 {
			fmt.Println(numbers[i])
		} else {
			fmt.Print(numbers[i])
			if numbers[i] != "❌" || numbers[i] != "⭕" {
				fmt.Print(" ")
			}
		}
	}
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	clearConsole()
	//give options
	fmt.Println("(1) 🔒 Mastermind")
	fmt.Println("(2) ⭕ Tic-Tac-Toe With Friend")
	fmt.Println("(3) ⛓️  Hangman")
	fmt.Println("(4) ⚔️  Attack")

	//decide...
	var mainChoice int
	fmt.Print("Choose a game to play (1-4): ")
	_, err := fmt.Scanln(&mainChoice)
	if err != nil {
		fmt.Println("Error reading input:", err)
	} else if mainChoice == 1 {

		// MASTERMIND
		colourList := [6]string{"🟠", "🟡", "🟢", "🔵", "🟣", "🟤"}

		fmt.Println("One Player")
		fmt.Println("Two Player")

		var mastermindPlayers int
		fmt.Print("Choose amount of players (1 or 2): ")
		_, err := fmt.Scanln(&mastermindPlayers)
		if err != nil {
			fmt.Println("Error reading input:", err)
		} else {
			if mastermindPlayers == 1 {
				for i := 0; i < 4; i++ {
					rand.Intn(len(colourList))

				}
			}
		}
	} else if mainChoice == 2 {
		clearConsole()
		//TIC TAC TOE
		currentPlayer := "❌"
		// Setup
		numbers := [9]string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣"}
		displayGrid(numbers)
		if currentPlayer == "❌" {
			var usrPlace int
			fmt.Print(currentPlayer, "'s turn: ")
			_, err := fmt.Scan(&usrPlace)
			if err != nil {
				fmt.Println("Error reading input:", err)
			} else {
				numbers[usrPlace - 1] = currentPlayer
				displayGrid(numbers)
			}
		}

	} else if mainChoice == 3 {
		clearConsole()
		//HANGMAN
		fmt.Println("IM NOT DONE DOING THIS")
	} else if mainChoice == 4 {
		clearConsole()
		//ATTACK
		fmt.Println("IM NOT DONE DOING THIS")
	}
}
