package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	gammaValues := [12]int{}

	for _, e := range list {
		for j, f := range e {
			if f == 48 {
				gammaValues[j]--
			} else {
				gammaValues[j]++
			}
		}
	}

	for i, e := range gammaValues {
		if e > 0 {
			gammaValues[i] = 1
		} else {
			gammaValues[i] = 0
		}
	}

	gamma := 0
	epsilon := 0

	for i, e := range gammaValues {
		power := len(gammaValues) - i - 1

		if e == 1 {
			gamma += int(math.Pow(float64(2), float64(power)))
		} else {
			epsilon += int(math.Pow(float64(2), float64(power)))
		}
	}

	result = gamma * epsilon

	fmt.Printf("Result: %[1]d", result)
}
