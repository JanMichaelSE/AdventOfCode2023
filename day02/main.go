package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	reds   = 12
	greens = 13
	blues  = 14
)

func main() {

	// TODO: Future improved passing expected values

	scanner := bufio.NewScanner(os.Stdin)
	counter := 1
	cubesColorMap := map[int][]string{}
	for scanner.Scan() {
		text := scanner.Text()

		valuesStr := strings.Split(text, ":")
		trimmedValuesStr := strings.Trim(valuesStr[1], " ")
		cubesColorMap[counter] = strings.Split(strings.ReplaceAll(trimmedValuesStr, ";", ","), ",")

		counter++
	}

	cubeMaxPerGame := map[int][3]int{}
	for key, values := range cubesColorMap {
		cubeMaxPerGame[key] = [3]int{0, 0, 0}
		for _, cube := range values {
			trimmedColor := strings.TrimLeft(cube, " ")
			colorSplit := strings.Split(trimmedColor, " ")
			amount, _ := strconv.Atoi(colorSplit[0])
			color := colorSplit[1]

			temp := cubeMaxPerGame[key]
			if color == "green" && cubeMaxPerGame[key][0] < amount {
				temp[0] = amount
				cubeMaxPerGame[key] = temp
			} else if color == "blue" && cubeMaxPerGame[key][1] < amount {
				temp[1] = amount
				cubeMaxPerGame[key] = temp
			} else if color == "red" && cubeMaxPerGame[key][2] < amount {
				temp[2] = amount
				cubeMaxPerGame[key] = temp
			}
		}
	}

	totalPossibleGames := 0
	totalSum := 0
	for i, values := range cubeMaxPerGame {

		possible := true
		totalPower := 1
		for j, cube := range values {
			switch {
			case j == 0 && cube > greens:
				possible = false
			case j == 1 && cube > blues:
				possible = false
			case j == 2 && cube > reds:
				possible = false
			}

			totalPower *= cube
		}

		totalSum += totalPower

		if possible {
			totalPossibleGames += i
		}
	}

	fmt.Println("Total Possible Games:", totalPossibleGames)
	fmt.Println("Total Power Sum:", totalSum)
}
