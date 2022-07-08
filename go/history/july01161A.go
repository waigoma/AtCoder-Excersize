package main

import (
	"fmt"
)

func main() {
	var str [3]string
	fmt.Scan(&str[0], &str[1], &str[2])

	str[0], str[1] = str[1], str[0]
	str[0], str[2] = str[2], str[0]

	fmt.Println(str[0], str[1], str[2])
}
