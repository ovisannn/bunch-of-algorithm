package main

import "fmt"

func Bday_cake(ar []int32) int32 {
	result := 0
	max := ar[0]

	for i := 1; i < len(ar); i++ {
		if ar[i] > max {
			max = ar[i]
		}
	}

	for i := 0; i < len(ar); i++ {
		if ar[i] == max {
			result++
		}
	}
	return int32(result)
}

func main() {
	fmt.Print(Bday_cake([]int32{3, 2, 1, 3}))
}
