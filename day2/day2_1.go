package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./day2/input.txt")
	check(err)
	defer file.Close()

	double_count := 0
	triple_count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		amount := map[int32]int{}

		for _, r := range text {
			amount[r] += 1
		}

		double := false
		triple := false
		for _, value := range amount {
			if !double && value == 2 {
				double = true
				double_count += 1
			} else if !triple && value == 3 {
				triple = true
				triple_count += 1
			}

			// We are finished when two letters are encountered twice resp. trice
			if double && triple {
				break
			}
		}
	}

	fmt.Println(double_count * triple_count)
}
