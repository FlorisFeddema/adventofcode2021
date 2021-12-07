package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := 0

	buf, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(buf))
	data = strings.ReplaceAll(data, "\n", ",")

	list := strings.Split(data, ",")

	crabs := []int{}

	for _, e := range list {
		value, _ := strconv.Atoi(e)
		crabs = append(crabs, value)
	}

	sort.Ints(crabs)

	mid := len(crabs) / 2

	median := (crabs[mid-1] + crabs[mid]) / 2

	if mid%2 == 1 {
		median = crabs[mid]
	}

	for _, e := range crabs {
		if e > median {
			result += e - median
		} else {
			result += median - e
		}
	}

	fmt.Printf("Result: %[1]d\n", result)
}
