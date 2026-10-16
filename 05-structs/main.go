// Lesson 5: structs, methods, and value vs pointer receivers.
package main

import "fmt"

type Player struct {
	name   string
	money  int
	health int
}

// TakeDamage lowers health but never below zero.
// Pointer receiver: it changes the original Player, not a copy.
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

// IsAlive only reads, so a value receiver is enough.
func (p Player) IsAlive() bool {
	return p.health > 0
}

func main() {
	p := Player{
		name:   "Ilia",
		money:  100,
		health: 100,
	}

	p.TakeDamage(20)
	p.TakeDamage(30)
	p.TakeDamage(60) // more than the remaining health, so it clamps to 0

	if p.IsAlive() {
		fmt.Println("Player is alive")
	} else {
		fmt.Println("Player is dead")
	}

	if p.Pay(100) {
		fmt.Println("1 : Payment successful")
	} else {
		fmt.Println("1 : Not enough money")
	}

	if p.Pay(50) {
		fmt.Println("2 : Payment successful")
	} else {
		fmt.Println("2 : Not enough money")
	}

	if !p.Pay(-1000) {
		fmt.Println("3 : Negative payment rejected")
	}

	fmt.Printf("%+v\n", p)
}
