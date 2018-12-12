package main

import (
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
)

type coordinate struct {
	x int
	y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input := 8444

	fixedC := maxFuelFixedSize(input, 3)
	fmt.Println("Part 1:", fixedC.x, ",", fixedC.y)

	dynamicC, size := maxFuel(input)
	fmt.Println("Part 2:", dynamicC.x, ",", dynamicC.y, ",", size)
}

func maxFuel(serial int) (coordinate, int) {
	fuelCell := [300][300]int{}
	for x, row := range fuelCell {
		for y := range row {
			fuelCell[x][y] = computePowerLevel(coordinate{x, y}, serial)
		}
	}

	best := -10
	var bestSize int
	var c coordinate

	b := pb.StartNew(300)
	for size := 1; size < 300; size++ {
		b.Increment()
		for x := 0; x < 300 - size; x++ {
			for y := 0; y < 300 - size; y++ {
				fuel := sumFuel(fuelCell, coordinate{x, y}, size)
				if fuel > best {
					best = fuel
					bestSize = size
					c = coordinate{x, y}
				}
			}
		}
	}
	b.Finish()
	return c, bestSize
}

func maxFuelFixedSize(serial int, size int) coordinate {
	fuelCell := [300][300]int{}
	for x, row := range fuelCell {
		for y := range row {
			fuelCell[x][y] = computePowerLevel(coordinate{x, y}, serial)
		}
	}

	best := -10
	var c coordinate
	for x := 0; x < 300 - size; x++ {
		for y := 0; y < 300 - size; y++ {
			fuel := sumFuel(fuelCell, coordinate{x, y}, size)
			if fuel > best {
				best = fuel
				c = coordinate{x, y}
			}
		}
	}
	return c
}

func sumFuel(fuelCell [300][300]int, c coordinate, size int) int {
	sum := 0
	for dx := 0; dx < size; dx++ {
		for dy := 0; dy < size; dy++ {
			sum += fuelCell[c.x+dx][c.y+dy]
		}
	}
	return sum
}

func computePowerLevel(c coordinate, serial int) int {
	rackId := c.x + 10
	return ((rackId * c.y + serial) * rackId / 100 % 10) - 5
}
