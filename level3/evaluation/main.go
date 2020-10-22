// Statement
// Create many test cases as needed to validate the right behavior of the Age
// Filter exercise from Level 1
package main

func filterByAgeRange(fromAge, toAge int, ages []int) []int {
	response := make([]int, 0)
	for _, ageToCheck := range ages {
		if fromAge <= ageToCheck && ageToCheck <= toAge {
			response = append(response, ageToCheck)
		}
	}
	return response
}
