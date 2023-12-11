package lists

func Sum(values [5]int) int {
	var totalValue int

	for _, value := range values {
		totalValue += value
	}

	return totalValue
}
