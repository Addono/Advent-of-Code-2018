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

	count := 0
	encountered := map[int]bool{}

	for {
		file, err := os.Open("./day1/input.txt")
		check(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			amount, err := strconv.Atoi(scanner.Text())
			check(err)

			count += amount

			_, ok := encountered[count]
			if ok {
				fmt.Println(count)
				return
			} else {
				encountered[count] = true
			}
		}
	}
}