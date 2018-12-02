package main

import (
	"bufio"
	"fmt"
	"os"
)

var data []string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkForIndex(i int, c chan string) {
	strings := map[string]bool{} // Tracks all sub strings we have generated so far.

	// Loop through all strings and remove the i-th letter.
	for _, text := range data {
		withoutRemovedChar := text[:i] + text[i+1:] // Generate the sub string

		// Check if we have encountered this sub string before, if so we have found our result
		if strings[withoutRemovedChar] {
			c <- withoutRemovedChar
			return // We are finished, since we are guaranteed that there is only one occurrence.
		} else {
			strings[withoutRemovedChar] = true // Store the occurance of this sub string
		}
	}
}

func main() {
	file, err := os.Open("./day2/input.txt")
	check(err)
	defer file.Close()

	// Read all data from the file into an array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	// Concurrently check for every index in the string if there are similar IDs
	c := make(chan string)
	for i := 0; i < len(data[0]); i++ {
		go checkForIndex(i, c)
	}

	fmt.Println(<-c)
}
