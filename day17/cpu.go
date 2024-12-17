package day17

import "fmt"

type Cpu struct {
	A, B, C            int
	instructionPointer int
}

func (cpu *Cpu) adv(comboOp int) {
	op := getComboOpValue(cpu, comboOp)

	numerator := cpu.A
	denominator := pow(2, op)

	cpu.A = numerator / denominator

	cpu.instructionPointer += 2
}

func (cpu *Cpu) bxl(literalOp int) {
	cpu.B = cpu.B ^ literalOp
	cpu.instructionPointer += 2
}

func (cpu *Cpu) bst(comboOp int) {
	op := getComboOpValue(cpu, comboOp)
	cpu.B = op % 8
	cpu.instructionPointer += 2
}

func (cpu *Cpu) jnz(literalOp int) {
	if cpu.A != 0 {
		cpu.instructionPointer = literalOp
	} else {
		cpu.instructionPointer += 2
	}
}

func (cpu *Cpu) bxc(_ int) {
	cpu.B = cpu.B ^ cpu.C
	cpu.instructionPointer += 2
}

func (cpu *Cpu) out(comboOp int) {
	op := getComboOpValue(cpu, comboOp)
	fmt.Printf("%d,", op%8)

	cpu.instructionPointer += 2
}

func (cpu *Cpu) bdv(comboOp int) {
	op := getComboOpValue(cpu, comboOp)

	numerator := cpu.A
	denominator := pow(2, op)

	cpu.B = numerator / denominator

	cpu.instructionPointer += 2
}

func (cpu *Cpu) cdv(comboOp int) {
	op := getComboOpValue(cpu, comboOp)

	numerator := cpu.A
	denominator := pow(2, op)

	cpu.C = numerator / denominator

	cpu.instructionPointer += 2
}

func (cpu *Cpu) execute(instructions []int) {
	operations := map[int]func(op int){
		0: cpu.adv,
		1: cpu.bxl,
		2: cpu.bst,
		3: cpu.jnz,
		4: cpu.bxc,
		5: cpu.out,
		6: cpu.bdv,
		7: cpu.cdv,
	}

	for cpu.instructionPointer < len(instructions)-1 {
		operation := operations[instructions[cpu.instructionPointer]]
		operand := instructions[cpu.instructionPointer+1]

		operation(operand)
	}
}

func pow(base int, power int) int {
	res := 1

	for range power {
		res *= base
	}
	return res
}

func getComboOpValue(cpu *Cpu, comboOp int) int {
	op := 0
	if comboOp == 4 {
		op = cpu.A
	} else if comboOp == 5 {
		op = cpu.B
	} else if comboOp == 6 {
		op = cpu.C
	} else if comboOp == 7 {
		panic("Invalid combo operand in adv")
	} else {
		op = comboOp
	}
	return op
}
