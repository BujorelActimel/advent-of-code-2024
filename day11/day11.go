package day11

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) map[int]int {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	input := string(data)

	stones := make(map[int]int)
	for _, s := range strings.Fields(input) {
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}
		stones[val]++
	}
	return stones
}

func FirstPuzzle() int {
	stones := parseInput("day11/input.txt")

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}

	return sum
}

func SecondPuzzle() int {
	stones := parseInput("day11/input.txt")

	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}

	return sum
}

func numDigits(i int) int {
	return int(math.Log10(float64(i))) + 1
}

func powTen(pow int) int {
	n := 1
	for i := 0; i < pow; i++ {
		n *= 10
	}
	return n
}

func blink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)

	add := func(key, incr int) {
		newStones[key] += incr
	}

	for stone, count := range stones {
		if stone == 0 {
			add(1, count)
		} else if digits := numDigits(stone); digits%2 == 0 {
			filter := powTen(digits / 2)
			left, right := stone/filter, stone%filter
			add(left, count)
			add(right, count)
		} else {
			add(stone*2024, count)
		}
	}
	return newStones
}
