package adapter_test

import (
	"magicalArena/internal/adapter"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	name     string
	testCase func(*testing.T)
}

func Test_AddPlayerToArena(t *testing.T) {
	tests := []test{{
		name: "happy path",
		testCase: func(*testing.T) {
			dice := adapter.NewDice(6)
			player := adapter.NewPlayer("sudh", 100, 20, 30)
			arena := adapter.NewArena(dice)
			err := arena.AddPlayerToArena(player)

			assert.NoError(t, err)
		},
	}}

	for _, test := range tests {
		t.Run(test.name, test.testCase)
	}
}

func Test_Fight(t *testing.T) {
	tests := []test{{
		name: "happy path",
		testCase: func(*testing.T) {
			dice := adapter.NewDice(6)
			arena := adapter.NewArena(dice)

			player1 := adapter.NewPlayer("1", 100, 20, 30)
			arena.AddPlayerToArena(player1)

			player2 := adapter.NewPlayer("2", 100, 20, 30)
			arena.AddPlayerToArena(player2)

			err := arena.Fight(player1.GetName(), player2.GetName())
			assert.NoError(t, err)
		},
	},
		{name: "playerNotfound",
			testCase: func(*testing.T) {
				dice := adapter.NewDice(6)
				arena := adapter.NewArena(dice)

				player1 := adapter.NewPlayer("1", 1, 20, 30)
				arena.AddPlayerToArena(player1)

				err := arena.Fight(player1.GetName(), "NotFound")
				assert.Error(t, err)
				assert.ErrorContains(t, err, "No player found with this name ")
				err = arena.Fight("NotFound", player1.GetName())
				assert.Error(t, err)
				assert.ErrorContains(t, err, "No player found with this name ")
			},
		},
		{name: "make the player attacker who has less health",
			testCase: func(*testing.T) {
				dice := adapter.NewDice(6)
				arena := adapter.NewArena(dice)

				player1 := adapter.NewPlayer("1", 100, 20, 30)
				arena.AddPlayerToArena(player1)
				player2 := adapter.NewPlayer("2", 10, 20, 30)
				arena.AddPlayerToArena(player2)

				err := arena.Fight(player1.GetName(), player2.GetName())
				assert.NoError(t, err)

			},
		},
		{name: "if either of the player health is zero other is automatically winner",
			testCase: func(*testing.T) {
				dice := adapter.NewDice(6)
				arena := adapter.NewArena(dice)

				player1 := adapter.NewPlayer("1", 0, 20, 30)
				arena.AddPlayerToArena(player1)
				player2 := adapter.NewPlayer("2", 10, 20, 30)
				arena.AddPlayerToArena(player2)

				err := arena.Fight(player1.GetName(), player2.GetName())
				assert.NoError(t, err)
				err = arena.Fight(player2.GetName(), player1.GetName())
				assert.NoError(t, err)

			},
		},
		{name: "draw",
			testCase: func(*testing.T) {
				dice := adapter.NewDice(6)
				arena := adapter.NewArena(dice)

				player1 := adapter.NewPlayer("1", 0, 20, 30)
				arena.AddPlayerToArena(player1)
				player2 := adapter.NewPlayer("2", 0, 20, 30)
				arena.AddPlayerToArena(player2)

				err := arena.Fight(player1.GetName(), player2.GetName())
				assert.NoError(t, err)
				err = arena.Fight(player2.GetName(), player1.GetName())
				assert.NoError(t, err)

			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.testCase)
	}
}

func Test_Turn(t *testing.T) {
	tests := []test{}

	for _, test := range tests {
		t.Run(test.name, test.testCase)
	}
}
