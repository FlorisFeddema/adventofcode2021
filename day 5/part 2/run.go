package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type line struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func main() {
	result := 0

	buf, err := ioutil.ReadFile("data2.txt")
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(buf))
	data = strings.ReplaceAll(data, "\n", "_")

	lines := []line{}

	dataLines := strings.Split(data, "_")
	for _, e := range dataLines {

		coordinates := strings.Split(e, " -> ")

		x := strings.Split(coordinates[0], ",")
		y := strings.Split(coordinates[1], ",")

		x1, _ := strconv.Atoi(x[0])
		x2, _ := strconv.Atoi(x[1])
		y1, _ := strconv.Atoi(y[0])
		y2, _ := strconv.Atoi(y[1])

		line := line{x1, x2, y1, y2}

		lines = append(lines, line)
	}

	grid := [10][10]int{}

	for _, e := range lines {
		//vertical lines
		if e.x2 == e.y2 {
			//get highest
			high := 0
			low := 0
			if e.x1 > e.y1 {
				high = e.x1
				low = e.y1
			} else {
				high = e.y1
				low = e.x1
			}

			for i := 0; i <= (high - low); i++ {
				grid[low+i][e.x2]++
			}

		}
		//horizonal lines
		if e.x1 == e.y1 {
			//get highest
			high := 0
			low := 0
			if e.x2 > e.y2 {
				high = e.x2
				low = e.y2
			} else {
				high = e.y2
				low = e.x2
			}

			for i := 0; i <= (high - low); i++ {
				grid[e.x1][low+i]++
			}
		}
		//diagonal lines
		if e.x1 != e.y1 && e.x2 != e.y2 {
			//get highest
			start1 := 0
			end1 := 0
			start2 := 0
			end2 := 0
			if e.x2 > e.y2 {
				start1 = e.x2
				end1 = e.y2

				start2 = e.x1
				end2 = e.y1
			} else {
				start1 = e.y2
				end1 = e.x2

				start2 = e.y1
				end2 = e.x1
			}

			fmt.Print("Start ")
			fmt.Print(start1)
			fmt.Print(", ")
			fmt.Println(end1)
			fmt.Print("End ")
			fmt.Print(start2)
			fmt.Print(", ")
			fmt.Println(end2)

			negative := false

			if start2 > end2 {
				negative = true
			}

			if start1 > end1 {
				for i := 0; i <= (start1 - end1); i++ {
					if negative {
						grid[start1-i][start2-i]++
					} else {
						grid[start1-i][start2+i]++
					}
				}
			} else {
				for i := 0; i <= (start1 - end1); i++ {
					if negative {
						grid[start1+i][start2-i]++
					} else {
						grid[start1+i][start2+i]++
					}
				}
			}

			end2++
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
