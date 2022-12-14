package handler

// Generic Request object which is used in every handler
type Requester interface {
	Body() string
	HeaderByName(name string) string
	PathByName(name string) string
	QueryByName(name string) string
	SetQueryByName(name, set string)
	GetAuthToken() string
}

func newResponse(code int, body string) {

}

// Genertic Response object which is used in every handler
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

type Logger interface {
	Error(args ...interface{})
}

// BeforeHandlerHook is a callback function called before a handler functions main logic is ran.
// A Callback function can be passed in when building a handler and is passed the raw API Gateway Request struct
type BeforeHandlerHook func(Requester) error

type HandlerFunc = func(request Requester) (*Response, error)
