package main

import (
	"./ring"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
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
	fmt.Println("Part 1 example:", playGame(9, 25), "(", 32, ")")
	fmt.Println("Part 1 demo 1:", playGame(10, 1618), "(", 8317, ")")
	fmt.Println("Part 1 demo 5:", playGame(30, 5807), "(", 37305, ")")
	fmt.Println("Part 1:", playGame(459, 71790))

	fmt.Println("Part 2:", playGame(459, 71790 * 100))

}

func playGame(playerCount int, lastMarbleValue int) int {
	score := map[int]int{}
	game := ring.New(1)
	game.Value = 0

	bar := pb.StartNew(lastMarbleValue)
	bar.SetWidth(80)

	for i := 1; i <= lastMarbleValue; i++ {
		bar.Increment()
		if i % 23 == 0 {
			game = game.Move(-8)
			score[i % playerCount] += i + game.Value.(int) // Update the score for the current player
			game = game.Pop().Move(1) // Remove the previous marble and continue with the next one.
		} else {
			marble := ring.New(1)
			marble.Value = i
			game = game.Link(marble)
		}
	}

	bar.Finish()

	return max(score)
}
