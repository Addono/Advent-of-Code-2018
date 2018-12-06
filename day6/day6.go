package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func min(i1 int, i2 int) int {
	if i1 < i2 {
		return i1
	} else {
		return i2
	}
}

func max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

type coordinate struct {
	x int
	y int
}

func distance(c1 coordinate, c2 coordinate) int {
	return abs(c1.x - c2.x) + abs(c1.y - c2.y)
}

func closest(data []coordinate, point coordinate) (int, int) {
	minDistance := math.MaxInt16
	minIndex := -1
	for i, c := range data {
		distance := distance(c, point)
		if distance < minDistance{
			minDistance = distance
			minIndex = i
		} else if distance == minDistance {
			minIndex = -1
		}
	}

	return minIndex, minDistance
}

func summedDistance(data []coordinate, point coordinate) int {
	total := 0
	for _, c := range data {
		total += distance(c, point)
	}
	return total
}

func main() {
	file, err := os.Open("./day6/input.txt")
	check(err)
	defer file.Close()

	var data []coordinate
	minC := coordinate{math.MaxInt16, math.MaxInt16} // The coordinate containing the smallest value for both x and y
	maxC := coordinate{math.MinInt16, math.MinInt16} // The coordinate containing the largest value for both x and y

	// Ingest the input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, ", ")
		x, _ := strconv.Atoi(splitted[0])
		y, _ := strconv.Atoi(splitted[1])

		minC = coordinate{min(minC.x, x), min(minC.y, y)}
		maxC = coordinate{max(maxC.x, x), max(maxC.y, y)}

		data = append(data, coordinate{x, y})
	}

	count := map[int]int{}
	notFinite := map[int]bool{}
	for x := minC.x; x <= maxC.x; x++ {
		for y := minC.y; y <= maxC.y; y++ {
			closestIndex, distance := closest(data, coordinate{x, y})

			if closestIndex == -1 {
				fmt.Print(".") // Print a full stop in case the distance is equal
			} else {
				// Print the current coordinate
				if distance == 0 {
					fmt.Print(" ") // Leave the current coordinate empty if this is the location of the coordinate
				} else {
					fmt.Print(string(closestIndex + 32)) // Otherwise print the area character
				}

				count[closestIndex] += 1

				// An area is not finite if it touches the outer border
				if x == minC.x || x == maxC.x || y == minC.y || y == maxC.y {
					notFinite[closestIndex] = true
				}
			}
		}

		fmt.Println()
	}

	// Find the largest finite area
	maxSize := -1
	maxIndex := -1
	for i, c := range count {
		if !notFinite[i] && c > maxSize {
			maxIndex = i
			maxSize = c
		}
	}

	fmt.Println("The largest finite area has size: ", maxSize, " (Day 1)")
	fmt.Println("The character to visualize the largest area is: " + string(maxIndex + 32))

	// Part 2

	// Brute force check for every coordinate within either vertical or horizontal 10000 units if the summed distance
	// is less than 10000.
	amount := 0
	for x := minC.x - 10000; x <= maxC.x + 10000; x++ {
		for y := minC.y - 10000; y <= maxC.y + 10000; y++ {
			if summedDistance(data, coordinate{x, y}) < 10000 {
				amount += 1
			}
		}
	}

	fmt.Println("Amount of positions within a summed distance of 10000: ", amount, " (Day 2)")
}