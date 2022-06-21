package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input [4]string
	fmt.Scan(&input[0], &input[1], &input[2], &input[3])

	w, _ := strconv.ParseFloat(input[0], 32)
	h, _ := strconv.ParseFloat(input[1], 32)
	x, _ := strconv.ParseFloat(input[2], 32)
	y, _ := strconv.ParseFloat(input[3], 32)

	minX := 0.0
	minY := 0.0

	if x*h < (w-x)*h {
		minX = x * h
	} else {
		minX = (w - x) * h
	}

	if w*y < w*(h-y) {
		minY = w * y
	} else {
		minY = w * (h - y)
	}

	if minX == minY {
		fmt.Println(minX, 1)
	} else {
		if minX < minY {
			fmt.Println(minY, 0)
		} else {
			fmt.Println(minX, 0)
		}
	}
}
