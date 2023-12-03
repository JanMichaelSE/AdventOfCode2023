package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var schematic []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		schematic = append(schematic, line)
	}

	totalGearRatio := 0
	for y, row := range schematic {
		for x, ch := range row {
			if ch == '*' {
				partNumbers := getAdjacentPartNumbers(schematic, y, x)
				if len(partNumbers) == 2 {
					totalGearRatio += partNumbers[0] * partNumbers[1]
				}
			}
		}
	}

	fmt.Println("Total Gear Ratio is:", totalGearRatio)
}

func getAdjacentPartNumbers(schematic []string, y, x int) []int {
	directions := []struct{ dx, dy int }{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	}

	var partNumbers []int
	for _, dir := range directions {
		newX, newY := x+dir.dx, y+dir.dy
		if newY >= 0 && newY < len(schematic) && newX >= 0 && newX < len(schematic[newY]) {
			if number, ok := getNumberAt(schematic, newY, newX); ok {
				partNumbers = append(partNumbers, number)
			}
		}
	}

	return partNumbers
}

func getNumberAt(schematic []string, y, x int) (int, bool) {
	if !unicode.IsDigit(rune(schematic[y][x])) {
		return 0, false
	}

	start, end := x, x
	for start > 0 && unicode.IsDigit(rune(schematic[y][start-1])) {
		start--
	}
	for end < len(schematic[y])-1 && unicode.IsDigit(rune(schematic[y][end+1])) {
		end++
	}

	number, _ := strconv.Atoi(schematic[y][start : end+1])
	return number, true
}
