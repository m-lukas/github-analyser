package httputil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// UnmarshalRequest reads the request body and parses it into the given pointer
func UnmarshalRequest(r *http.Request, v interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"url":   r.URL.Path,
			"ip":    r.RemoteAddr,
		}).Error("Error reading request body")
		return err
	}

	if err = json.Unmarshal(b, v); err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"url":   r.URL.Path,
		}).Error("Error unmarhsaling request")
		return err
	}

	return nil
}
