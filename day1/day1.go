package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// The code for day 1, star 1 was generated with ChatGPT, as I'd never written Go before.
// Prompt (along the lines of):
// "Create a Go program that reads a file, splits into lines,
// splits by whitespace, writes the first and second item as numbers
// into two separate lists, sorts the lists, and computes the sums of
// pairwise distances between entries in the two lists."

// Function to read the file and process the numbers
func processFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list1, list2 []int
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by whitespace
		parts := strings.Fields(line)

		if len(parts) < 2 {
			continue // Skip lines that do not have at least two numbers
		}

		// Convert the first and second parts to integers
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", parts[0])
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", parts[1])
		}

		// Append the numbers to their respective lists
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return list1, list2, nil
}

// Function to calculate the sum of distances between corresponding elements in two lists
func calculateDistanceSum(list1, list2 []int) int {
	if len(list1) != len(list2) {
		fmt.Println("The lists must have the same length.")
		return 0
	}

	var sum int
	for i := 0; i < len(list1); i++ {
		sum += abs(list1[i] - list2[i])
	}

	return sum
}

// Function to return the absolute value of an integer
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func star1(list1, list2 []int) {
	// Sort the lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Calculate the sum of distances
	distanceSum := calculateDistanceSum(list1, list2)

	// Output the results
	// fmt.Println("Sorted List 1:", list1)
	// fmt.Println("Sorted List 2:", list2)
	fmt.Println("Sum of distances:", distanceSum)
}

func star2(list1, list2 []int) {
	var j = 0
	var score = 0
	for _, val := range list1 {
		for j < len(list2) && list2[j] < val {
			j++
		}
		for j < len(list2) && list2[j] == val {
			score += val
			j++
		}
	}
	fmt.Println("Score:", score)
}

func main() {
	// Replace with the path to your file
	filename := "day1.txt"

	// Process the file
	list1, list2, err := processFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Sort the lists
	sort.Ints(list1)
	sort.Ints(list2)

	star1(list1, list2);
	star2(list1, list2);
}