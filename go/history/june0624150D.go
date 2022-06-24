package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...[]int) int {
	result := a * b / GCD(a, b)

	if integers == nil {
		return result
	}

	for i := 0; i < len(integers[0]); i++ {
		result = LCM(result, integers[0][i]/2)
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N_M := scanner.Text()
	nmSplit := strings.Split(N_M, " ")

	n, _ := strconv.Atoi(nmSplit[0])
	m, _ := strconv.Atoi(nmSplit[1])

	scanner.Scan()
	a_n := scanner.Text()
	anSplit := strings.Split(a_n, " ")

	var an []int
	for _, v := range anSplit {
		num, _ := strconv.Atoi(v)
		an = append(an, num)
	}

	anSlice := an[2:n]
	lcm := LCM(an[0]/2, an[1]/2, anSlice)

	mult := 1
	count := 0

	if lcm > m {
		fmt.Println(0)
		os.Exit(0)
	}

	lcmTmp := lcm
	lt := lcmTmp

	for {
		if lt <= m {
			count++
			mult += 2
			lt = lcmTmp * mult
		} else {
			break
		}
	}

	fmt.Println(count)
}
