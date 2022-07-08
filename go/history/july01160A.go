package main

import (
	"fmt"
)

func main() {
	var str [1]string
	fmt.Scan(&str[0])

	tmp := str[0]

	if tmp[2] == tmp[3] && tmp[4] == tmp[5] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
