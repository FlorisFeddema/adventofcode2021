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

	oxigen := strings.Split(data, ",")

	for i := range oxigen[0] {
		count := 0
		for _, e := range oxigen {
			value := e[i]
			if value == 48 {
				count--
			} else {
				count++
			}
		}

		zeroMode := false
		if count < 0 {
			zeroMode = true
		}

		tmp := []string{}
		for _, e := range oxigen {
			value := e[i]
			if zeroMode {
				if value == 48 { //non zero value
					tmp = append(tmp, e)
				}
			} else {
				if value != 48 { // zero value
					tmp = append(tmp, e)
				}
			}
		}

		oxigen = tmp
		if len(oxigen) == 1 {
			break
		}
	}

	co2 := strings.Split(data, ",")

	for i := range co2[0] {
		count := 0
		for _, e := range co2 {
			value := e[i]
			if value == 48 {
				count--
			} else {
				count++
			}
		}

		zeroMode := false
		if count < 0 {
			zeroMode = true
		}

		tmp := []string{}
		for _, e := range co2 {
			value := e[i]
			if zeroMode {
				if value != 48 { //non zero value
					tmp = append(tmp, e)
				}
			} else {
				if value == 48 { // zero value
					tmp = append(tmp, e)
				}
			}
		}

		co2 = tmp
		if len(co2) == 1 {
			break
		}
	}

	co2Value, _ := strconv.ParseInt(co2[0], 2, 64)
	oxigenValue, _ := strconv.ParseInt(oxigen[0], 2, 64)

	result = int(co2Value) * int(oxigenValue)

	fmt.Printf("Result: %[1]d\n", result)
}
