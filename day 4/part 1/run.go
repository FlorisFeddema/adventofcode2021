package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bingo struct {
	data [][]string
}

func main() {
	result := 0

	buf, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(buf))
	data = strings.ReplaceAll(data, "\n ", "-")
	data = strings.ReplaceAll(data, "\n", "-")

	rolls := strings.Split(strings.Split(data, "-")[0], ",")

	data = strings.ReplaceAll(data, "  ", " ")
	data = strings.ReplaceAll(data, " ", ",")
	data = strings.ReplaceAll(data, "--", "-")

	cardData := strings.Split(data, "-")[1:]

	cards := []bingo{}

	for i := 0; i < len(cardData); i += 5 {
		row1 := strings.Split(cardData[i], ",")
		row2 := strings.Split(cardData[i+1], ",")
		row3 := strings.Split(cardData[i+2], ",")
		row4 := strings.Split(cardData[i+3], ",")
		row5 := strings.Split(cardData[i+4], ",")

		card := bingo{
			data: [][]string{row1, row2, row3, row4, row5},
		}
		cards = append(cards, card)
	}

	winner, currentRolls := playgame(rolls, cards)

	for _, e := range winner.data {
		for _, f := range e {
			if !contains(currentRolls, f) {
				value, _ := strconv.Atoi(f)
				result += value
			}
		}
	}

	lastRoll, _ := strconv.Atoi(currentRolls[len(currentRolls)-1])
	result = result * lastRoll

	fmt.Printf("Result: %[1]d\n", result)
}

func contains(s []string, input string) bool {
	for _, e := range s {
		if e == input {
			return true
		}
	}
	return false
}

func playgame(rolls []string, cards []bingo) (bingo, []string) {
	currentRolls := []string{}

	for _, e := range rolls {
		currentRolls = append(currentRolls, e)
		for _, f := range cards {
			if checkHorizontal(f, currentRolls) {
				fmt.Println("Horizontal winner")
				return f, currentRolls
			}
			if checkVertical(f, currentRolls) {
				fmt.Println("Vertical winner")
				return f, currentRolls
			}
		}
	}
	return *new(bingo), currentRolls
}

func checkHorizontal(card bingo, rolls []string) bool {
	for _, line := range card.data {
		match := 0
		for _, number := range line {
			if contains(rolls, number) {
				match++
				if match == 5 {
					return true
				}
			}
		}
	}
	return false
}

func checkVertical(card bingo, rolls []string) bool {
	for i := 0; i < len(card.data[0]); i++ {
		match := 0
		for j := 0; j < len(card.data); j++ {
			if contains(rolls, card.data[j][i]) {
				match++
				if match == 5 {
					return true
				}
			}
		}
	}
	return false
}
