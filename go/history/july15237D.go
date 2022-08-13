package main

import (
	"fmt"
)

func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func main() {
	var input [2]string
	fmt.Scan(&input[0])
	fmt.Scan(&input[1])

	//n, _ := strconv.Atoi(input[0])
	s := input[1]

	res := make([]int, 1)
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
