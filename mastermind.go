package main

func masternimdGame() {
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
}