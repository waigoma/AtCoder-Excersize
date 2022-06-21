package main

import "fmt"

func isGood(s string) bool {
	tmp := ""
	for _, str := range s {
		if string(str) == tmp {
			return false
		}
		tmp = string(str)
	}
	return true
}

func main() {
	var s [1]string
	fmt.Scan(&s[0])

	if isGood(s[0]) {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
}
