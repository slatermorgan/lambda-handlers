package handler

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/oneiota/serviceerror"
)

// Generic Request object which is used in every handler
type Requester interface {
	Body() string
	HeaderByName(name string) string
	PathByName(name string) string
	QueryByName(name string) string
	GetAuthToken() string
}

// Genertic Response object which is used in every handler
type Responder interface {
	StatusCode() int
	Headers() map[string]string
	Body() string
	SetStatusCode(code int)
	SetHeaders(headers map[string]string)
	SetBody(body string)
}

type logger interface {
	Error(args ...interface{})
}

// Genertic Handler object which is the reciever in every handler method
type ResponseHandler struct {
	DefaultHeaders map[string]string
	responder      Responder
	logger         logger
}

func NewResponseHandler(logger logger, responder Responder) *ResponseHandler {
	return &ResponseHandler{
		logger:    logger,
		responder: responder,
	}
}

// Body gets request payload
func (r *ResponseHandler) BuildResponse(code int, model interface{}) (Responder, error) {
	body := ""
	if model != nil {
		bodyBytes, err := json.Marshal(model)
		if err != nil {
			r.responder.SetStatusCode(http.StatusInternalServerError)

			return r.responder, err
		}

		body = string(bodyBytes)
	}

	return r.BuildResponder(code, body)
}

// BuildRawJSONResponse builds an Response with the given status code & response body
// The Response will contain the raw response body and appropriate JSON header
func (r *ResponseHandler) BuildResponder(statusCode int, respBody string) (Responder, error) {
	r.responder.SetStatusCode(code)
	r.responder.SetHeaders(r.DefaultHeaders)
	r.responder.SetBody(string(body))

	return r.responder, nil
}

func (r *ResponseHandler) BuildErrorResponse(err error) (Responder, error) {
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

// BeforeHandlerHook is a callback function called before a handler functions main logic is ran.
// A Callback function can be passed in when building a handler and is passed the raw API Gateway Request struct
type BeforeHandlerHook func(Requester) error

type HandlerFunc = func(request Requester) (Responder, error)
