package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input [3]string
	fmt.Scan(&input[0], &input[1], &input[2])

	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])

	if (a >= 0 && b >= 0) || c%2 == 0 {
		a = int(math.Abs(float64(a)))
		b = int(math.Abs(float64(b)))
	}

	if a > b {
		fmt.Println(">")
	} else if a < b {
		fmt.Println("<")
	} else {
		fmt.Println("=")
	}
}
