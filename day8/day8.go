package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

type tree struct {
	children []tree
	metadata []int
}

func main() {
	exampleTree := ingestInput("./day8/input_example_part_1.txt")
	fmt.Println("Part 1 example input:", part1(exampleTree))
	fmt.Println("Part 2 example input:", part2(exampleTree))


	inputTree := ingestInput("./day8/input.txt")
	fmt.Println("Part 1:", part1(inputTree))
	fmt.Println("Part 2:", part2(inputTree))

}

func sortedKeys(input map[string]bool) []string {
	var keys []string
	for k := range input {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func ingestInput(filename string) tree {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	// Ingest the input
	scanner := bufio.NewScanner(file)
	var splitted []string
	for scanner.Scan() {
		text := scanner.Text()
		splitted = strings.Split(text, " ")
	}

	var input []int
	for _, s := range splitted {
		i, err := strconv.Atoi(s)
		check(err)
		input = append(input, i)
	}

	root, _ := createTree(input)

	return root
}

func part1(t tree) int {
	return sumMetadata(t)
}

func part2(t tree) int {
	s := 0
	if len(t.children) == 0 {
		s = sum(t.metadata)
	} else {
		for _, index := range t.metadata {
			if 0 < index && index <= len(t.children) {
				s += part2(t.children[index-1]) // Index is decremented to take convert it into a zero-based index
			}
		}
	}
	return s
}

// Creates a tree from the start of the input, returns the tree and the amount of the input which is consumed
func createTree(input []int) (tree, int) {
	childrenCount := input[0]
	metadataCount := input[1]

	var children []tree
	pointer := 2 // Track where each child starts in the input
	for i := 0; i < childrenCount; i++ {
		childTree, childLength := createTree(input[pointer: len(input)-metadataCount])
		pointer += childLength
		children = append(children, childTree)
	}

	metadata := input[pointer:pointer + metadataCount]

	return tree{children: children, metadata: metadata}, pointer + metadataCount
}

// Sums all metadata in a tree
func sumMetadata(t tree) int {
	sum := sum(t.metadata)
	for _, child := range t.children {
		sum += sumMetadata(child)
	}
	return sum
}

// Sums the values of an integer array
func sum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

