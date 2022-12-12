package aws

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestRequestGetters(t *testing.T) {
	body := "abcdef"

	token := "abc.def.ghi"

	headName := "header"
	headVal := "h1"

	pathKey := "id"
	pathVal := "p1"

	path2Key := "subid"
	path2Val := "p2"

	queryKey := "q"
	queryVal := "football"

	query2Key := "t"
	query2Val := "red"

	req := AWSRequest{
		body: body,
		headers: map[string]string{
			headName:        headVal,
			"Authorization": token,
		},
		pathParams: map[string]string{
			pathKey:  pathVal,
			path2Key: path2Val,
		},
		queryParams: map[string]string{
			queryKey:  queryVal,
			query2Key: query2Val,
		},
	}

	assert.Equal(t, body, req.Body())

	assert.Equal(t, headVal, req.HeaderByName(headName))

	assert.Equal(t, token, req.HeaderByName("Authorization"))

	assert.Equal(t, pathVal, req.PathByName(pathKey))

	assert.Equal(t, queryVal, req.QueryByName(queryKey))
}

func TestBody_Empty(t *testing.T) {
	req := AWSRequest{}

	assert.Equal(t, "", req.Body())
	assert.Equal(t, "", req.HeaderByName("headName"))
	assert.Equal(t, "", req.HeaderByName("Authorization"))
	assert.Equal(t, "", req.PathByName("pathKey"))
	assert.Equal(t, "", req.QueryByName("queryKey"))
}

func TestSetQueryByName(t *testing.T) {
	queryKey := "q"
	queryVal := "football"

	query2Key := "t"
	query2Val := "red"

	newQueryVal := "soccer"

	req := AWSRequest{
		queryParams: map[string]string{
			queryKey:  queryVal,
			query2Key: query2Val,
		},
	}

	// BEFORE
	assert.Equal(t, queryVal, req.QueryByName(queryKey))
	assert.Equal(t, query2Val, req.QueryByName(query2Key))

	req.SetQueryByName(queryKey, newQueryVal)
	// AFTER
	assert.Equal(t, newQueryVal, req.QueryByName(queryKey))
	assert.Equal(t, query2Val, req.QueryByName(query2Key))
}

func TestNewAWSRequest(t *testing.T) {
	pathKey := "pathKey"
	pathVal := "path"
	queryKey := "q"
	queryVal := "football"
	body := "body here"
	headKey := "headerKey"
	headVal := "headerValue"

	req := &events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			queryKey: queryVal,
		},
		PathParameters: map[string]string{
			pathKey: pathVal,
		},
		Body: body,
		Headers: map[string]string{
			headKey: headVal,
		},
	}

	actual := NewAWSRequest(req)

	// BEFORE
	assert.Equal(t, body, actual.Body())
	assert.Equal(t, pathVal, actual.PathByName(pathKey))
	assert.Equal(t, headVal, actual.HeaderByName(headKey))
	assert.Equal(t, queryVal, actual.QueryByName(queryKey))
}
