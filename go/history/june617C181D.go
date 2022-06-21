package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input [1]string
	fmt.Scan(&input[0])

	if len(input[0]) == 1 {
		n, _ := strconv.Atoi(input[0])
		if n%8 == 0 {
			fmt.Println("Yes")
			os.Exit(0)
		}
	}

	if len(input[0]) == 2 {
		split := strings.Split(input[0], "")
		n1str := split[0] + split[1]
		n2str := split[1] + split[0]
		n1, _ := strconv.Atoi(n1str)
		n2, _ := strconv.Atoi(n2str)

		if n1%8 == 0 || n2%8 == 0 {
			fmt.Println("Yes")
			os.Exit(0)
		}
	}

	var list []string
	for i := 100; i < 1000; i++ {
		if i%8 == 0 {
			str := strconv.Itoa(i)
			if strings.Contains(str, "0") {
				continue
			}
			list = append(list, strconv.Itoa(i))
		}
	}

	//var number []int
	number := make([]int, 9)
	for _, str := range input[0] {
		n, _ := strconv.Atoi(string(str))
		number[n-1]++
	}

	for _, li := range list {
		isFind := true
		tmp := make([]int, len(number))
		copy(tmp, number)

		for _, str := range li {
			n, _ := strconv.Atoi(string(str))
			tmp[n-1]--
		}

		for _, num := range tmp {
			if num < 0 {
				isFind = false
			}
		}

		if isFind {
			fmt.Println("Yes")
			os.Exit(0)
		}
	}

	fmt.Println("No")
}
