package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id string
	x int
	y int
	w int
	h int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitLink(s string, sep string) (string, string) {
	x := strings.Split(s, sep)
	return x[0], x[1]
}

func intSplitLink(s string, sep string) (int, int) {
	a1, a2 := splitLink(s, sep)
	i1, _ := strconv.Atoi(a1)
	i2, _ := strconv.Atoi(a2)
	return i1, i2
}

func checkCollision(data claim, cloth *[1001][1001]int) bool {
	for i := data.x; i < data.x+data.w; i++ {
		for j := data.y; j < data.y+data.h; j++ {
			if cloth[i][j] > 1 { // Check if there is more than one other claim at this position, since this square is also included.
				return true
			}
		}
	}
	return false
}

func main() {
	file, err := os.Open("./day3/input.txt")
	check(err)
	defer file.Close()

	var input []claim

	// Ingest all input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// Extract the text
		id, remainder := splitLink(text, " @ ")
		coordinates, size := splitLink(remainder, ": ")
		x, y := intSplitLink(coordinates, ",")
		w, h := intSplitLink(size, "x")

		input = append(input, claim{id, x, y, w, h})
	}

	// Define the variables which tracks the progress
	cloth := [1001][1001]int{}
	count := 0

	for _, data := range input {
		for i := data.x; i < data.x+data.w; i++ {
			for j := data.y; j < data.y+data.h; j++ {
				cloth[i][j] += 1 // Track that we found another occurrence of this coordinate.
				if cloth[i][j] == 2 { // If this was the second one, then it is used double for the first time
					count += 1 // Count the amount of coordinates which is used twice or more.
				}
			}
		}
	}

	fmt.Println("Task 1:\t" + strconv.Itoa(count))

	for _, data := range input {
		if !checkCollision(data, &cloth) {
			fmt.Println("Task 2:\t" + data.id)
			break
		}
	}
}
