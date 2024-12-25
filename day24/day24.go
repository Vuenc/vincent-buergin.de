package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"math/bits"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Rule struct {
	operand1  string
	operand2  string
	operation string
	target    string
	role      string
	id int
}

func evaluateCircuit(registerValues map[string]bool, rulesIn []Rule) int {
	// fmt.Println(rules)
	rules := append([]Rule{}, rulesIn...)
	// outerLoop:
	for len(rules) > 0 {
		for i := 0; i < len(rules); i++ {
			rule := rules[i]
			op1, exists1 := registerValues[rule.operand1]
			op2, exists2 := registerValues[rule.operand2]
			if exists1 && exists2 {
				if rule.operation == "AND" {
					registerValues[rule.target] = op1 && op2
				} else if rule.operation == "OR" {
					registerValues[rule.target] = op1 || op2
				} else if rule.operation == "XOR" {
					registerValues[rule.target] = op1 != op2
				} else {
					panic("Unknown operation: " + rule.operation)
				}
				rules[i] = rules[len(rules)-1]
				rules = rules[:len(rules)-1]
				
				// fmt.Printf("%s %s %s -> %s\n", rule.operand1, rule.operation, rule.operand2, rule.target)
			}
		}
		// fmt.Println()
	}
	outputNumber := 0
	for name, value := range registerValues {
		if name[0] == 'z' && value {
			bit, _ := strconv.Atoi(name[1:])
			outputNumber += 1 << bit
		}
	}
	return outputNumber
}

// z00 = x00 XOR y00
// c00 = x00 AND y00
// z01 = x01 XOR y01 XOR c00
// c01 = (x01 AND y01) OR ((x01 OR y01) AND c00)



func assignRoles(rulesIn []Rule) ([]Rule, map[string]string) {
	rules := append([]Rule{}, rulesIn...)
	// x01 AND y01 =: A01
	// x01 OR y01 =: B01
	// B01 AND D00 =: C01
	// A01 OR C01 =: D01
	// x01 XOR y01 =: E01
	// z01 = D01 XOR E01
	roles := map[string]string{}

	hasChanges := true
	for hasChanges {
		hasChanges = false
		for i := range rules {
			rule := rules[i]
			if rule.role != "" {continue}
			if rule.operand1[0] == 'x' && rule.operand2[0] == 'y' && rule.operand1[1:] == rule.operand2[1:] {
				if rule.operation == "AND" {
					rules[i].role = "A" + rule.operand1[1:]
				// }
				// else if rule.operation == "OR" {
				// 	rules[i].role = "B" + rule.operand1[1:]
				// }
				} else if rule.operation =="XOR" {
					rules[i].role = "E" + rule.operand1[1:]
				}
				roles[rule.target] = rules[i].role
				hasChanges = true
			} else if roles[rule.operand1] != "" && roles[rule.operand2] != "" {
				if roles[rule.operand1] == "E01" {fmt.Println(rule)}
				if roles[rule.operand1] > roles[rule.operand2] {
					rules[i].operand1, rules[i].operand2 = rules[i].operand2, rules[i].operand1
					fmt.Println("Swap:", rules[i].operand1, rules[i].operand2 )
					rule = rules[i]
					hasChanges = true
				}
				role1, role2 := roles[rule.operand1], roles[rule.operand2]
				if roles[rule.operand1] == "E01" || roles[rule.operand1] == "A00" {fmt.Println(rule, role1, role2, role1 == "A00" && role2 == "E01" && rule.operation == "AND")}
				if role1[0] == 'D' && role2[0] == 'E' && role1[1:] == role2[1:] && rule.operation == "AND" {
					rules[i].role = "C" + role1[1:]
					roles[rules[i].target] = rules[i].role
					hasChanges = true
				} else if role1[0] == 'A' && role2[0] == 'C' && role1[1:] == role2[1:] && rule.operation == "OR" {
					num, _ := strconv.Atoi(role1[1:])
					rules[i].role = "D" + fmt.Sprintf("%02d", num+1)
					roles[rules[i].target] = rules[i].role
					hasChanges = true
				} else if role1 == "A00" && role2 == "E01" && rule.operation == "AND" {
					rules[i].role = "C01"
					roles[rules[i].target] = rules[i].role
					hasChanges = true
				}
				// } else if role1 == 
				// if roles[rule.operand1][0] == 'B' && roles[rule.operand2][0] == 'D' {
				// 	num1, _ := strconv.Atoi(roles[rule.operand1][1:])
				// 	num2, _ := strconv.Atoi(roles[rule.operand2][1:])
				// 	if num1 == num2 + 1 {
				// 		rules[i].role = "D" + roles[rule.operand1][1:]
				// 		roles[rule.target] = rules[i].role
				// 		hasChanges = true
				// 	}
				// }
			}
		}
	}

	// for _, rule := range rules {
	// 	o1 := rule.operand1
	// 	if roles[o1] != "" {o1 = roles[o1]}
	// 	o2 := rule.operand2
	// 	if roles[o2] != "" {o2 = roles[o2]}
	// 	t := rule.target
	// 	if rule.role != "" {t = rule.role}
	// 	fmt.Printf("%s %s %s -> %s\n", o1, rule.operation, o2, t)
	// }

	fmt.Println(roles)
	return rules, roles
}

func evaluateBackwards(registerValues map[string]bool, rulesIn []Rule, roles map[string]string) {
	// fmt.Println(rules)
	rules := append([]Rule{}, rulesIn...)
	requestedValues := []string{}
	requestedPreviously := []string{}

	fmt.Println(rules)

	for i := range 45 {
		requestedValues = []string{fmt.Sprintf("z%02d", i)}
		fmt.Println("Requested", fmt.Sprintf("z%02d", i))
		numChanges := 0
		hasChanges := true
		processedRules := []int{}
		for hasChanges {
			hasChanges = false
			for i := 0; i < len(rules); i++ {
				rule := rules[i]
				if (slices.Contains(requestedPreviously, rule.target) || slices.Contains(requestedValues, rule.target)) && !slices.Contains(processedRules, i) {
					processedRules = append(processedRules, i)
					// requestedValues[slices.Index(requestedValues, rule.target)] = requestedValues[len(requestedValues)-1]
					// requestedValues = requestedValues[:len(requestedValues)-1]

					if !slices.Contains(requestedPreviously, rule.operand1) {
						o1 := rule.operand1
						if roles[o1] != "" {o1 = roles[o1]}
						o2 := rule.operand2
						if roles[o2] != "" {o2 = roles[o2]}
						t := rule.target
						if rule.role != "" {t = rule.role}
						// fmt.Println(rule.operand1, o1, rule.operand2, o2, rule.target, t)
						fmt.Println("Requested", o1, 
							fmt.Sprintf("(%s = %s %s %s)", t, o1, rule.operation, o2))
						if !slices.Contains(requestedValues, rule.operand1) {
						// if rule.operand1[0] != 'x' && rule.operand1[0] != 'y' {
							requestedValues = append(requestedValues, rule.operand1)
							hasChanges = true
							numChanges++
						}
					}
					// }

					// if rule.operand2[0] != 'x' && rule.operand2[0] != 'y' {

					if !slices.Contains(requestedPreviously, rule.operand1) {
							o1 := rule.operand1
							if roles[o1] != "" {o1 = roles[o1]}
							o2 := rule.operand2
							if roles[o2] != "" {o2 = roles[o2]}
							t := rule.target
							if rule.role != "" {t = rule.role}
							fmt.Println("Requested", o2, 
								fmt.Sprintf("(%s = %s %s %s)", t, o1, rule.operation, o2))
						if !slices.Contains(requestedValues, rule.operand2) {
							requestedValues = append(requestedValues, rule.operand2)
							hasChanges = true
							numChanges++
						}
					}
					// }
				}
			}
		}
		fmt.Println(numChanges)
		requestedPreviously = append(requestedPreviously, requestedValues...)
	}

}

func printRules(rules []Rule, roles map[string]string) {
	for i := range 45 {
		requestedValues := []string{fmt.Sprintf("z%02d", i)}
		depth := 0
		nextRequestedValues := []string{}
		relevantRules := map[string]Rule{}
		
		for depth < 4 {
			depth++
			for i := 0; i < len(rules); i++ {
				rule := rules[i]
				if slices.Contains(requestedValues, rule.target) {
					nextRequestedValues = append(nextRequestedValues, rule.operand1)
					nextRequestedValues = append(nextRequestedValues, rule.operand2)
					relevantRules[rule.target] = rule
				}
			}
			requestedValues = nextRequestedValues 
		}

		var printRule func(rule Rule) string
		printRule = func(rule Rule) string {
			if rule.target == rule.operand1 || rule.target == rule.operand2 {
				return fmt.Sprintf("@@%s[%03d]=%s %s %s", rule.target, rule.id, rule.operand1, rule.operation, rule.operand2)
			}
			o1 := rule.operand1
			r1, exists := relevantRules[o1]
			if exists {o1 = "(" + printRule(r1) + ")"
			} else if roles[o1] != "" {o1 = roles[o1]}
			o2 := rule.operand2
			r2, exists := relevantRules[o2]
			if exists {o2 = "(" + printRule(r2) + ")"
			} else if roles[o2] != "" {o2 = roles[o2]}
			t := rule.target
			if roles[t] != "" {t = roles[t]}
			// return fmt.Sprintf("%s=%s %s %s", t, o1, rule.operation, o2)
			return fmt.Sprintf("%s[%03d]=%s %s %s", t, rule.id, o1, rule.operation, o2)

			// return fmt.Sprintf("%s %s %s", o1, rule.operation, o2)
		}
		// fmt.Println(relevantRules)
		rule := relevantRules[fmt.Sprintf("z%02d", i)]
		fmt.Printf("%s = %s\n", rule.target, printRule(rule))
	}
}

func star1(registerValues map[string]bool, rules []Rule) {
	fmt.Println("Circuit output:", evaluateCircuit(registerValues, rules))
}

func badness(rules []Rule) int {
	summedHammingDistance := 0
	for _, y := range []int{0, (1 << 46) - 1} {
		registerValues := map[string]bool{}
		for t := range 45 {
			x := 1 << t
			for i := range 45 {
				registerValues[fmt.Sprintf("x%02d", i)] = ((x >> i) & 1) == 1
				registerValues[fmt.Sprintf("y%02d", i)] = ((y >> i) & 1) == 1
			}
			// fmt.Println(x, registerValues)
			out := evaluateCircuit(registerValues, rules)
			// fmt.Println(out)
			summedHammingDistance += bits.OnesCount64(uint64(out))
		}
	}
	return summedHammingDistance
}

func star2(registerValues map[string]bool, rules []Rule) {
	rules[36].target, rules[135].target = rules[135].target, rules[36].target
	rules[51].target, rules[55].target = rules[55].target, rules[51].target
	rules[121].target, rules[169].target = rules[169].target, rules[121].target
	rules[94].target, rules[64].target = rules[64].target, rules[94].target
	// rules[185].target, rules[51].target = rules[51].target, rules[185].target
	// rules[170].target, rules[8].target = rules[8].target, rules[170].target


	// rules[55].target, rules[8].target = rules[8].target, rules[55].target

	rules, roles := assignRoles(rules)
	for i, rule := range rules {
		// if roles[rule.operand1] != "" && roles[rule.operand2] != "" && roles[rule.operand1] > roles[rule.operand2] {
		if roles[rule.operand2] != "" && roles[rule.operand2][0] == 'E' {
			rules[i].operand2, rules[i].operand1 = rules[i].operand1, rules[i].operand2
		} else if roles[rule.operand2] != "" && roles[rule.operand2][0] == 'A' && !(roles[rule.operand1] != "" && roles[rule.operand1][0] == 'E') {
			rules[i].operand2, rules[i].operand1 = rules[i].operand1, rules[i].operand2
		} 
	}
	// evaluateBackwards(registerValues, rules, roles)
	printRules(rules, roles)
	// registerNames := []string
	// for name := range registerValues {
	// 	if name[0] != 'x' && name[0] != 'y' {
	// 		registerNames = append(registerNames, name)
	// 	}
	// }

	// bestBadness := badness(rules)
	// bestBadnessSwap := [2]int{0, 0}
	// for badness(rules) > 0 {
	// 	for i := range rules {
	// 		for j := range rules[i+1:] {
	// 			rules[i].target, rules[j].target = rules[j].target, rules[i].target
	// 			// fmt.Println(rules[i])
	// 			swapBadness := badness(rules)
	// 			// fmt.Println(bestBadness, swapBadness)
	// 			if swapBadness < bestBadness {
	// 				bestBadness = swapBadness
	// 				bestBadnessSwap = [2]int{i, j}
	// 				fmt.Println("New best badness:", bestBadness, "for swap", bestBadnessSwap)
	// 			}
	// 			rules[i].target, rules[j].target = rules[j].target, rules[i].target
	// 		}
	// 	}
	// 	fmt.Println("Swapping", bestBadnessSwap)
	// 	rules[bestBadnessSwap[0]].target, rules[bestBadnessSwap[1]].target = rules[bestBadnessSwap[1]].target, rules[bestBadnessSwap[0]].target
	// }
	// y := 0
	// lastOut := -1
	// for x := range 45 {
	// 	x = 1 << x
	// 	for i := range 45 {
	// 		// fmt.Println(i, ((x >> i) & 1), x, x >> i)
	// 		registerValues[fmt.Sprintf("x%02d", i)] = ((x >> i) & 1) == 1
	// 		// fmt.Println(fmt.Sprintf("x%02d", x), registerValues[fmt.Sprintf("x%02d", x)], ((x >> i) & 1) == 1)
	// 		registerValues[fmt.Sprintf("y%02d", i)] = ((y >> i) & 1) == 1
	// 	}
	// 	// registerValues["x00"] = true
	// 	// fmt.Println(registerValues["x00"], fmt.Sprintf("x%02d", 0), x >> 0, (x >> 0) & 1, (x >> 0) & 1 == 1)

	// 	fmt.Printf("  %046b\n", x)
	// 	fmt.Printf("+ %046b\n", y)
	// 	// fmt.Printf("= %046b\n", x+y)
	// 	out := evaluateCircuit(registerValues, rules)
	// 	fmt.Printf("≠ %046b\n", out)
	// 	if lastOut != -1 {
	// 		fmt.Printf("d %046b\n", out ^ lastOut)
	// 	}
	// 	lastOut = out

	// 	fmt.Println()
	// 	// ignore(x, y)
	// }

	swapIndices := []int{36, 135, 51, 55, 121, 169, 94, 64};
	names := []string{}
	for _, i := range swapIndices {
		names = append(names, rules[i].target)
	}
	sort.Strings(names)
	fmt.Println(strings.Join(names, ","))
}

func main() {
	var lines = common.ReadLines("day24.txt")
	registerValues := map[string]bool{}
	rules := []Rule{}
	i := 0
	for ; lines[i] != ""; i++ {
		val, _ := strconv.Atoi(lines[i][5:])
		registerValues[lines[i][:3]] = val == 1
	}
	for i++; i < len(lines); i++ {
		split := strings.Split(lines[i], " ")
		rules = append(rules, Rule{min(split[0], split[2]), max(split[0], split[2]), split[1], split[4], "", len(rules)})
	}
	star1(registerValues, rules)
	star2(registerValues, rules)
}
