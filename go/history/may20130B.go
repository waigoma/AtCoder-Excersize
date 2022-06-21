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

	split1 := strings.Split(scanner.Text(), " ")

	n, _ := strconv.Atoi(split1[0])
	x, _ := strconv.Atoi(split1[1])

	scanner.Scan()

	split := strings.Split(scanner.Text(), " ")

	l := make([]int, 0, len(split))
	for _, str := range split {
		atoint, _ := strconv.Atoi(str)
		l = append(l, atoint)
	}

	ls := 0
	count := 0

	tmp := make([]int, 0, len(l))

	for i := 0; i < n+1; i++ {
		tmp = append(tmp, ls)
		if len(l) > i {
			ls += l[i]
		}
	}

	for _, i := range tmp {
		if i <= x {
			count++
		}
	}

	fmt.Println(count)
}
