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
	nmSplit := strings.Split(scanner.Text(), " ")
	//n, _ := strconv.Atoi(nmSplit[0])
	m, _ := strconv.Atoi(nmSplit[1])

	scanner.Scan()
	voteSplit := strings.Split(scanner.Text(), " ")

	totalVote := 0

	for _, v := range voteSplit {
		num, _ := strconv.Atoi(v)
		totalVote += num
	}

	selected := 0

	for i := 0; i < len(voteSplit); i++ {
		num, _ := strconv.Atoi(voteSplit[i])
		tmp := float32(num)
		tmp2 := float32(totalVote) * float32(1) / float32(4*m)

		if tmp < tmp2 {
			continue
		}
		selected++
	}

	if selected < m {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
