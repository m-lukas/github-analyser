package graphql

import "io/ioutil"

func ReadQuery(path string) (string, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil

}
