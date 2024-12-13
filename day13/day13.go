package day13

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Button struct {
	x, y int64
}

type Prize struct {
	x, y int64
}

type Triplet struct {
	A, B  Button
	prize Prize
}

func parseInput(filename string) []Triplet {
	var triplets []Triplet

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for {
		if !scanner.Scan() {
			break
		}

		buttonALine := scanner.Text()
		if !scanner.Scan() {
			break
		}
		buttonBLine := scanner.Text()
		if !scanner.Scan() {
			break
		}
		prizeLine := scanner.Text()
		if !scanner.Scan() {
			buttonA := parseButton(buttonALine)
			buttonB := parseButton(buttonBLine)
			prize := parsePrize(prizeLine)
			triplets = append(triplets, Triplet{A: buttonA, B: buttonB, prize: prize})
			break
		}
		_ = scanner.Text()

		buttonA := parseButton(buttonALine)
		buttonB := parseButton(buttonBLine)
		prize := parsePrize(prizeLine)

		triplets = append(triplets, Triplet{A: buttonA, B: buttonB, prize: prize})
	}

	return triplets
}

func parseButton(line string) Button {
	parts := strings.Split(line, ", ")
	xPart := strings.Split(parts[0], "+")
	yPart := strings.Split(parts[1], "+")
	x, _ := strconv.ParseInt(xPart[1], 10, 64)
	y, _ := strconv.ParseInt(yPart[1], 10, 64)
	return Button{x: x, y: y}
}

func parsePrize(line string) Prize {
	parts := strings.Split(line, ", ")
	xPart := strings.Split(parts[0], "=")
	yPart := strings.Split(parts[1], "=")
	x, _ := strconv.ParseInt(xPart[1], 10, 64)
	y, _ := strconv.ParseInt(yPart[1], 10, 64)
	return Prize{x: x, y: y}
}

func FirstPuzzle() int64 {
	triplets := parseInput("day13/input.txt")

	var result int64 = 0

	for _, triplet := range triplets {
		cost, canWin := calculateMinCost(triplet)
		if canWin {
			result += cost
		}
	}

	return result
}

func SecondPuzzle() int64 {
	triplets := parseInput("day13/input.txt")

	var offset int64 = 10000000000000
	for i := range triplets {
		triplets[i].prize.x += offset
		triplets[i].prize.y += offset
	}

	var result int64 = 0

	for _, triplet := range triplets {
		cost, canWin := calculateMinCost(triplet)
		if canWin {
			result += cost
		}
	}

	return result
}

func calculateMinCost(triplet Triplet) (int64, bool) {
	if !checkIfPossible(triplet) {
		return 0, false
	}

	gcd1, x1, y1 := extendedGCD(triplet.A.x, triplet.B.x)
	baseX := x1 * (triplet.prize.x / gcd1)
	baseY := y1 * (triplet.prize.x / gcd1)

	a := triplet.A.y*(triplet.B.x/gcd1) - triplet.B.y*(triplet.A.x/gcd1)
	b := triplet.prize.y - triplet.A.y*baseX - triplet.B.y*baseY

	if a == 0 {
		if b == 0 {
			return findMinCost(baseX, baseY, triplet.B.x/gcd1, -triplet.A.x/gcd1), true
		}
		return 0, false
	}

	if b%a != 0 {
		return 0, false
	}

	k1 := b / a
	A := baseX + (triplet.B.x/gcd1)*k1
	B := baseY - (triplet.A.x/gcd1)*k1

	if A < 0 || B < 0 {
		return 0, false
	}

	return 3*A + B, true
}

func checkIfPossible(triplet Triplet) bool {
	return triplet.prize.x%gcd(triplet.A.x, triplet.B.x) == 0 &&
		triplet.prize.y%gcd(triplet.A.y, triplet.B.y) == 0
}

func findMinCost(baseX, baseY, dx, dy int64) int64 {
	var kMin, kMax int64
	if dx > 0 {
		kMin = (-baseX + dx - 1) / dx
	} else if dx < 0 {
		kMax = -baseX / dx
	}
	if dy > 0 {
		kMin = max64(kMin, (-baseY+dy-1)/dy)
	} else if dy < 0 {
		kMax = min64(kMax, -baseY/dy)
	}

	var minCost int64 = math.MaxInt64
	for k := kMin; k <= kMax; k++ {
		x := baseX + dx*k
		y := baseY + dy*k
		if x >= 0 && y >= 0 {
			cost := 3*x + y
			minCost = min64(minCost, cost)
		}
	}
	return minCost
}

func extendedGCD(a, b int64) (int64, int64, int64) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return abs64(a)
}

func abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
