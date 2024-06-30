# Magical-Arena
# Magical Arena

## Overview

The Magical Arena game simulates a battle between two players, each defined by health, strength, and attack attributes. Players attack each other in turns, and the game continues until one player's health reaches 0.

## Rules of the Game

- Each player has a "health", "strength", and "attack" attribute.
- Players attack in turns. The player with lower health attacks first.
- Attacking player rolls the attacking dice and the defending player rolls the defending dice.
- The attack value multiplied by the outcome of the attacking dice roll determines the damage created by the attacker.
- The defender’s strength value multiplied by the outcome of the defending dice roll determines the damage defended by the defender.
- Any damage created by the attacker that exceeds the defended damage will reduce the defender’s health.
- The game ends when any player's health reaches 0.

## Example

Given:
- Player A: Health = 50, Strength = 5, Attack = 10
- Player B: Health = 100, Strength = 10, Attack = 5
- Both dice are 6-sided (values range from 1 to 6)

Game flow:
1. Player A attacks and rolls a die (roll: 5). Player B defends and rolls a die (roll: 2).
    - Attack damage: 5 * 10 = 50
    - Defend strength: 10 * 2 = 20
    - Player B's health reduced by 30 (to 70).
2. Player B attacks and rolls a die (roll: 4). Player A defends and rolls a die (roll: 3).
    - Attack damage: 4 * 5 = 20
    - Defend strength: 5 * 3 = 15
    - Player A's health reduced by 5 (to 45).

And so on.

## Installation

1. Ensure you have Go installed on your system.
2. Clone the repository or download the ZIP file and extract it.

## Running the Game

Navigate to the project directory and run:

```sh
go run cmd/main.go

