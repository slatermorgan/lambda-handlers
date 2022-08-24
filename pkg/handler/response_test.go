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

	r := mocks.NewResponder(t)
	r.On(
		"SetStatusCode",
		code,
	).Times(1)
	r.On(
		"SetHeaders",
		heads,
	).Times(1)
	r.On(
		"SetBody",
		"model",
	).Times(1)

	l := mocks.NewLogger(t)

	hand := ResponseHandler{
		DefaultHeaders: heads,
		responder:      r,
		logger:         l,
	}

	res, err := hand.BuildResponder(code, body)

	assert.NoError(t, err)
	assert.Implements(t, (*Responder)(nil), res)
}

func TestBuildResponse_Empty(t *testing.T) {
	code := 200
	heads := map[string]string{
		"default": "header",
	}

	r := mocks.NewResponder(t)
	r.On(
		"SetStatusCode",
		code,
	).Times(1)
	r.On(
		"SetHeaders",
		heads,
	).Times(1)
	r.On(
		"SetBody",
		"",
	).Times(1)

	l := mocks.NewLogger(t)

	hand := ResponseHandler{
		DefaultHeaders: heads,
		responder:      r,
		logger:         l,
	}

	res, err := hand.BuildResponse(code, nil)

	assert.NoError(t, err)
	assert.Implements(t, (*Responder)(nil), res)
}

func TestBuildResponse(t *testing.T) {
	model := Model{
		Success: true,
	}

	code := 200
	heads := map[string]string{
		"default": "header",
	}

	r := mocks.NewResponder(t)
	r.On(
		"SetStatusCode",
		code,
	).Times(1)
	r.On(
		"SetHeaders",
		heads,
	).Times(1)
	r.On(
		"SetBody",
		"{\"success\":true}",
	).Times(1)

	l := mocks.NewLogger(t)

	hand := ResponseHandler{
		DefaultHeaders: heads,
		responder:      r,
		logger:         l,
	}

	res, err := hand.BuildResponse(code, model)

	assert.NoError(t, err)
	assert.Implements(t, (*Responder)(nil), res)
}
