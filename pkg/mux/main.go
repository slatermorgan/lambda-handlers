package mux

import (
	"io"
	"net/http"

	"github.com/slatermorgan/lambda-handlers/pkg/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

func CreateHandler(
	h handler.HandlerFunc,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := h(NewRequest(r))

		WriteResponse(res, w)
	}
}

func WriteResponse(r handler.Responder, w http.ResponseWriter) {
	w.Write([]byte(r.Body()))
}
