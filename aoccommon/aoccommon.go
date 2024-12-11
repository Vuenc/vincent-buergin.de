package aoccommon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ReadLines(filename string) []string {
	file, _ := os.Open(filename)
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func Array2D(height, width int, defaultValue ...int) [][]int {
	a := make([][]int, height)
	for i := range a {
		a[i] = make([]int, width)
	}
	if len(defaultValue) > 0 {
		for y, row := range a {
			for x := range row {
				a[y][x] = defaultValue[0]
			}
		}
	}
	return a
}

func SplitToInts(input, separator string) []int {
	parts := strings.Split(input, separator)
	var output []int

	for _, val := range(parts) {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Errorf("invalid number: %s", parts[0]))
		}
		output = append(output, num)
	}

	return output
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a % b)
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elements []T) Set[T] {
    set := Set[T]{}
    for _, e := range elements {
        set[e] = struct{}{}
    }
    return set
}

func Union[T comparable](set1, set2 Set[T]) Set[T] {
    result := NewSet([]T{})

    for elem := range set1 {
        result[elem] = struct{}{}
    }
    for elem := range set2 {
        result[elem] = struct{}{}
    }

    return result
}