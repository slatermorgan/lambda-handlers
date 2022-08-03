package example

import (
	"net/http"

	"github.com/slatermorgan/lambda-handlers/pkg/handler"
)

type ExampleModel struct {
	Success bool `json:"success"`
}

type Connector interface {
	Authorize(token string) error
	Find(postcode string) (interface{}, error)
}

const findHandlerDefaultCount = 10

// AfterFindHandlerHook is a hook/callback function definition, triggered after the Find connector call on for the FindHandler
type AfterFindHandlerHook func(interface{}) error

// FindHandler returns a handlers.HandlerFunc which is used for the Find endpoint.
// The handler calls the Find method of the connector
func FindHandler(
	resHander *handler.ResponseHandler,
	connector Connector,
	beforeHook handler.BeforeHandlerHook,
	afterHook AfterFindHandlerHook,
) handler.HandlerFunc {
	return func(request handler.Requester) (handler.Responder, error) {
		if beforeHook != nil {
			if err := beforeHook(request); err != nil {
				return resHander.BuildErrorResponse(err)
			}
		}

		token := request.GetAuthToken()
		if err := connector.Authorize(token); err != nil {
			return resHander.BuildErrorResponse(err)
		}

		postcode := request.QueryByName("postcode")

		addresses, err := connector.Find(postcode)
		if err != nil {
			return resHander.BuildErrorResponse(err)
		}

		if afterHook != nil {
			if err := afterHook(addresses); err != nil {
				return resHander.BuildErrorResponse(err)
			}
		}

		return resHander.BuildResponse(http.StatusOK, addresses)
	}
}
