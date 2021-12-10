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

	rows := strings.Count(data, ",") + 3
	columns := len(strings.Split(data, ",")[0]) + 2

	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]int, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			grid[i][j] = 9
		}
	}

	for i, line := range strings.Split(data, ",") {
		for j, f := range line {
			grid[i+1][j+1], _ = strconv.Atoi(string(f))
		}
	}

	for i := 1; i < rows-1; i++ {
		for j := 1; j < columns-1; j++ {
			value := grid[i][j]
			if value < grid[i-1][j] &&
				value < grid[i+1][j] &&
				value < grid[i][j-1] &&
				value < grid[i][j+1] {
				result += value + 1
			}
		}
	}

	fmt.Printf("Result: %[1]d\n", result)
}
