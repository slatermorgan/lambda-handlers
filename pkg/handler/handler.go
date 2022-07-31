package handler

type Requester interface {
	Body() string
	HeaderByName(name string) string
	PathByName(name string) string
	QueryByName(name string) string
}

type Responder interface {
	StatusCode() int
	Headers() map[string]string
	Body() string
}

type HandlerFunc = func(request Requester) (Responder, error)
