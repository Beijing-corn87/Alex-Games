package main

import (
	"fmt"
	"math/rand"
)

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
		clearConsole()
		// MASTERMIND

		masternimdGame()
	} else if mainChoice == 2 {
		
		//TIC TAC TOE
		ticTacToeGame()
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
