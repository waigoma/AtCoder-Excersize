package main

import (
	"fmt"
	"strconv"
	"strings"
)

var a int
var b int
var x int

func calc(n int) int {
	dn := strings.Split(string(rune(n)), "")
	return a*n + b*len(dn)
}

func find(start int, end int) bool {
	for i := start; i < end; i++ {
		if calc(i) > x {
			fmt.Println(i - 1)
			return true
		}
	}
	return false
}

func half(down int, up int) {

}

func main() {
	var input [3]string
	fmt.Scan(&input[0], &input[1], &input[2])

	a, _ = strconv.Atoi(input[0])
	b, _ = strconv.Atoi(input[1])
	x, _ = strconv.Atoi(input[2])

	if x < calc(1) {
		fmt.Println("0")
		return
	}

	upHalfX := int(float64(x) * 0.75)
	halfX := x / 2
	downHalfX := x / 4

	if calc(halfX) > x {
		if calc(downHalfX) > x {
			for i := 0; i < downHalfX; i++ {
				if calc(i) > x {
					fmt.Println(i - 1)
					return
				}
			}
		} else {
			for i := downHalfX; i < halfX; i++ {
				if calc(i) > x {
					fmt.Println(i - 1)
					return
				}
			}
		}
	} else {
		if calc(upHalfX) > x {
			for i := 0; i < upHalfX; i++ {
				if calc(i) > x {
					fmt.Println(i - 1)
					return
				}
			}
		} else {
			for i := upHalfX; i < x; i++ {
				if calc(i) > x {
					fmt.Println(i - 1)
					return
				}
			}
		}
	}

}
