package solve

func abs(x, y int) int {
	if x > y {
		return 1
	}
	return -1
}

func CalculateSimilarity(col1, col2 []int) int {
	var sim int = 0
	for _, elem1 := range col1 {
		var count int = 0
		for _, elem2 := range col2 {
			if elem1 == elem2 {
				count++
			}
		}
		sim += (elem1 * count)
	}

	return sim
}

func CalculateDistance(col1, col2 []int) int {
	var dist int = 0
	for i := range col1 {
		dist += ((col1[i] - col2[i]) * abs(col1[i], col2[i]))
	}

	return dist
}
