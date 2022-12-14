package handler

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/oneiota/serviceerror"
)

// Genertic Handler object which is the reciever in every handler method
type ResponseHandler struct {
	defaultHeaders map[string]string
	logger         Logger
}

func NewResponseHandler(
	logger Logger,
	defaultHeads map[string]string,
) *ResponseHandler {
	return &ResponseHandler{
		logger:         logger,
		defaultHeaders: defaultHeads,
	}
}

// Body gets request payload
func (r *ResponseHandler) BuildResponse(code int, model interface{}) (*Response, error) {
	body := ""
	if model != nil {
		bodyBytes, err := json.Marshal(model)
		if err != nil {
			return &Response{StatusCode: http.StatusInternalServerError}, err
		}

		body = string(bodyBytes)
	}

	return r.BuildResponder(code, body)
}

// BuildRawJSONResponse builds an Response with the given status code & response body
// The Response will contain the raw response body and appropriate JSON header
func (r *ResponseHandler) BuildResponder(code int, body string) (*Response, error) {
	return &Response{
		StatusCode: code,
		Body:       body,
		Headers:    r.defaultHeaders,
	}, nil
}

func (r *ResponseHandler) BuildErrorResponse(err error) (*Response, error) {
	statusCode := http.StatusInternalServerError
	var serviceErr *serviceerror.ServiceError

	switch err := err.(type) {
	case *serviceerror.ServiceError:
		statusCode = err.StatusCode()
		serviceErr = err
	default:
		// If its a general error - we don't want to return the message as its a code/integration issue.
		// We don't want those messages being shown to users.
		serviceErr = serviceerror.Unknown("An unknown error occurred")
	}

	if statusCode == http.StatusInternalServerError {
		r.logger.Error(err)
	}

	return r.BuildResponse(statusCode, serviceErr)
}
