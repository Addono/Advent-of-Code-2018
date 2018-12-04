package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type activity struct {
	minute int
	move string
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


func main() {
	file, err := os.Open("./day4/input.txt")
	check(err)
	defer file.Close()


	// Ingest all input
	var source []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		source = append(source, scanner.Text())
	}

	sort.Strings(source)

	var input []activity
	for _, text := range source {
		// Extract the text
		x := strings.Split(text, " ")
		minutes, _ := strconv.Atoi(strings.Split(x[1], ":")[1][:2])

		input = append(input, activity{
			minute: minutes,
			move: x[3],
		})
	}

	count := [4000]int{}
	sleepTime := [4000][60]int{}
	var guard int
	for i, a := range input {
		if a.move == "asleep" {
			asleep := a.minute
			up := input[i+1].minute

			count[guard] += up - asleep
			for i := asleep; i < up; i++ {
				sleepTime[guard][i] += 1
			}
		} else if a.move != "up" {
			guard, err = strconv.Atoi(a.move[1:])
			check(err)
		}
	}

	mostSleepingGuard := guard
	for g, minutes := range count {
		if (minutes >= count[mostSleepingGuard]) {
			mostSleepingGuard = g
		}
	}

	mostSleptMinute, _ := getMostSleptMinute(&sleepTime[mostSleepingGuard])

	fmt.Println(mostSleptMinute * mostSleepingGuard)

	// Part 2
	mostSleepingGuardByMinute := 0
	mostSleptMinute = 0
	mostSleptMinuteDuration := 0
	for g, st := range sleepTime {
		maxMinute, maxDuration := getMostSleptMinute(&st)

		if (maxDuration > mostSleptMinuteDuration) {
			mostSleptMinuteDuration = maxDuration
			mostSleptMinute = maxMinute
			mostSleepingGuardByMinute = g
		}
	}

	fmt.Println(mostSleepingGuardByMinute * mostSleptMinute)
}

func getMostSleptMinute(sleepTime *[60]int) (int, int) {
	mostSleptMinute := -1
	mostSleptMinuteAmount := -1
	for minute, amount := range sleepTime {
		if (amount > mostSleptMinuteAmount) {
			mostSleptMinuteAmount = amount
			mostSleptMinute = minute
		}
	}
	return mostSleptMinute, mostSleptMinuteAmount
}
