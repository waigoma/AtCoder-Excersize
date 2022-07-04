package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func strListToIntList(strList []string) []int {
	intList := make([]int, len(strList))
	for i, str := range strList {
		intList[i], _ = strconv.Atoi(str)
	}
	return intList
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nm := scanner.Text()

	scanner.Scan()
	an := scanner.Text()

	scanner.Scan()
	bn := scanner.Text()

	nmStrList := strings.Split(nm, " ")
	anStrList := strings.Split(an, " ")
	bnStrList := strings.Split(bn, " ")

	n, _ := strconv.Atoi(nmStrList[0])
	m, _ := strconv.Atoi(nmStrList[1])

	//anStrList = anStrList[:n]
	//bnStrList = bnStrList[:m]

	//anList := strListToIntList(anStrList)
	//bnList := strListToIntList(bnStrList)
	//
	min := math.MaxInt64
	//for _, a := range anList {
	//	for _, b := range bnList {
	//		tmp := int(math.Abs(float64(a - b)))
	//		if tmp < min {
	//			min = tmp
	//		}
	//	}
	//}

	fmt.Println(anStrList, bnStrList, n, m, min)
}
