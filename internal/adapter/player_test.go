package adapter_test

import (
	"magicalArena/internal/adapter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	name := "Test Player"
	health := 100
	strength := 20
	attack := 30

	player := adapter.NewPlayer(name, health, strength, attack)
	actualName := player.GetName()
	assert.Equal(t, actualName, name)
	actualHealth := player.GetHealth()
	assert.Equal(t, actualHealth, health)
	actualStrength := player.GetStrength()
	assert.Equal(t, actualStrength, strength)
	actualAttack := player.GetAttack()
	assert.Equal(t, actualAttack, attack)
}

func TestIsAlive(t *testing.T) {
	player := adapter.NewPlayer("Test Player", 100, 20, 30)

	if !player.IsAlive() {
		t.Errorf("Expected player to be alive")
	}
	actual := player.IsAlive()
	assert.Equal(t, actual, true)
	player.TakeDamage(100)
	actual = player.IsAlive()
	assert.Equal(t, actual, false)

}

func TestTakeDamage(t *testing.T) {
	player := adapter.NewPlayer("Test Player", 100, 20, 30)

	player.TakeDamage(30)
	actualHealth := player.GetHealth()
	assert.Equal(t, actualHealth, 70)

	player.TakeDamage(70)
	actualHealth = player.GetHealth()
	assert.Equal(t, actualHealth, 0)

	player.TakeDamage(10)
	actualHealth = player.GetHealth()
	assert.Equal(t, actualHealth, 0)
}

func TestGetName(t *testing.T) {
	player := adapter.NewPlayer("Test Player", 100, 20, 30)

	actualName := player.GetName()
	assert.Equal(t, actualName, "Test Player")

}

func TestGetHealth(t *testing.T) {
	player := adapter.NewPlayer("Test Player", 100, 20, 30)

	actualHealth := player.GetHealth()
	assert.Equal(t, actualHealth, 100)

}

func TestGetStrength(t *testing.T) {
	player := adapter.NewPlayer("Test Player", 100, 20, 30)

	assert.Equal(t, player.GetStrength(), 20)
}

func TestGetAttack(t *testing.T) {
	player := adapter.NewPlayer("Test Player", 100, 20, 30)

	assert.Equal(t, player.GetAttack(), 30)
}
