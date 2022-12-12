package handler

import (
	"testing"

	"github.com/slatermorgan/lambda-handlers/pkg/handler/mocks"
	"github.com/stretchr/testify/assert"
)

type Model struct {
	Success bool `json:"success"`
}

func TestBuildResponder(t *testing.T) {
	body := "model"
	code := 200
	heads := map[string]string{
		"default": "header",
	}

	l := mocks.NewLogger(t)

	hand := NewResponseHandler(l, heads)

	res, err := hand.BuildResponder(code, body)

	assert.NoError(t, err)
	assert.IsType(t, (*Response)(nil), res)
}

func TestBuildResponse_Empty(t *testing.T) {
	code := 200
	heads := map[string]string{
		"default": "header",
	}

	l := mocks.NewLogger(t)

	hand := NewResponseHandler(l, heads)

	res, err := hand.BuildResponse(code, nil)

	assert.NoError(t, err)
	assert.IsType(t, (*Response)(nil), res)
}

func TestBuildResponse(t *testing.T) {
	model := Model{
		Success: true,
	}

	code := 200
	heads := map[string]string{
		"default": "header",
	}

	l := mocks.NewLogger(t)

	hand := NewResponseHandler(l, heads)

	res, err := hand.BuildResponse(code, model)

	assert.NoError(t, err)
	assert.IsType(t, (*Response)(nil), res)
}
