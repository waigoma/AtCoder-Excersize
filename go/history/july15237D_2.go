package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input [2]string
	fmt.Scan(&input[0])
	fmt.Scan(&input[1])

	n, _ := strconv.Atoi(input[0])
	s := input[1]

	for i {

	}

	res := make([]int, n+1)
	fmt.Println(len(res))
	mv := 0

	for i, v := range s {
		if string(v) == "R" {
			mv++
		}
		res = insert(res, mv, i+1)
		fmt.Println(res)
	}

	for _, v := range res {
		fmt.Print(v)
		fmt.Print(" ")
	}
}
