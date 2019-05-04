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
