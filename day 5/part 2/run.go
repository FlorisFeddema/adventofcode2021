package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type line struct {
	startX int
	startY int
	endX   int
	endY   int
}

func main() {
	result := 0

	buf, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(buf))
	data = strings.ReplaceAll(data, "\n", "_")

	lines := []line{}

	dataLines := strings.Split(data, "_")
	for _, e := range dataLines {

		coordinates := strings.Split(e, " -> ")

		start := strings.Split(coordinates[0], ",")
		end := strings.Split(coordinates[1], ",")

		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])
		endX, _ := strconv.Atoi(end[0])
		endY, _ := strconv.Atoi(end[1])

		line := line{startX, startY, endX, endY}

		lines = append(lines, line)
	}

	grid := [1000][1000]int{}

	for _, e := range lines {
		//horizontal
		if e.startY == e.endY {

			start, end := 0, 0

			if e.startX > e.endX {
				start = e.endX
				end = e.startX
			} else {
				end = e.endX
				start = e.startX
			}

			for i := start; i <= end; i++ {
				grid[e.startY][i]++
			}
		}
		//vertical
		if e.startX == e.endX {

			start, end := 0, 0

			if e.startY > e.endY {
				start = e.endY
				end = e.startY
			} else {
				end = e.endY
				start = e.startY
			}

			for i := start; i <= end; i++ {
				grid[i][e.startX]++
			}
		}
		//diagonal
		if e.endX != e.startX && e.endY != e.startY {

			lenght := e.startX - e.endX
			if lenght < 0 {
				lenght *= -1
			}

			modX, modY := 1, 1

			if e.startX > e.endX {
				modX = -1
			}

			if e.startY > e.endY {
				modY = -1
			}

			for i := 0; i <= lenght; i++ {
				grid[e.startY+(i*modY)][e.startX+(i*modX)]++
			}
		}
	}

	for _, e := range grid {
		fmt.Println(e)
		for _, f := range e {
			if f > 1 {
				result++
			}
		}
	}

	fmt.Printf("Result: %[1]d\n", result)
}
