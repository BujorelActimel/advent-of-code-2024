package day7

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	target int
	nums   []int
}

func parseInput(fileName string) []Pair {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res []Pair

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(parts[0])
		stringNums := strings.Fields(parts[1])
		var nums []int
		for _, num := range stringNums {
			intNum, _ := strconv.Atoi(num)
			nums = append(nums, intNum)
		}
		res = append(res, Pair{target, nums})
	}
	return res
}

func FirstPuzzle() int {
	input := parseInput("day7/input.txt")
	result := 0
	for _, pair := range input {
		if canReachTarget(pair.nums, pair.target) {
			result += pair.target
		}
	}
	return result
}

func canReachTarget(nums []int, target int) bool {
	return tryOperations(nums, 1, nums[0], target)
}

func tryOperations(nums []int, pos int, current int, target int) bool {
	if pos == len(nums) {
		return current == target
	}

	// Try adunare
	if tryOperations(nums, pos+1, current+nums[pos], target) {
		return true
	}

	// Try inmultire
	if tryOperations(nums, pos+1, current*nums[pos], target) {
		return true
	}

	return false
}

func SecondPuzzle() int {
	input := parseInput("day7/input.txt")
	result := 0
	for _, pair := range input {
		if canReachTarget2(pair.nums, pair.target) {
			result += pair.target
		}
	}
	return result
}

func canReachTarget2(nums []int, target int) bool {
	return tryOperations2(nums, 1, nums[0], target)
}

func tryOperations2(nums []int, pos int, current int, target int) bool {
	if pos == len(nums) {
		return current == target
	}

	// Try adunare
	if tryOperations2(nums, pos+1, current+nums[pos], target) {
		return true
	}

	// Try inmultire
	if tryOperations2(nums, pos+1, current*nums[pos], target) {
		return true
	}

	// Try concat
	concat, err := strconv.Atoi(strconv.Itoa(current) + strconv.Itoa(nums[pos]))
	if err != nil {
		log.Fatalln(err)
	}
	if tryOperations2(nums, pos+1, concat, target) {
		return true
	}

	return false
}
