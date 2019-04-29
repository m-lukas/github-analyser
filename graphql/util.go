package graphql

import "io/ioutil"

type rateLimit struct {
	Cost      int
	Remaining int
}

func readQuery(path string) (string, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil

}

func avg(slice []float64) float64 {

	var sum float64

	for _, value := range slice {

		sum += value

	}

	return sum / float64(len(slice))

}

func removeDuplicates(slice []string) []string {

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
