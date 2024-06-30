package main

import (
	"fmt"
	"magicalArena/internal/adapter"
)

func main() {
	var player1Name, player2Name string
	var player1Health, player1Strength, player1Attack, player2Health, player2Strength, player2Attack int

	// Function to get player details with validation
	getPlayerDetails := func(playerNum int) (string, int, int, int) {
		var name string
		var health, strength, attack int

		fmt.Printf("Enter the Player %d details\n", playerNum)
		fmt.Print("Enter The name: ")
		fmt.Scan(&name)

		for {
			fmt.Print("Enter The Health: ")
			fmt.Scan(&health)
			if health >= 0 {
				break
			}
			fmt.Println("Health of a player can't be less than 0. Please enter a valid value.")
		}

		for {
			fmt.Print("Enter The Strength: ")
			fmt.Scan(&strength)
			if strength >= 0 {
				break
			}
			fmt.Println("Strength of a player can't be less than 0. Please enter a valid value.")
		}

		for {
			fmt.Print("Enter The Attack: ")
			fmt.Scan(&attack)
			if attack >= 0 {
				break
			}
			fmt.Println("Attack of a player can't be less than 0. Please enter a valid value.")
		}

		return name, health, strength, attack
	}

	// Get details for player 1
	player1Name, player1Health, player1Strength, player1Attack = getPlayerDetails(1)

	// Get details for player 2
	player2Name, player2Health, player2Strength, player2Attack = getPlayerDetails(2)

	// Create players, dice, and arena
	playerA := adapter.NewPlayer(player1Name, player1Health, player1Strength, player1Attack)
	playerB := adapter.NewPlayer(player2Name, player2Health, player2Strength, player2Attack)
	dice := adapter.NewDice(6)
	arena := adapter.NewArena(dice)
	arena.AddPlayerToArena(playerA)
	arena.AddPlayerToArena(playerB)

	fmt.Println("Let's Start the Game")
	arena.Fight(player1Name, player2Name)
}
