package main

import (
	"fmt"
	"strconv"
)

func main() {
	var time [4]string
	fmt.Scan(&time[0], &time[1], &time[2], &time[3])

	takahashi := time[0] + "." + time[1]
	aoki := time[2] + "." + time[3] + "1"

	takahashiF, _ := strconv.ParseFloat(takahashi, 64)
	aokiF, _ := strconv.ParseFloat(aoki, 64)

	if takahashiF < aokiF {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
