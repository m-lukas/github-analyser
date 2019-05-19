package db

import "encoding/json"

func ConvertUsers(input []json.RawMessage) ([]*ElasticUser, error) {

	var output []*ElasticUser
	for _, message := range input {
		var userData ElasticUser
		err := json.Unmarshal(message, &userData)
		if err != nil {
			return nil, err
		}
		output = append(output, &userData)
	}

	return output, nil
}
