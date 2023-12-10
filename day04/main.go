package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	winningCards := []map[string]bool{}
	playerCards := [][]string{}
	for scanner.Scan() {
		text := scanner.Text()

		winningNumbers, playerNumbers := getNumbersFromText(text)
		winningCards = populateWinnerCards(winningNumbers, winningCards)
		playerCards = append(playerCards, playerNumbers)
	}

	wonTotal := 0
	for i, cards := range playerCards {
		roundTotal := 0
		for _, card := range cards {
			if winningCards[i][card] && roundTotal == 0 {
				roundTotal = 1
			} else if winningCards[i][card] {
				roundTotal *= 2
			}
		}
		wonTotal += roundTotal
	}

	println("Part 1:", wonTotal)
}

func getNumbersFromText(text string) ([]string, []string) {
	card := strings.Split(text, ":")[1]
	cardNumbers := strings.Split(strings.TrimSpace(card), "|")
	leftNumbers := strings.Split(strings.TrimSpace(cardNumbers[0]), " ")
	rightNumbers := strings.Split(strings.TrimSpace(cardNumbers[1]), " ")

	winnerNumbers := []string{}
	for _, card := range leftNumbers {
		if card != "" {
			winnerNumbers = append(winnerNumbers, card)
		}
	}

	playerNumbers := []string{}
	for _, card := range rightNumbers {
		if card != "" {
			playerNumbers = append(playerNumbers, card)
		}
	}

	return winnerNumbers, playerNumbers
}

func populateWinnerCards(cards []string, winningCards []map[string]bool) []map[string]bool {
	winners := map[string]bool{}
	for _, card := range cards {
		winners[card] = true
	}
	winningCards = append(winningCards, winners)
	return winningCards
}
