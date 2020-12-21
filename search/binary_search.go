package search

func binary(search int, list []int) bool {
	midIndex := len(list) / 2
	mid := list[midIndex]

	if search == mid {
		return true
	} else if len(list) == 1 {
		return false
	} else if search <= mid {
		return binary(search, list[:midIndex])
	} else if search >= mid {
		return binary(search, list[midIndex:])
	}

	return false
}
