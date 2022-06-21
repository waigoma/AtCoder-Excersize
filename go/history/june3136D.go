package main

import (
	"fmt"
	"strings"
)

func main() {
	var input [1]string
	fmt.Scan(&input[0])

	split := strings.Split(input[0], "")

	length := len(split)

	children := make([]int, length)

	for i := 0; i < length; i++ {
		children[i] = 1
	}

	for i := 0; i <= 1000; i++ {
		for j := 0; j < length; j++ {
			if children[j] == 0 {
				continue
			}
			if split[j] == "R" {
				children[j]--
				children[j+1]++
			} else if split[j] == "L" {
				children[j]--
				children[j-1]++
			}
		}
	}

	for _, tmp := range children {
		fmt.Print(tmp, " ")
	}
}
