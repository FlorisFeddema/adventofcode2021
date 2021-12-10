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

	fmt.Println("-----------------------")

	buf, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(buf))
	data = strings.ReplaceAll(data, "\n", ",")

	rows := strings.Count(data, ",") + 3
	columns := len(strings.Split(data, ",")[0]) + 2

	basins := []basin{}

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

				basin := basin{
					row:    i,
					column: j,
				}

				basins = append(basins, basin)
			}
		}
	}

	values := []int{}

	for _, basin := range basins {
		points := make([][]bool, rows)
		for i := 0; i < rows; i++ {
			points[i] = make([]bool, columns)
		}

		values = append(values, calculate(basin.row, basin.column, grid, points))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	result = values[0] * values[1] * values[2]

	fmt.Printf("Result: %[1]d\n", result)
}

func calculate(i int, j int, grid [][]int, points [][]bool) int {

	if points[i][j] {
		return 0
	} else {
		points[i][j] = true
	}

	value := 1

	if grid[i-1][j] != 9 {
		value += calculate(i-1, j, grid, points)
	}
	if grid[i+1][j] != 9 {
		value += calculate(i+1, j, grid, points)
	}
	if grid[i][j-1] != 9 {
		value += calculate(i, j-1, grid, points)
	}
	if grid[i][j+1] != 9 {
		value += calculate(i, j+1, grid, points)
	}

	return value
}

type basin struct {
	row    int
	column int
}
