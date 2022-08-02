package aws

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/slatermorgan/lambda-handlers/pkg/handler"
)

type AWSResponse struct {
	statusCode int
	headers    map[string]string
	body       string
}

func (a *AWSResponse) StatusCode() int {
	return a.statusCode
}

func (a *AWSResponse) Headers() map[string]string {
	return a.headers
}

func (a *AWSResponse) Body() string {
	return a.body
}

func (a *AWSResponse) SetStatusCode(code int) {
	a.statusCode = code
}

func (a *AWSResponse) SetHeaders(headers map[string]string) {
	a.headers = headers
}

func (a *AWSResponse) SetBody(body string) {
	a.body = body
}

func NewEvent(r handler.Responder) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: r.StatusCode(),
		Headers:    r.Headers(),
		Body:       r.Body(),
	}
}
