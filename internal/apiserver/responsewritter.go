package apiserver

import "net/http"

type response struct {
	writer http.ResponseWriter
	code   int
}

func (r *response) WriteHeader(statusCode int) {
	r.code = statusCode
	r.writer.WriteHeader(statusCode)
}
