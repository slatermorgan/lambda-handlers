package aws

import "github.com/aws/aws-lambda-go/events"

type AWSRequest struct {
	body        string
	pathParams  map[string]string
	queryParams map[string]string
	headers     map[string]string
}

func NewAWSRequest(r *events.APIGatewayProxyRequest) *AWSRequest {
	return &AWSRequest{
		body:        r.Body,
		pathParams:  r.PathParameters,
		queryParams: r.QueryStringParameters,
		headers:     r.Headers,
	}
}

// Body gets request payload
func (r *AWSRequest) Body() string {
	return r.body
}

// HeaderByName gets a header by its name eg. "content-type"
func (r *AWSRequest) HeaderByName(name string) string {
	return r.headers[name]
}

// PathByName gets a path parameter by its name eg. "productID"
func (r *AWSRequest) PathByName(name string) string {
	return r.pathParams[name]
}

// QueryByName gets a query parameter by its name eg. "locale"
func (r *AWSRequest) QueryByName(name string) string {
	return r.queryParams[name]
}

// PathByName sets a query parameter by its name eg. "locale"
// This is used to alter requests in middleware functions.
func (r *AWSRequest) SetQueryByName(name, set string) {
	r.queryParams[name] = set
}

// PathByName gets a query parameter by its name eg. "locale"
func (r *AWSRequest) GetAuthToken() string {
	if r.HeaderByName("Authorization") != "" {
		return r.HeaderByName("Authorization")
	} else {
		return r.HeaderByName("authorization")
	}
}
