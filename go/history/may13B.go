package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var apple [2]string
	fmt.Scan(&apple[0], &apple[1])

	n, _ := strconv.Atoi(apple[0])
	l, _ := strconv.Atoi(apple[1])

	minimum := 100000000
	minimumI := -1
	tmpL := l

	for i := 0; i < n; i++ {
		tmp := int(math.Abs(float64(tmpL)))
		if tmp < minimum {
			minimum = tmp
			minimumI = i
		}
		tmpL++
	}

	result := 0
	for i := 0; i < n; i++ {
		if i == minimumI {
			l++
			continue
		}
		result += l
		l++
	}

	fmt.Println(result)
}
