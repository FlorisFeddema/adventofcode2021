package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	result := 0

	buf, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	data := string(buf)
	data = strings.ReplaceAll(data, "\n", ",")

	list := strings.Split(data, ",")

	input := []int{}

	for _, e := range list {
		i, _ := strconv.Atoi(e)
		input = append(input, i)
	}

	for i := range input {
		if i == 0 {
			continue
		}

		if input[i] > input[i-1] {
			result++
		}
	}

	fmt.Printf("Result: %[1]d", result)
}
