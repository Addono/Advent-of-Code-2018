package main

import (
	"container/ring"
	"fmt"
	"math"
)

func max(data map[int]int) int {
	max := math.MinInt8
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

func printRing(r *ring.Ring) {
	r.Do(func(i interface{}) {fmt.Print(i, " ")})
	fmt.Println()
}

func main() {
	fmt.Println("Part 1 example:", part1(9, 25), "(", 32, ")")
	fmt.Println("Part 1 demo 1:", part1(10, 1618), "(", 8317, ")")
	fmt.Println("Part 1 demo 5:", part1(30, 5807), "(", 37305, ")")
	fmt.Println("Part 1:", part1(459, 71790))
}

func part1(playerCount int, lastMarbleValue int) int {
	score := map[int]int{}
	game := ring.New(1)
	game.Value = 0

	for i := 1; i <= lastMarbleValue; i++ {
		if i % 23 == 0 {
			game = game.Move(-8)
			score[i % playerCount] += i + game.Value.(int) // Update the score for the current player
			game = game.Unlink(game.Len()-1).Move(1) // Remove all but the previous marble and continue with the next one.
		} else {
			marble := ring.New(1)
			marble.Value = i
			game = game.Link(marble)
		}
	}

	return max(score)
}