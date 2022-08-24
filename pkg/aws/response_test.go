package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetters(t *testing.T) {
	code := 200
	headers := map[string]string{
		"head1": "h1",
	}
	body := "message"

	r := AWSResponse{
		statusCode: code,
		headers:    headers,
		body:       body,
	}

	assert.Equal(t, code, r.StatusCode())
	assert.Equal(t, headers, r.Headers())
	assert.Equal(t, body, r.Body())
}

func TestSetters(t *testing.T) {
	code := 200
	headers := map[string]string{
		"head1": "h1",
	}
	body := "message"

	r := AWSResponse{}

	// Set
	r.SetStatusCode(code)
	r.SetHeaders(headers)
	r.SetBody(body)

	assert.Equal(t, code, r.StatusCode())
	assert.Equal(t, headers, r.Headers())
	assert.Equal(t, body, r.Body())
}
