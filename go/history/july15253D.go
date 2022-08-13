package main

import (
	"fmt"
	"strconv"
)

func strToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func main() {
	var inputs [3]string

	fmt.Scan(&inputs[0], &inputs[1], &inputs[2])

	n := strToInt(inputs[0])
	a := strToInt(inputs[1])
	b := strToInt(inputs[2])

	res := 0

	for i := 0; i < n; i++ {
		if i%a != 0 && i%b != 0 {
			res += i
		}
	}

	fmt.Println(res)
}
