// main.go
package main

import (
	"fmt"
	"magicalArena/adapter"
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
	fmt.Println("Enter The Strength")
	fmt.Scan(&player1Strength)
	fmt.Println("Enter The Attack")
	fmt.Scan(&player1Attack)
	fmt.Println("Enter The Player 2 details")
	fmt.Println("Enter The name")
	fmt.Scan(&player2Name)
	fmt.Println("Enter The Health")
	fmt.Scan(&player2Health)
	fmt.Println("Enter The Strength")
	fmt.Scan(&player2Strength)
	fmt.Println("Enter The Attack")
	fmt.Scan(&player2Attack)
	fmt.Println("Lets Start the Game")

	playerA := adapter.NewPlayer(player1Name, player1Health, player1Strength, player1Attack)
	playerB := adapter.NewPlayer(player2Name, player2Health, player2Strength, player2Attack)
	dice := adapter.NewDice(6)
	arena := adapter.NewArena(dice)
	arena.AddPlayerToArena(playerA)
	arena.AddPlayerToArena(playerB)

	arena.Fight(player1Name, player2Name)
}
