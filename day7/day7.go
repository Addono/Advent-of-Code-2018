package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Part 1 example input:", computePath("./day7/input_example_part_1.txt"))
	fmt.Println("Part 1:", computePath("./day7/input.txt"))
}

func sortedKeys(input map[string]bool) []string {
	var keys []string
	for k := range input {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func computePath(filename string) string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	// Ingest the input
	dependencies := map[string]map[string]bool{}
	nodes := map[string]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, " ")

		dependency := splitted[1]
		dependant := splitted[7]
		if dependencies[dependant] == nil {
			dependencies[dependant] = map[string]bool{}
		}
		dependencies[dependant][dependency] = true
		nodes[dependency] = true
		nodes[dependant] = true // Not sure if this is entirely necessary, at least can we be sure that we won't start with these
	}

	output := ""
	progress := true
	for len(nodes) > 0 {
		progress = false
		for _, k := range sortedKeys(nodes) {
			if len(dependencies[k]) == 0 {
				output += k
				delete(nodes, k)
				for _, d := range dependencies {
					delete(d, k)
				}
				progress = true
				break
			}
		}

		if !progress {
			fmt.Println("Could not solve dependencies, are they cyclic?")
			break
		}
	}

	return output
}

