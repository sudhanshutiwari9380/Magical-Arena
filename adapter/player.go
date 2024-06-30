package adapter

import "magicalArena/port"

type player struct {
	Name     string
	Health   int
	Strength int
	Attack   int
}

func NewPlayer(name string, health, strength, attack int) port.IPlayer {

	return &player{
		Name:     name,
		Health:   health,
		Strength: strength,
		Attack:   attack,
	}
}
func (p *player) IsAlive() bool {
	return p.Health > 0
}

func (p *player) TakeDamage(damage int) {
	p.Health -= damage
	if p.Health < 0 {
		p.Health = 0
	}
}
func (p *player) GetName() string {
	return p.Name
}
func (p *player) GetHealth() int {
	return p.Health
}
func (p *player) GetStrength() int {
	return p.Strength
}
func (p *player) GetAttack() int {
	return p.Attack
}
