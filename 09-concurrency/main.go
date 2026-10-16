// Lesson 9: concurrency — goroutines and channels.
// Loads 10 players from a fake database, first one by one, then all at once.
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Player struct {
	id    int
	name  string
	money int
}

// loadPlayer simulates a database query that takes 0.5–1s.
func loadPlayer(id int) Player {
	time.Sleep(time.Duration(500+rand.IntN(500)) * time.Millisecond)
	return Player{id: id, name: fmt.Sprintf("player_%d", id), money: rand.IntN(1000)}
}

// loadSequential loads players one after another (~7.5s for 10).
func loadSequential(count int) []Player {
	players := make([]Player, 0, count)
	for id := 1; id <= count; id++ {
		players = append(players, loadPlayer(id))
	}
	return players
}

// loadConcurrent starts one goroutine per player and collects the results
// from a channel (~1s for 10).
//
// No mutex is needed: the goroutines never touch shared variables, they only
// send their result. Only main reads the channel and writes to the slice.
func loadConcurrent(count int) []Player {
	ch := make(chan Player, count) // buffered: senders never block

	for id := 1; id <= count; id++ {
		go func() {
			ch <- loadPlayer(id) // Go 1.22+: each iteration has its own id
		}()
	}

	players := make([]Player, 0, count)
	for range count {
		players = append(players, <-ch)
	}
	return players
}

func totalMoney(players []Player) int {
	total := 0
	for _, p := range players {
		total += p.money
	}
	return total
}

func main() {
	const count = 10

	start := time.Now()
	seq := loadSequential(count)
	fmt.Printf("sequential: %d players in %v (total $%d)\n",
		len(seq), time.Since(start).Round(time.Millisecond), totalMoney(seq))

	start = time.Now()
	con := loadConcurrent(count)
	fmt.Printf("concurrent: %d players in %v (total $%d)\n",
		len(con), time.Since(start).Round(time.Millisecond), totalMoney(con))
}
