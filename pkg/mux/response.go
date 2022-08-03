package mux

type Response struct {
	statusCode int
	headers    map[string]string
	body       string
}

func (r *Response) StatusCode() int {

	return r.statusCode
}

func (r *Response) Headers() map[string]string {
	return r.headers
}

func (r *Response) Body() string {
	return r.body
}

func (r *Response) SetStatusCode(code int) {
	r.statusCode = code
}

func (r *Response) SetHeaders(headers map[string]string) {
	r.headers = headers
}

func (r *Response) SetBody(body string) {
	r.body = body
}
