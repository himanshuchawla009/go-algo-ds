package main

import (
	"fmt"
	"math"
)

/**
 * @input A : Integer
 *
 * @Output Integer
 */
func sqrt(A int) int {
	// base cases

	if A == 0 || A == 1 {
		return A
	}

	// binary search will always have four variables
	var max float64 = float64(A)
	var min float64 = 1

	var ans float64

	for {

		if min <= max {
			var mid float64 = float64(math.Floor((max + min) / 2))

			// square of mid
			sqMid := mid * mid
			if sqMid == float64(A) {
				ans = mid
				break
			}
			if sqMid < float64(A) {
				min = mid + 1
				ans = mid
			} else {
				max = mid - 1
			}

		} else {

			break

		}

	}
	return int(math.Floor(ans))

}

func main() {
	ans := sqrt(11)
	fmt.Println(ans)
}
