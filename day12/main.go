package main

import (
	"bufio"
	"fmt"
	"go/types"
	"golang.org/x/text/encoding/charmap"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./day12/initial_state.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var initialState string
	for scanner.Scan() {
		initialState = scanner.Text()
	}

	file, err = os.Open("./day12/conversion.txt")
	check(err)
	defer file.Close()

	scanner = bufio.NewScanner(file)
	conversion := map[string]string{}
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, " => ")
		conversion[splitted[0]] = splitted[1]
	}

	fmt.Println("Part 1:", part1(initialState, conversion))
}

func part1(initialState string, conversion map[string]string) int {
	state := map[int]int32{}

	for index, char := range initialState {
		state[index] = char
	}
}
