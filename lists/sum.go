package lists

func Sum(values []int) int {
	var totalValue int

	for _, value := range values {
		totalValue += value
	}

	return totalValue
}

func SumAllTails(numbersLists ...[]int) []int {
	var sums []int

	for _, numberList := range numbersLists {
		if len(numberList) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numberList[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
