// Lesson 8: errors — error values, sentinel errors, wrapping with %w,
// and checking them with errors.Is.
package main

import (
	"errors"
	"fmt"
)

// Sentinel errors live at package level so they are created once and
// callers can compare against them with errors.Is.
var (
	ErrNotEnoughMoney = errors.New("not enough money")
	ErrInvalidAmount  = errors.New("invalid amount")
	ErrItemNotFound   = errors.New("item not found")
)

type Player struct {
	name   string
	money  int
	health int
}

func newPlayer(name string, money int) *Player {
	return &Player{name: name, money: money, health: 100}
}

// Pay returns an error instead of a bool, so the caller knows *why* it failed.
func (p *Player) Pay(amount int) error {
	if amount < 0 {
		return ErrInvalidAmount
	}
	if p.money < amount {
		return ErrNotEnoughMoney
	}
	p.money -= amount
	return nil
}

// Buy wraps errors with %w: the message gains context ("buy pistol: ...")
// while errors.Is can still find the original sentinel inside.
func Buy(p *Player, shop map[string]int, item string) error {
	price, ok := shop[item]
	if !ok {
		return fmt.Errorf("buy %s: %w", item, ErrItemNotFound)
	}
	if err := p.Pay(price); err != nil {
		return fmt.Errorf("buy %s: %w", item, err)
	}
	return nil
}

func main() {
	player := newPlayer("Ilia", 500)
	shop := map[string]int{
		"phone":  200,
		"bread":  10,
		"pistol": 1500,
	}

	for _, item := range []string{"phone", "pistol", "car"} {
		err := Buy(player, shop, item)
		switch {
		case err == nil:
			fmt.Printf("✅ bought %s ($%d left)\n", item, player.money)
		case errors.Is(err, ErrNotEnoughMoney):
			fmt.Println("💸", err)
		case errors.Is(err, ErrItemNotFound):
			fmt.Println("❓", err)
		default:
			fmt.Println("⚠️ unexpected error:", err)
		}
	}
}
