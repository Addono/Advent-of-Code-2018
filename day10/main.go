package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)


type point struct {
	x int
	y int
	vx int
	vy int
}

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


func deltaPoints(points []point) (int, int) {
	minX, minY, maxX, maxY := minMaxPoints(points)

	dx := maxX - minX
	dy := maxY - minY
	return dx, dy
}

func minMaxPoints(points []point) (int, int, int, int) {
	minX := math.MaxInt32
	minY := math.MaxInt32
	maxX := math.MinInt32
	maxY := math.MinInt32
	for _, p := range points {
		minX = min(minX, p.x)
		minY = min(minY, p.y)
		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)
	}
	return minX, minY, maxX, maxY
}

func main() {
	fmt.Println("Example input:")
	solve("./day10/input_example_part_1.txt")

	fmt.Println("Full input:")
	solve("./day10/input.txt")
}

func solve(filename string) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	re := regexp.MustCompile(`position=<\s*(-?\d*?),\s*(-?\d*?)> velocity=<\s*(-?\d*?),\s*(-?\d*?)>`)

	var points []point

	// Ingest the input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		matchAll := re.FindStringSubmatch(text)

		if len(matchAll) != 5 {
			panic("Not enough elements matched")
		}

		x, _ := strconv.Atoi(matchAll[1])
		y, _ := strconv.Atoi(matchAll[2])
		vx, _ := strconv.Atoi(matchAll[3])
		vy, _ := strconv.Atoi(matchAll[4])

		points = append(points, point{x, y, vx, vy})
	}

	dx, dy := math.MaxInt32, math.MaxInt32
	var oldPoints []point
	for i := 0; ; i++ {
		points, oldPoints = progressPoints(points), points
		newdx, newdy := deltaPoints(points)
		dx, dy = min(dx, newdx), min(dy, newdy)
		if newdx > dx && newdy > dy { // Check if we stopped making progess, if so then the previous second displayed the message
			fmt.Println("Solution part 1:")
			printPoints(oldPoints) // Print the layout of the previous selection
			fmt.Println("Solution part 2:", i)
			break
		}
	}
}

func printPoints(points []point) {
	// Map all points onto a 2 dimensional map
	plane := map[int]map[int]bool{}
	for _, p := range points {
		if plane[p.y] == nil {
			plane[p.y] = map[int]bool{}
		}
		plane[p.y][p.x] = true
	}

	// Print each tile in the map
	minX, minY, maxX, maxY := minMaxPoints(points)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if plane[y][x] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func progressPoints(points []point) []point {
	var newPoints []point
	for _, p := range points {
		newPoints = append(newPoints, point{
			p.x + p.vx,
			p.y + p.vy,
			p.vx,
			p.vy,
		})
	}
	return newPoints
}
