package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	var split []string

	in := 0
	for 2 > in {
		in++
		scanner.Scan()
		if in == 1 {
			n, _ = strconv.Atoi(scanner.Text())
		} else if in == 2 {
			split = strings.Split(scanner.Text(), " ")
		}
	}

	ints := make([]int, 0, n)
	for _, str := range split {
		atoint, _ := strconv.Atoi(str)
		ints = append(ints, atoint)
	}

	sort.Ints(ints)

	for i := 0; i < n+1; i++ {
		hit := false
		for _, num := range ints {
			if i == num {
				hit = true
				break
			}
		}
		if !hit {
			fmt.Println(i)
			break
		}
	}
}
