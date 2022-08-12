package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	points := strings.Split(scanner.Text(), " ")

	max := 0
	current := 0

	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			p, _ := strconv.Atoi(points[j])
			current += p

			if current > max {
				max = current
			}
		}
	}

	fmt.Println(max)
}
