package main

import (
	"fmt"
	"math"
	"strconv"
)

func calcCeilFloor(x, r float64) (float64, float64) {
	_ceil := math.Ceil(x - r)
	_floor := math.Floor(x + r)
	return _ceil, _floor
}

func main() {
	var str [3]string
	fmt.Scan(&str[0], &str[1], &str[2])

	x, _ := strconv.ParseFloat(str[0], 64)
	y, _ := strconv.ParseFloat(str[1], 64)
	r, _ := strconv.ParseFloat(str[2], 64)

	if x > y {
		x, y = y, x
	}

	count := 0

	_ceil, _floor := calcCeilFloor(x, r)

	for i := _ceil; i <= _floor; i++ {
		newPoint := math.Sqrt(math.Pow(r, 2) - math.Pow(x-i, 2))

		if newPoint == 0 {
			if math.Floor(y) == y {
				count++
			}
			continue
		}

		_bottom, _top := calcCeilFloor(y, newPoint)
		count += int(math.Abs(_top-_bottom)) + 1

		if math.Abs(_bottom-_top) == r*2 {
			if math.Floor(x) != x {
				count -= 2
			}
		}
	}

	fmt.Println(count)
}
