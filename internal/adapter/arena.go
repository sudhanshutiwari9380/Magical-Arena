// arena.go
package adapter

import (
	"errors"
	"fmt"
	"magicalArena/port"
)

type Arena struct {
	MapOfPlayers map[string]port.IPlayer
	Dice         port.IDice
}

func NewArena(dice port.IDice) port.IArena {
	return &Arena{
		MapOfPlayers: make(map[string]port.IPlayer),
		Dice:         dice,
	}
}
func (a *Arena) AddPlayerToArena(player port.IPlayer) error {

	a.MapOfPlayers[player.GetName()] = player
	return nil
}
func (a *Arena) Fight(playerA, playerB string) error {
	var p1, p2 port.IPlayer
	var ok bool
	if p1, ok = a.MapOfPlayers[playerA]; !ok {
		return errors.New("No player found with this name " + playerA)
	}
	if p2, ok = a.MapOfPlayers[playerB]; !ok {
		return errors.New("No player found with this name " + playerB)
	}
	if !(p1.GetHealth() <= p2.GetHealth()) {
		temp := p2
		p2 = p1
		p1 = temp
	}
	fmt.Printf("first Attacker is %s\n", p1.GetName())
	if !p1.IsAlive() {
		fmt.Printf("Winner is %s", p2.GetName())
		return nil
	} else if !p2.IsAlive() {
		fmt.Printf("Winner is %s", p1.GetName())
		return nil
	} else if !p1.IsAlive() && !p2.IsAlive() {
		fmt.Printf("Draw")
		return nil
	}
	for p1.IsAlive() && p2.IsAlive() {
		a.Turn(p1, p2)
		if !p2.IsAlive() {
			fmt.Printf("%s won!\n", p1.GetName())
			return nil
		}

		a.Turn(p2, p1)
		if !p1.IsAlive() {
			fmt.Printf("%s won!\n", p2.GetName())
			return nil
		}
	}
	return nil
}

func (a *Arena) Turn(attacker, defender port.IPlayer) {
	attackRoll := a.Dice.Roll()
	defendRoll := a.Dice.Roll()

	attackDamage := attacker.GetAttack() * attackRoll
	defendStrength := defender.GetStrength() * defendRoll
	damage := attackDamage - defendStrength

	if damage > 0 {
		defender.TakeDamage(damage)

	}

	fmt.Printf("%s attacks %s: attack roll %d, defend roll %d, damage %d\n",
		attacker.GetName(), defender.GetName(), attackRoll, defendRoll, damage)
	fmt.Printf("%s health: %d, %s health: %d\n",
		attacker.GetName(), attacker.GetHealth(), defender.GetName(), defender.GetHealth())
}
