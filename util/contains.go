package util

import "sort"

func Contains(array []string, compare string) bool {

	for _, item := range array {
		if item == compare {
			return true
		}
	}

	return false
}

func BinaryContains(array []string, compare string) bool {

	sort.Strings(array)
	index := sort.SearchStrings(array, compare)
	if index < len(array) && array[index] == compare {
		return true
	}

	return false
}

func RemoveDuplicates(slice []string) []string {

	keys := make(map[string]bool)
	cleaned := []string{}

	for _, value := range slice {
		exists := keys[value]
		if !exists {
			keys[value] = true
			cleaned = append(cleaned, value)
		}
	}
	return cleaned
}
