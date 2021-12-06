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

	fish := [9]int{}

	for _, e := range list {
		value, _ := strconv.Atoi(e)
		fish[value]++
	}

	fmt.Println(fish)

	for i := 0; i < 80; i++ {
		nFish := [9]int{}
		for i, e := range fish {
			if i == 0 {
				nFish[6] += e
				nFish[8] += e
			} else {
				nFish[i-1] += e
			}
		}
		fish = nFish
		fmt.Println(fish)
	}

	for _, e := range fish {
		result += e
	}

	fmt.Printf("Result: %[1]d\n", result)
}
