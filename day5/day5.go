package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func isValidUpdate(update []int, rules [][]int) bool {
	var forbiddenPages []int
	for _, page := range update {
		// fmt.Println(forbiddenPages)
		if slices.Contains(forbiddenPages, page) {
			return false
		}
		for _, rule := range rules {
			// fmt.Println(rule, rule[1], page)
			if rule[1] == page {
				forbiddenPages = append(forbiddenPages, rule[0])
			}
		}
	}
	return true
}

func star1(updates, rules [][]int) {
	var middlePageNumberSum = 0
	for _, update := range updates {
		if isValidUpdate(update, rules) {
			middlePageNumberSum += update[(len(update)-1)/2]
		}
	}
	fmt.Println("Sum of middle page numbers of valid updates:", middlePageNumberSum)
}

func violatedRules(update []int, rules [][]int) [][]int {
	var activeRules [][]int
	var violatedRules [][]int

	for _, page := range update {
		for _, rule := range activeRules {
			if page == rule[0] {
				violatedRules = append(violatedRules, rule)
			}
		}
		for _, rule := range rules {
			if rule[1] == page {
				activeRules = append(activeRules, rule)
			}
		}
	}
	return violatedRules
}

func fixUpdate(update []int, rules [][]int) {
	var violated = violatedRules(update, rules)
	// fmt.Println("Invalid:", update)
	for len(violated) > 0 {
		// fmt.Println("Violated rules:", violated)
		var movedNumbers []int
		for _, rule := range violated {
			if !slices.Contains(movedNumbers, rule[0]) {
				movedNumbers = append(movedNumbers, rule[0])
				i := slices.Index(update, rule[0])
				j := slices.Index(update, rule[1])
				update = append(update[:j], append([]int{rule[0]}, append(update[j:i], update[i+1:]...)...)...)
				// fmt.Println("-> ", update)
			}
		}
		violated = violatedRules(update, rules)
	}
	// fmt.Println("Done.")
}

func star2(updates, rules [][]int) {
	var middlePageNumberSum = 0
	for _, update := range updates {
		if !isValidUpdate(update, rules) {
			fixUpdate(update, rules)
			// fmt.Println("Fixed:", update)
			middlePageNumberSum += update[(len(update)-1)/2]
		}
	}
	fmt.Println("Sum of middle page numbers of fixed updates:", middlePageNumberSum)
}

func main() {
	var lines = common.ReadLines("day5.txt")

    var re = regexp.MustCompile(`(\d+)\|(\d+)`)
	var rules [][]int
	var i = 0
	for ; lines[i] != ""; i++ {
		var nums = re.FindStringSubmatch(lines[i])
		var val1, _ = strconv.Atoi(nums[1])
		var val2, _ = strconv.Atoi(nums[2])
		rules = append(rules, []int{ val1, val2 })
	}
	i++
	var updates [][]int
	for ; i < len(lines); i++ {
		updates = append(updates, common.SplitToInts(lines[i], ","))
	}

	star1(updates, rules)
	star2(updates, rules)
}