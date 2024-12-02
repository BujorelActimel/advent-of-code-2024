package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(nums []int) bool {
	increasing := nums[0] < nums[1]

	for i := range len(nums) - 1 {
		if increasing && (nums[i+1]-nums[i] > 3 || nums[i+1]-nums[i] < 1) {
			return false
		}
		if !increasing && (nums[i]-nums[i+1] > 3 || nums[i]-nums[i+1] < 1) {
			return false
		}
	}
	return true
}

func stringToIntSlice(line string) []int {
	var nums []int

	for _, num := range strings.Fields(line) {
		converted_num, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		nums = append(nums, converted_num)
	}

	return nums
}

func removeElement(slice []int, index int) []int {
	var resultSlice []int
	resultSlice = append(resultSlice, slice[:index]...)
	resultSlice = append(resultSlice, slice[index+1:]...)
	return resultSlice
}

func FirstPuzzle() int {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		nums := stringToIntSlice(line)

		if isSafe(nums) {
			result++
		}
	}

	return result
}

func SecondPuzzle() int {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		nums := stringToIntSlice(line)

		if isSafe(nums) {
			result++
			continue
		}

		for i := range len(nums) {
			if isSafe(removeElement(nums, i)) {
				result++
				break
			}
		}
	}

	return result
}
