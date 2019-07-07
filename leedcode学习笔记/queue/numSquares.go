package main

import "fmt"
import "math"

func main() {
	fmt.Println(48, numSquares(48))
	//	fmt.Println(17, numSquares(17))
	//	fmt.Println(19, numSquares(19))
	//	fmt.Println(20, numSquares(20))
}

func numSquares(n int) int {
	min := int((^uint(0) >> 1))
	max := int(math.Floor(math.Sqrt(float64(n))))
	max = 4
	for m := max; m > 0; m-- {
		t := m
		numTmp := 0
		yu := n
		for t > 0 {
			fmt.Println("111111", yu, t)
			if yu-t*t < 0 {
				continue
			}
			yu -= t * t
			fmt.Println("222222", yu, t)
			t = int(math.Floor(math.Sqrt(float64(yu))))
			fmt.Println("444444", yu, t)
			numTmp++
		}
		fmt.Println("33333", min, numTmp)
		if min > numTmp {
			min = numTmp
		}
	}
	return min
}
