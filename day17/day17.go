package main

import (
	common "aoc2024/aoccommon"
	"fmt"
	"strconv"
)

func comboOperand(operand int64, registers [3]int64) int64 {
	if operand <= 3 {
		return operand
	}
	return registers[operand - 4]
}

func star1(program []int64, registerA int64) {
	registers := [3]int64{registerA, 0, 0}
	output := []int64{}
	for instructionPointer := int64(0); instructionPointer < int64(len(program)); instructionPointer += 2 {
		instruction, operand := program[instructionPointer], program[instructionPointer+1]
		switch instruction {
		case 0: // adv instruction
			denominator := int64(1) << comboOperand(operand, registers)
			registers[0] = registers[0] / denominator
		case 1: // bxl instruction
			registers[1] = registers[1] ^ operand
		case 2: // bst instruction
			registers[1] = comboOperand(operand, registers) % 8
		case 3: // jnz instruction
			if registers[0] != 0 {
				instructionPointer = operand - 2
			}
		case 4: // bxc instruction
			registers[1] = registers[1] ^ registers[2]
		case 5: // out instruction
			output = append(output, comboOperand(operand, registers) % 8)
		case 6: // bdv instruction
			denominator := int64(1) << comboOperand(operand, registers)
			registers[1] = registers[0] / denominator
		case 7: // cdv instruction
			denominator := int64(1) << comboOperand(operand, registers)
			registers[2] = registers[0] / denominator
		}
	}
	fmt.Print("Program output: ")
	for i, out := range output {
		fmt.Print(out)
		if i < len(output) - 1 {
			fmt.Print(",")
		}
	}
	fmt.Println()
}

func star2_recursive(program []int64, i int, registerA int64) (int64, bool) {
	if (i < 0) {
		return registerA, true
	}
	code := program[i]
	for aThreeBits := range int64(8) {
		b := aThreeBits ^ 1
		c := (((registerA << 3) | aThreeBits) >> b)
		b = (b ^ c ^ 4) % 8
		if b == code {
			solution, ok := star2_recursive(program, i-1, (registerA << 3) | aThreeBits)
			if ok {
				return solution, true
			}
		}
	}
	return -1, false
}

func star2(program []int64) int64 {
	registerA, ok := star2_recursive(program, len(program)-1, int64(0))
	if ok {
		fmt.Println("Initial register A value for quine: ", registerA)
	}
	return registerA
}

func main() {
	var lines = common.ReadLines("day17.txt")

	_registerA, _ := strconv.Atoi(lines[0][len("Register A: "):])
	registerA := int64(_registerA)
	_program := common.SplitToInts(lines[4][len("Program: "):], ",")
	program := []int64{}
	for _, val := range _program {
		program = append(program, int64(val))
	}

	star1(program, registerA)
	registerA = star2(program)
	fmt.Print("Program:        ")
	for i, code := range program {
		fmt.Print(code)
		if i < len(program) - 1 {
			fmt.Print(",")
		}
	}
	fmt.Println()
	star1(program, registerA)
}