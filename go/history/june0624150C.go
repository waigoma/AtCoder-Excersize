package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getPermutations(elements []int) [][]int {
	permutations := [][]int{}
	if len(elements) == 1 {
		permutations = [][]int{elements}
		return permutations
	}
	for i := range elements {
		el := make([]int, len(elements))
		copy(el, elements)

		// or copy via append
		// el := append([]int(nil), elements...)

		for _, perm := range getPermutations(append(el[0:i], el[i+1:]...)) {
			permutations = append(permutations, append([]int{elements[i]}, perm...))
		}
	}
	return permutations
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nums := scanner.Text()
	num, _ := strconv.Atoi(nums)

	scanner.Scan()
	match1 := scanner.Text()
	match1Split := strings.Split(match1, " ")

	scanner.Scan()
	match2 := scanner.Text()
	match2Split := strings.Split(match2, " ")

	if match1 == match2 {
		fmt.Println(0)
		os.Exit(0)
	}

	var numList []int
	for i := 1; i <= num; i++ {
		numList = append(numList, i)
	}

	var listNum []int

	for i, v := range getPermutations(numList) {
		if fmt.Sprint(v) == fmt.Sprint(match1Split) || fmt.Sprint(v) == fmt.Sprint(match2Split) {
			listNum = append(listNum, i+1)
		}

		if len(listNum) == 2 {
			break
		}
	}

	fmt.Println(int(math.Abs(float64(listNum[0] - listNum[1]))))
}
