package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func FirstPuzzle() int {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	rules := map[int][]int{}

	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		first, second := parts[0], parts[1]

		key, err := strconv.Atoi(first)
		if err != nil {
			log.Fatalln(err)
		}

		value, err := strconv.Atoi(second)
		if err != nil {
			log.Fatalln(err)
		}

		rules[key] = append(rules[key], value)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		update := []int{}

		for _, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalln(err)
			}
			update = append(update, val)
		}

		if validUpdate(update, rules) {
			result += update[len(update)/2]
		}
	}

	return result
}

func validUpdate(update []int, rules map[int][]int) bool {
	for i := 1; i < len(update); i++ {
		key := update[i]
		for j := 0; j <= i; j++ {
			val := update[j]
			if contains(rules[key], val) {
				return false
			}
		}
	}
	return true
}

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func SecondPuzzle() int {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	rules := map[int][]int{}

	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		first, second := parts[0], parts[1]

		key, err := strconv.Atoi(first)
		if err != nil {
			log.Fatalln(err)
		}

		value, err := strconv.Atoi(second)
		if err != nil {
			log.Fatalln(err)
		}

		rules[key] = append(rules[key], value)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		update := []int{}

		for _, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalln(err)
			}
			update = append(update, val)
		}

		if !validUpdate(update, rules) {
			sortUpdateElements(update, rules)
			result += update[len(update)/2]
		}
	}

	return result
}

// returns true if the first comes before the second
func compareUpdateElement(first int, second int, rules map[int][]int) bool {
	return !contains(rules[second], first)
}

func sortUpdateElements(update []int, rules map[int][]int) {
	n := len(update)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if !compareUpdateElement(update[i], update[j], rules) {
				update[i], update[j] = update[j], update[i]
			}
		}
	}
}
