package day17

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) (Cpu, []int) {
	var cpu Cpu
	var instructions []int

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	cpu.A = parseRegister(scanner.Text())
	scanner.Scan()
	cpu.B = parseRegister(scanner.Text())
	scanner.Scan()
	cpu.C = parseRegister(scanner.Text())

	scanner.Scan()
	scanner.Scan()
	program := scanner.Text()

	parts := strings.Split(program, ": ")

	nums := parts[1]
	splitNums := strings.Split(nums, ",")

	for _, num := range splitNums {
		instruction, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction)
	}

	return cpu, instructions
}

func parseRegister(line string) int {
	parts := strings.Split(line, ": ")
	val := parts[1]

	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	return intVal
}

func FirstPuzzle() int {
	cpu, instructions := parseInput("day17/input.txt")

	cpu.execute(instructions)

	fmt.Println()

	return 0
}

func (cpu *Cpu) executeForOutput(instructions []int) []int {
	var output []int
	operations := map[int]func(op int, out *[]int){
		0: func(op int, _ *[]int) { cpu.adv(op) },
		1: func(op int, _ *[]int) { cpu.bxl(op) },
		2: func(op int, _ *[]int) { cpu.bst(op) },
		3: func(op int, _ *[]int) { cpu.jnz(op) },
		4: func(op int, _ *[]int) { cpu.bxc(op) },
		5: func(op int, out *[]int) {
			op = getComboOpValue(cpu, op)
			*out = append(*out, op%8)
			cpu.instructionPointer += 2
		},
		6: func(op int, _ *[]int) { cpu.bdv(op) },
		7: func(op int, _ *[]int) { cpu.cdv(op) },
	}

	for cpu.instructionPointer < len(instructions)-1 {
		operation := operations[instructions[cpu.instructionPointer]]
		operand := instructions[cpu.instructionPointer+1]
		operation(operand, &output)
	}

	return output
}

func SecondPuzzle() int {
	cpu, instructions := parseInput("day17/input.txt")

	valid := []int{0}

	for length := 1; length < len(instructions)+1; length++ {
		previousValid := valid
		valid = []int{}

		for _, num := range previousValid {
			for offset := 0; offset < 8; offset++ {
				candidateA := 8*num + offset

				testCpu := Cpu{
					A:                  candidateA,
					B:                  cpu.B,
					C:                  cpu.C,
					instructionPointer: 0,
				}

				output := testCpu.executeForOutput(instructions)

				targetSuffix := instructions[len(instructions)-length:]
				if isValidOutput(output, targetSuffix) {
					valid = append(valid, candidateA)
				}
			}
		}
	}

	if len(valid) == 0 {
		return -1
	}

	answer := valid[0]
	for _, val := range valid {
		if val < answer {
			answer = val
		}
	}

	return answer
}

func isValidOutput(output []int, target []int) bool {
	if len(output) < len(target) {
		return false
	}

	for i := range target {
		if output[i] != target[i] {
			return false
		}
	}

	return true
}
