package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var str [2]string
	fmt.Scan(&str[0], &str[1])

	n, _ := strconv.Atoi(str[0])
	k, _ := strconv.Atoi(str[1])

	if n%k == 0 || n == 0 || n == k {
		fmt.Println(0)
		os.Exit(0)
	}

	divide := n / k

	tmp := n - k*divide

	tmp1 := math.Abs(float64(tmp))
	tmp2 := math.Abs(float64(tmp - k))

	if tmp1 < tmp2 {
		fmt.Println(int(tmp1))
	} else {
		fmt.Println(int(tmp2))
	}
}
