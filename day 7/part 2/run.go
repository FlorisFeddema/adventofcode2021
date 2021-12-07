package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	total := 0

	for _, e := range list {
		value, _ := strconv.Atoi(e)
		crabs = append(crabs, value)
		total += value
	}

	mean := int(math.Round(float64(total) / float64(len(crabs))))

	mean -= 1

	for _, e := range crabs {
		if e > mean {
			n := e - mean
			result += (n * (n + 1)) / 2
		} else {
			n := mean - e
			result += (n * (n + 1)) / 2
		}
	}

	fmt.Printf("Result: %[1]d\n", result)
}
