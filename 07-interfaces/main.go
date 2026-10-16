// Lesson 7: interfaces — implicit implementation, polymorphism,
// fmt.Stringer and type switches.
package main

import "fmt"

// Damageable is anything that can take damage and be alive or not.
// No "implements" keyword: any type with these methods satisfies it.
type Damageable interface {
	TakeDamage(dmg int)
	IsAlive() bool
}

// ---------- Player ----------

type Player struct {
	name   string
	money  int
	health int
}

func newPlayer(name string) *Player {
	return &Player{name: name, money: 500, health: 100}
}

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

func (p *Player) Pay(amount int) bool {
	if amount < 0 || p.money < amount {
		return false
	}
	p.money -= amount
	return true
}

func (p *Player) IsAlive() bool {
	return p.health > 0
}

// String makes *Player a fmt.Stringer, so fmt.Println prints it nicely.
func (p *Player) String() string {
	return fmt.Sprintf("%s [HP:%d $%d]", p.name, p.health, p.money)
}

// ---------- Vehicle ----------

type Vehicle struct {
	model  string
	engine int
}

func (v *Vehicle) TakeDamage(dmg int) {
	if dmg < 0 {
		return
	}
	v.engine -= dmg
	if v.engine < 0 {
		v.engine = 0
	}
}

func (v *Vehicle) IsAlive() bool {
	return v.engine > 0
}

func (v *Vehicle) String() string {
	return fmt.Sprintf("%s [Engine:%d]", v.model, v.engine)
}

// ---------- Working with any Damageable ----------

// explode doesn't know or care whether a target is a Player or a Vehicle.
func explode(targets []Damageable, dmg int) {
	for _, t := range targets {
		t.TakeDamage(dmg)
	}
}

// report uses the interface for IsAlive, and a type switch only when it
// needs fields that belong to one concrete type (name / model).
func report(targets []Damageable) {
	for _, t := range targets {
		if t.IsAlive() {
			fmt.Println("✅", t)
			continue
		}

		switch v := t.(type) {
		case *Player:
			fmt.Printf("💀 %s is dead\n", v.name)
		case *Vehicle:
			fmt.Printf("🔥 %s is destroyed\n", v.model)
		}
	}
}

func main() {
	targets := []Damageable{
		newPlayer("Ilia"),
		&Vehicle{model: "Sultan", engine: 1000},
		newPlayer("Mobina"),
		&Vehicle{model: "Banshee", engine: 500},
		&Vehicle{model: "Faggio", engine: 100},
	}

	fmt.Println("--- Explosion 1 (60 dmg) ---")
	explode(targets, 60)
	report(targets)

	fmt.Println("--- Explosion 2 (50 dmg) ---")
	explode(targets, 50)
	report(targets)
}
