package codewars

import (
	"fmt"
	"math"
)

func Codewars() {
	fmt.Println(EquableTriangle(5, 12, 13))
	//LengthOfSequence([]int{-79, 9, -79}, -79)
	Solution("13", "200")
}

// codewars
func EquableTriangle(a, b, c int) bool {
	var perimeter = float64(a + b + c)
	p := perimeter / 2
	square := math.Sqrt(p * (p - float64(a)*(p-float64(b)*(p-float64(c)))))
	//fmt.Println(square == perimeter)
	return square == perimeter
}

func combat(health, damage float64) float64 {
	result := health - damage
	if result < 0 {
		return 0.0
	}
	return result
}

func LengthOfSequence(arr []int, key int) int {
	var first, last, count int

	for index, element := range arr {
		if element == key {
			fmt.Println(index, element)
			if count == 0 {
				first = index
			} else if count == 1 {
				last = index
			}
			count++
		}
	}
	if count != 2 {
		return 0
	}
	return last - first + 1
}

func Solution(a, b string) string {
	if len(a) < len(b) {
		return a + b + a

	}
	return b + a + b
}
