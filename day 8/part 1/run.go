package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type display struct {
	numbers []string
	code    []string
}

func main() {
	result := 0

	buf, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(buf))
	data = strings.ReplaceAll(data, "\n", ",")

	list := strings.Split(data, ",")

	displays := []display{}

	for _, e := range list {
		values := strings.Split(e, " | ")
		display := display{
			numbers: strings.Split(values[0], " "),
			code:    strings.Split(values[1], " "),
		}

		displays = append(displays, display)
	}

	for _, e := range displays {
		result += checkDigits(e)
	}

	fmt.Printf("Result: %[1]d\n", result)
}

func checkDigits(display display) int {
	total := 0
	for _, e := range display.code {
		if checkOne(e) || checkFour(e) || checkSeven(e) || checkEight(e) {
			total++
			continue
		}
	}
	return total
}

func checkOne(input string) bool {
	if len(input) == 2 {
		return true
	}
	return false
}

func checkFour(input string) bool {
	if len(input) == 4 {
		return true
	}
	return false
}

func checkSeven(input string) bool {
	if len(input) == 3 {
		return true
	}
	return false
}

func checkEight(input string) bool {
	if len(input) == 7 {
		return true
	}
	return false
}
