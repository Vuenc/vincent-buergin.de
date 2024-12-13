package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"regexp"
	"strconv"
)

type claw_machine struct {
	x1, y1, x2, y2 int
	xPrize, yPrize int
}

func star1(machines []claw_machine, star2Mode ...bool) {
	// (a x1 + b x2) * y2 - (a y1 + b y2) * x2 = xp * y2 + yp * x2
	// a (x1 y2 - y1 x2) = xp y2 + yp x2
	// a = (xp y2 + yp x2) / (x1 y2 - y1 x2)
	// (a x1 + b x2) * y1 - (a y1 + b y2) * x1 = xp * y1 + yp * x1
	// b (x2 y1 - y2 x1) = xp y1 + yp x1
	// b = (xp y1 + yp x1) / (x2 y1 - y2 x1)

	tokenSum := 0
	for _, machine := range machines {
		detA := (machine.x1 * machine.y2 - machine.y1 * machine.x2)
		detB := (machine.x2 * machine.y1 - machine.y2 * machine.x1)
		if detA != 0 && detB != 0 {
			a := (machine.xPrize * machine.y2 - machine.yPrize * machine.x2) / detA
			b := (machine.xPrize * machine.y1 - machine.yPrize * machine.x1) / detB
			if a >= 0 && b >= 0 && machine.x1 * a + machine.x2 * b == machine.xPrize && machine.y1 * a + machine.y2 * b == machine.yPrize {
				tokenSum += 3 * a + b
			}
		}
	}
	isStar2Mode := (len(star2Mode) == 1 && star2Mode[0])
	if !isStar2Mode {
		fmt.Println("Num tokens:", tokenSum)
	} else {
		fmt.Println("Num tokens (+10000000000000):", tokenSum)
	}
}

func star2(machines []claw_machine) {
	for i := range machines {
		machines[i].xPrize += 10000000000000
		machines[i].yPrize += 10000000000000
	}
	star1(machines, true)
}

func main() {
	var lines = common.ReadLines("day13.txt")

    buttonARegex:= regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
    buttonBRegex:= regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
    prizeRegex:= regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	var machines []claw_machine
	for i := 0; i < len(lines); i+=4 {
		machine := claw_machine{}
		numsA := buttonARegex.FindStringSubmatch(lines[i])
		machine.x1, _ = strconv.Atoi(numsA[1])
		machine.y1, _ = strconv.Atoi(numsA[2])
		numsB := buttonBRegex.FindStringSubmatch(lines[i+1])
		machine.x2, _ = strconv.Atoi(numsB[1])
		machine.y2, _ = strconv.Atoi(numsB[2])
		numsPrize := prizeRegex.FindStringSubmatch(lines[i+2])
		machine.xPrize, _ = strconv.Atoi(numsPrize[1])
		machine.yPrize, _ = strconv.Atoi(numsPrize[2])

		machines = append(machines, machine)
	}

	// star1([]claw_machine{{94, 34, 22, 67, 8400, 5400}})
	star1(machines)

	// star2([]claw_machine{{26, 66, 67, 21, 12748, 12176}})
	star2(machines)
}