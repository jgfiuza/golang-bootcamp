package main

func filter(arr []int, min, max int) (result []int) {
	if len(arr) == 0 {
		return
	}
	if min > max {
		min, max = max, min
	}
	for _, num := range arr {
		if num >= min && num <= max {
			result = append(result, num)
		}
	}
	return
}
