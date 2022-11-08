package response

import (
	"encoding/json"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

type ResponseObject struct {
	// StatusCode int
	Data interface{}
}

func Response(status int, data interface{}, rw http.ResponseWriter) {
	respBytes, err := json.Marshal(ResponseObject{data})
	if err != nil {
		logger.Error(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
