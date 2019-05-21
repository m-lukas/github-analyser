package httputil

import (
	"encoding/json"
	"net/http"
)

const (
	SUCCESS           = 200
	SERVER_ERROR      = 500
	NOT_FOUND         = 404
	INVALID_ARGUMENTS = 400
)

//ErrorResponse is used for marshaling http errors
type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Message    string `json:"msg"`
}

//NewError creates new ErrorResponse object with statuscode and message
func NewError(statusCode int, message string) *ErrorResponse {
	return &ErrorResponse{StatusCode: statusCode, Message: message}
}

//WriteError writes an http response using the ErrorResponse object
func (e *ErrorResponse) WriteError(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	//check if statuscode is out of range
	if e.StatusCode > 0 {
		w.WriteHeader(e.StatusCode)
	} else {
		w.WriteHeader(SERVER_ERROR)
	}

	bytes, err := json.Marshal(e)
	if err != nil {
		w.WriteHeader(SERVER_ERROR)
		w.Write([]byte("Internal server error!"))
	}

	w.Write(bytes)
}

//WriteSuccess writes and http-response with a given statuscode and a marshaled data body
func WriteSuccess(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	//check if statuscode is out of range
	if statusCode > 0 {
		w.WriteHeader(statusCode)
	} else {
		w.WriteHeader(SERVER_ERROR)
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(SERVER_ERROR)
		w.Write([]byte("Internal server error!"))
	}

	w.Write(bytes)
}
