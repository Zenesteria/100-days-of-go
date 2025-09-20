package forloops

func contains(arr []int, target int) bool {
	result := false
	for _, val := range arr {
		if val == target {
			result = true
		}
	}
	return result
}

func TotalFruit(fruits []int) int {

	maxLength := 0

	for i := range fruits {
		var basketOne = []int{}
		var basketTwo = []int{}
		for j := i; j < len(fruits); j++ {
			if contains(basketOne, fruits[j]) || len(basketOne) == 0 {
				basketOne = append(basketOne, fruits[j])
				continue
			}
			if contains(basketTwo, fruits[j]) || len(basketTwo) == 0 {
				basketTwo = append(basketTwo, fruits[j])
				continue
			}

			break

		}
		if len(basketTwo)+len(basketOne) > maxLength {
			maxLength = len(basketTwo) + len(basketOne)
		}
	}

	// Replace this placeholder return statement with your code
	return maxLength
}
