package katas

// LargestNumbers returns a slice with the largest number from each of the lists of numbers passed to it
func LargestNumbers(listOfLists [][]int) []int {
	var result []int
	for _, list := range listOfLists {
		max := list[0]
		for _, n := range list {
			if n > max {
				max = n
			}
		}
		result = append(result, max)
	}
	return result
}
