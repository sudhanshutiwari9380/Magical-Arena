package adapter

import (
	"magicalArena/port"
	"math/rand"
)

type Dice struct {
	Sides int
}

func NewDice(sidesofDice int) port.IDice {
	return &Dice{Sides: sidesofDice}
}

func (d *Dice) Roll() int {
	return rand.Intn(d.Sides) + 1
}
