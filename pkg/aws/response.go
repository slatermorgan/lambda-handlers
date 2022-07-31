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

func NewAWSResponse(r handler.Responder) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: r.StatusCode(),
		Headers:    r.Headers(),
		Body:       r.Body(),
	}
}
