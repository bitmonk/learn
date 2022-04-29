package arrayslice

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum...[]int) (sums []int) {
	return
}