package main

import (
	"fmt"
	"math"
	"strconv"
)

func floor(x, b int) int {
	return int(float64(x) / float64(b))
}

func main() {
	var inputs [3]string
	fmt.Scan(&inputs[0], &inputs[1], &inputs[2])

	a, _ := strconv.Atoi(inputs[0])
	b, _ := strconv.Atoi(inputs[1])
	n, _ := strconv.Atoi(inputs[2])

	max := math.MinInt32

	for i := n; i > 0; i-- {
		calc := floor(a*i, b) - (a * floor(i, b))
		if max < calc {
			max = calc
		}
	}

	fmt.Println(max)
}
