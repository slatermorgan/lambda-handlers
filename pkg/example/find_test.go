package example

import (
	"testing"

	"bitbucket.org/oneiota/mesh-connect-models/address"
	"github.com/aws/aws-lambda-go/events"
	"github.com/slatermorgan/lambda-handlers/internal/mocks"
	"github.com/slatermorgan/lambda-handlers/pkg/aws"
	"github.com/slatermorgan/lambda-handlers/pkg/handler"
	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus/hooks/test"
)

func TestFind_AWS(t *testing.T) {
	expectToken := "authToken"
	expectAddrs := []*address.Address{}
	expectQuery := "M36FJ"
	awsReq := &events.APIGatewayProxyRequest{
		Path: "test/123",
		QueryStringParameters: map[string]string{
			"postcode": expectQuery,
		},
		Headers: map[string]string{
			"Accept":        "application/json",
			"Authorization": expectToken,
		},
	}
	req := aws.NewAWSRequest(awsReq)

	// Mocks
	c := new(mocks.Connector)
	c.On("Authorize",
		expectToken,
	).Return(
		nil,
	).Times(1)

	c.On("Find",
		expectQuery,
	).Return(
		expectAddrs,
		nil,
	).Times(1)

	logger, _ := test.NewNullLogger()

	resHander := handler.NewResponseHandler(logger, &aws.AWSResponse{})

	// Asserts
	resp, err := FindHandler(resHander, c, nil, nil)(req)
	assert.NoError(t, err)

	awsRes := aws.NewEvent(resp)
	expectAwsRes := &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "application/json"},
		MultiValueHeaders: map[string][]string(nil),
		Body:              "[]",
		IsBase64Encoded:   false,
	}

	assert.IsType(t, &events.APIGatewayProxyResponse{}, awsRes)
	assert.Equal(t, expectAwsRes, awsRes)

}
