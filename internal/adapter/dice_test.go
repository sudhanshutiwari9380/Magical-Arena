package adapter_test

import (
	"magicalArena/internal/adapter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDice(t *testing.T) {
	sides := 6
	dice := adapter.NewDice(sides)

	assert.NotNil(t, dice, "NewDice should not return nil")
	assert.Equal(t, sides, dice.(*adapter.Dice).Sides, "The number of sides should be equal to the value passed to NewDice")
}

func TestRoll(t *testing.T) {
	sides := 6
	dice := adapter.NewDice(sides).(*adapter.Dice)

	for i := 0; i < 100; i++ {
		roll := dice.Roll()
		assert.GreaterOrEqual(t, roll, 1, "Roll should be greater than or equal to 1")
		assert.LessOrEqual(t, roll, sides, "Roll should be less than or equal to the number of sides")
	}
}
