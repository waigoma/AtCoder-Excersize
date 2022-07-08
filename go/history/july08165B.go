package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input [1]string
	fmt.Scan(&input[0])

	x, _ := strconv.Atoi(input[0])
	total := 100

	for i := 1; ; i++ {
		total += int(float64(total) * 0.01)
		if total >= x {
			fmt.Println(i)
			break
		}
	}
}
