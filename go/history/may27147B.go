package main

import (
	"fmt"
	"strings"
)

func main() {
	var input [1]string
	fmt.Scan(&input[0])

	slice := strings.Split(input[0], "")
	inLength := len(slice)

	front := make([]string, inLength)
	back := make([]string, inLength)

	for i := 0; i < inLength/2; i++ {
		front[i] = slice[i]
		back[i] = slice[inLength-1-i]
	}

	count := 0
	for i := 0; i < len(front); i++ {
		if front[i] != back[i] {
			count++
		}
	}

	fmt.Println(count)
}
