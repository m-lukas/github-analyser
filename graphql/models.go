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
