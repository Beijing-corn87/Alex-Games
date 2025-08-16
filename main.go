package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//give options
	fmt.Println("(1) 🔒 Mastermind")
	fmt.Println("(2) ⭕ Tic-Tac-Toe With Friend")
	fmt.Println("(3) ⛓️ Hangman")
	fmt.Println("(4) ⚔️ Attack")

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

		//TIC TAC TOE
	} else if mainChoice == 3 {

		//HANGMAN
	} else if mainChoice == 4 {

		//ATTACK
	}
}
