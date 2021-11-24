package main

/*
 * Complete the 'updateTimes' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY signalOne
 *  2. INTEGER_ARRAY signalTwo
 */

func updateTimes(signalOne []int32, signalTwo []int32) int32 {
	counter := 0
	var max int32
	//first equal
	for i := 0; i < len(signalOne); i++ {
		if (signalOne[i] == signalTwo[i]) && signalOne[i] > max {
			max = signalOne[i]
			counter++
		}
	}
	return int32(counter)
}

func main() {
	var a []int32 = []int32{1, 2, 3, 4, 1}
	var b []int32 = []int32{5, 4, 3, 4, 1}
	print(updateTimes(a, b))
}
