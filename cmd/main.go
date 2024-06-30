// main.go
package main

import (
	"fmt"
	"magicalArena/internal/adapter"
)

func main() {
	// Name: "Player A", Health: 50, Strength: 5, Attack: 10
	var player1Name, player2Name string
	var player1Health, player1Strength, player1Attack, player2Health, player2Strength, player2Attack int
	fmt.Println("Enter the Player 1 details")
	fmt.Println("Enter The name")
	fmt.Scan(&player1Name)
	fmt.Println("Enter The Health")
	fmt.Scan(&player1Health)
	if player1Health < 0 {
		fmt.Printf("health of a player can't be less than 0 %d ", player1Health)
		return
	}
	fmt.Println("Enter The Strength")

	fmt.Scan(&player1Strength)
	if player1Strength < 0 {
		fmt.Printf("Strength of a player can't be less than 0 %d ", player1Strength)
		return
	}
	fmt.Println("Enter The Attack")
	fmt.Scan(&player1Attack)
	if player1Attack < 0 {
		fmt.Printf("Attack of a player can't be less than 0 %d ", player1Attack)
		return
	}
	fmt.Println("Enter The Player 2 details")
	fmt.Println("Enter The name")
	fmt.Scan(&player2Name)
	fmt.Println("Enter The Health")
	fmt.Scan(&player2Health)
	if player2Health < 0 {
		fmt.Printf("health of a player can't be less than 0 %d ", player2Health)
		return
	}
	fmt.Println("Enter The Strength")
	fmt.Scan(&player2Strength)
	if player2Strength < 0 {
		fmt.Printf("Strength of a player can't be less than 0 %d ", player2Strength)
		return
	}
	fmt.Println("Enter The Attack")
	fmt.Scan(&player2Attack)
	if player2Attack < 0 {
		fmt.Printf("Attack of a player can't be less than 0 %d ", player2Attack)
		return
	}
	fmt.Println("Lets Start the Game")

	playerA := adapter.NewPlayer(player1Name, player1Health, player1Strength, player1Attack)
	playerB := adapter.NewPlayer(player2Name, player2Health, player2Strength, player2Attack)
	dice := adapter.NewDice(6)
	arena := adapter.NewArena(dice)
	arena.AddPlayerToArena(playerA)
	arena.AddPlayerToArena(playerB)

	arena.Fight(player1Name, player2Name)
}
