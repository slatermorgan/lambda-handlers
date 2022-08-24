package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
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
