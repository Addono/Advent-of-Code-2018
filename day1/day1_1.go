package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./day1/input.txt")
	check(err)
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		amount, err := strconv.Atoi(scanner.Text())
		check(err)

		count += amount
	}

	fmt.Println(count)
}