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

// PathByName gets a query parameter by its name eg. "locale"
func (r *AWSRequest) QueryByName(name string) string {
	return r.headers[name]
}
