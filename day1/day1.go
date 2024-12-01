package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func FirstPuzzle() int {
	file, err := os.Open("day1/input1.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left_nums []int
	var right_nums []int

	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Fields(line)

		left, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		right, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		left_nums = append(left_nums, left)
		right_nums = append(right_nums, right)
	}

	sort.Ints(left_nums)
	sort.Ints(right_nums)

	result := 0

	for i := range len(left_nums) {
		result += int(math.Abs(float64(left_nums[i] - right_nums[i])))
	}

	return result
}

func SecondPuzzle() int {
	file, err := os.Open("day1/input2.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left_nums []int
	var right_nums []int

	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Fields(line)

		left, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		right, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		left_nums = append(left_nums, left)
		right_nums = append(right_nums, right)
	}

	result := 0

	for _, elem := range left_nums {
		result += (elem * count(right_nums, elem))
	}

	return result
}

func count(nums []int, element int) int {
	count := 0
	for _, val := range nums {
		if val == element {
			count++
		}
	}
	return count
}
