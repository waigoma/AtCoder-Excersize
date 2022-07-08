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
	nkSplit := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(nkSplit[0])
	k, _ := strconv.Atoi(nkSplit[1])

	people := make([]int, n)
	//snack := make([]string, k)

	for i := 1; i < k*2; i++ {
		scanner.Scan()
		if i%2 == 1 {
			continue
		}
		text := scanner.Text()
		split := strings.Split(text, " ")
		for _, sp := range split {
			s, _ := strconv.Atoi(sp)
			people[s-1] += 1
		}
	}

	count := 0
	for _, p := range people {
		if p == 0 {
			count += 1
		}
	}

	fmt.Println(count)
}
