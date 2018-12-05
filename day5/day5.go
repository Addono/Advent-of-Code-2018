package main

import (
	"bufio"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"os"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func react(c1 int32, c2 int32) bool {
	return c1 + 32 == c2 || c2 + 32 == c1
}

func main() {
	file, err := os.Open("./day5/input.txt")
	check(err)
	defer file.Close()

	// Ingest the input
	var source string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		source = scanner.Text()
	}

	remainingLength := reduce(source, 0)

	fmt.Println(remainingLength)

	// Part 2
	minLength := len(source)
	for i := "a"[0]; i <= "z"[0]; i++ {
		length := reduce(source, int32(i))
		if (length < minLength) {
			minLength = length
		}
	}

	fmt.Println(minLength)
}

func reduce(polymer string, ignore int32) int {
	var leftover stack.Stack
	for _, char := range polymer {
		if char == ignore || char + 32 == ignore {
			// Do nothing
		} else if leftover.Len() == 0 || !react(leftover.Peek().(int32), char) {
			leftover.Push(char)
		} else {
			leftover.Pop()
		}
	}
	return leftover.Len()
}
