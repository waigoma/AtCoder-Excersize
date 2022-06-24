package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	scanner.Text()

	scanner.Scan()
	str := scanner.Text()

	fmt.Println(strings.Count(str, "ABC"))
}
