package mux

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"
)

type Request struct {
	request *http.Request
}

func NewRequest(r *http.Request) *Request {
	return &Request{
		request: r,
	}
}

// Body gets request payload
func (r *Request) Body() string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.request.Body)

	return buf.String()
}

// HeaderByName gets a header by its name eg. "content-type"
func (r *Request) HeaderByName(name string) string {
	head := r.request.Header

	return head.Get(name)
}

// PathByName gets a path parameter by its name eg. "productID"
func (r *Request) PathByName(name string) string {
	vars := mux.Vars(r.request)

	return vars[name]
}

// PathByName gets a query parameter by its name eg. "locale"
func (r *Request) QueryByName(name string) string {
	v := r.request.URL.Query()

	return v.Get(name)
}

// PathByName gets a query parameter by its name eg. "locale"
func (r *Request) GetAuthToken() string {
	if r.HeaderByName("Authorization") != "" {
		return r.HeaderByName("Authorization")
	} else {
		return r.HeaderByName("authorization")
	}
}
