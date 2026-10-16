// Lesson 6: pointers — &, *, constructors and slices of pointers.
package main

import "fmt"

type Player struct {
	name   string
	money  int
	health int
}

// newPlayer is a constructor: Go has no classes, so a NewX function
// builds the value and returns a pointer to it.
func newPlayer(name string) *Player {
	return &Player{name: name, money: 500, health: 100}
}

// TakeDamage lowers health but never below zero.
func (p *Player) TakeDamage(dmg int) {
	if dmg < 0 {
		return
	}
	if dmg > p.health {
		p.health = 0
		return
	}
	p.health -= dmg
}

// Pay spends money only if the amount is valid and affordable.
func (p *Player) Pay(amount int) bool {
	if amount < 0 || p.money < amount {
		return false
	}
	p.money -= amount
	return true
}

// IsAlive only reads, but it uses a pointer receiver too so that
// every method on Player is consistent.
func (p *Player) IsAlive() bool {
	return p.health > 0
}

// damagePlayers works on []*Player, so each p points at the real player
// (with []Player, p would be a copy and nobody would take damage).
func damagePlayers(players []*Player, dmg int) {
	for _, p := range players {
		p.TakeDamage(dmg)
	}
}

// swap exchanges two ints through their addresses.
func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	// Copy vs pointer
	a := 5
	b := a  // b is a copy
	c := &a // c is a's address

	b = 10
	*c = 20

	fmt.Println(a, b, *c) // 20 10 20

	a1, b1 := 1, 2
	swap(&a1, &b1)
	fmt.Println(a1, b1) // 2 1

	players := []*Player{
		newPlayer("Ilia"),
		newPlayer("John"),
		newPlayer("Mobina"),
	}

	damagePlayers(players, 30)
	players[1].TakeDamage(80)

	for _, p := range players {
		if p.IsAlive() {
			fmt.Printf("%s is alive (%d HP)\n", p.name, p.health)
		} else {
			fmt.Printf("%s is dead\n", p.name)
		}
	}
}
