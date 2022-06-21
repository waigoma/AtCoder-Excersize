package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input [1]string
	fmt.Scan(&input[0])

	n, _ := strconv.Atoi(input[0])

	count := 0

	for i := 1; i <= n; i++ {
		split := strings.Split(strconv.Itoa(i), "")

		if len(split)%2 != 0 {
			count++
		}
	}

	fmt.Println(count)
}
