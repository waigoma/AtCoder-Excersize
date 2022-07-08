package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str [1]string
	fmt.Scan(&str[0])

	num, _ := strconv.Atoi(str[0])

	divide := num / 500
	num -= 500 * divide

	div := num / 5

	fmt.Println(1000*divide + 5*div)
}
