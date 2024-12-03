package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func FirstPuzzle() int {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		os.Exit(1)
	}

	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) == 3 {
				num1, err1 := strconv.Atoi(match[1])
				num2, err2 := strconv.Atoi(match[2])
				if err1 != nil || err2 != nil {
					fmt.Println("Error converting string to int:", err1, err2)
					os.Exit(1)
				}
				result += num1 * num2
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return result
}

func SecondPuzzle() int {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}

	r, err := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	if err != nil {
		log.Fatalln(err)
	}

	res := r.FindAllString(input, -1)
	fmt.Println(res)

	total := 0
	disabled := false
	for _, text := range res {
		if text == "don't()" {
			disabled = true
			continue
		}

		if text == "do()" {
			disabled = false
			continue
		}

		if disabled {
			continue
		}

		var num1, num2 int
		fmt.Sscanf(text, "mul(%d,%d)", &num1, &num2)

		total += (num1 * num2)
	}
	return total
}
