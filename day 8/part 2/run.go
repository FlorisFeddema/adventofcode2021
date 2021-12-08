package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
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

		numbers := strings.Split(values[0], " ")
		code := strings.Split(values[1], " ")

		for i := range numbers {
			numbers[i] = SortString(numbers[i])
		}

		for i := range code {
			code[i] = SortString(code[i])
		}

		display := display{
			numbers: numbers,
			code:    code,
		}

		displays = append(displays, display)
	}

	for _, e := range displays {
		result += calculateCode(e)
	}

	fmt.Printf("Result: %[1]d\n", result)
}

func calculateCode(display display) int {
	mappings := make(map[string]int)

	one, four, six := "", "", ""

	for _, e := range display.numbers {
		if len(e) == 2 {
			mappings[e] = 1
			one = e
			continue
		}
		if len(e) == 3 {
			mappings[e] = 7
			continue
		}
		if len(e) == 4 {
			mappings[e] = 4
			four = e
			continue
		}
		if len(e) == 7 {
			mappings[e] = 8
			continue
		}
	}

	for _, e := range display.numbers {
		if len(e) == 6 {
			if isSix(one, e) {
				mappings[e] = 6
				six = e
			} else if isZero(four, e) {
				mappings[e] = 0
			} else {
				mappings[e] = 9
			}
		}
	}

	for _, e := range display.numbers {
		if len(e) == 5 {
			if isFiveOrTwo(one, e) {
				if isFive(six, e) {
					mappings[e] = 2
				} else {
					mappings[e] = 5
				}
			} else {
				mappings[e] = 3
			}
		}
	}

	output := ""

	for _, e := range display.code {
		output += strconv.Itoa(mappings[e])
	}

	result, _ := strconv.Atoi(output)
	return result
}

func isFive(six string, e string) bool {
	for _, f := range e {
		// if not contains from 6
		if !strings.Contains(six, string(f)) {
			return true
		}
	}
	return false
}

func isFiveOrTwo(one, e string) bool {
	for _, f := range one {
		// if not contains from 1
		if !strings.Contains(e, string(f)) {
			return true
		}
	}
	return false
}

func isSix(one string, e string) bool {
	for _, f := range one {
		// if not contains from 1
		if !strings.Contains(e, string(f)) {
			return true
		}
	}
	return false
}

func isZero(four string, e string) bool {
	for _, f := range four {
		// if not contains from 4
		if !strings.Contains(e, string(f)) {
			return true
		}
	}
	return false
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
