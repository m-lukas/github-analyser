package httputil

import (
	"encoding/json"
	"net/http"

	"github.com/m-lukas/github-analyser/translate"

	log "github.com/sirupsen/logrus"
)

// ErrorResponse can be written to http response writers
type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Message    string `json:"msg"`
}

// FromTranslationKey creates a new error object by translating the given key to german
func FromTranslationKey(statusCode int, key translate.Key) *ErrorResponse {
	t := translate.Get(translate.German, key)

	return &ErrorResponse{
		Message:    t,
		StatusCode: statusCode,
	}
}

// Write the error to a http response writer
func (resp *ErrorResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	if resp.StatusCode > 0 {
		w.WriteHeader(resp.StatusCode)
	} else {
		w.WriteHeader(500)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"data":  resp,
		}).Error("Unable to marshal error response")

		return
	}
	w.Write(b)
}

// WriteSuccess writes the given data to a http writer
func WriteSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"data":  data,
		}).Error("Unable to marshal response")
		FromTranslationKey(500, translate.ServerError).Write(w)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.WriteHeader(statusCode)
	w.Write(b)
}

// WriteBlankSuccess writes headers but does not send any payload
func WriteBlankSuccess(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.WriteHeader(statusCode)
}
