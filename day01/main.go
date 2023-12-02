package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	inputFile = "input.txt"
	numberMap = map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	resultsStr := []string{}
	for scanner.Scan() {
		min := float64(math.MaxInt)
		max := float64(math.MinInt)
		minStr := ""
		maxStr := ""
		text := scanner.Text()
		for option := range numberMap {
			if newMin := math.Min(min, float64(strings.Index(text, option))); newMin < min && newMin != -1 {
				min = newMin
				minStr = option
			}

			if newMax := math.Max(max, float64(strings.LastIndex(text, option))); newMax > max && newMax != -1 {
				max = newMax
				maxStr = option
			}
		}

		resultsStr = append(resultsStr, numberMap[minStr])
		resultsStr[len(resultsStr)-1] += numberMap[maxStr]
	}

	total := 0
	for _, result := range resultsStr {
		value, _ := strconv.Atoi(result)
		total += value
	}

	fmt.Println("Final Result:", total)
}
