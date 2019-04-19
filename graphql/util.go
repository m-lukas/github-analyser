package graphql

import "io/ioutil"

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
