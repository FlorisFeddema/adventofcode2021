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

	data := strings.TrimSpace(string(buf))
	data = strings.ReplaceAll(data, "\n", ",")

	list := strings.Split(data, ",")

	postition := 0
	aim := 0
	depth := 0

	for _, e := range list {
		value, _ := strconv.Atoi(strings.Split(e, " ")[1])

		if strings.Contains(e, "forward") {
			postition += value

			depth += value * aim
		}
		if strings.Contains(e, "down") {
			aim += value
		}
		if strings.Contains(e, "up") {
			aim -= value
		}
	}

	result = depth * postition

	fmt.Printf("Result: %[1]d\n", result)
}
