package port

type IArena interface {
	AddPlayerToArena(IPlayer) error
	Fight(string, string) error
	Turn(attacker, defender IPlayer)
}
type IPlayer interface {
	IsAlive() bool
	TakeDamage(int)
	GetName() string
	GetHealth() int
	GetStrength() int
	GetAttack() int
}
type IDice interface {
	Roll() int
}
