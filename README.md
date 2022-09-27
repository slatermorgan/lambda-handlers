# Lambda Handlers

Lambda Handlers is a go module allowing Serverless handler functions to be ran as a local [Gorilla Mux](https://github.com/gorilla/mux) server or within any cloud server provider event.

Currently supported:
 - AWS Lambda.
 - Standard library HTTP.

## Usage

The first step is to swap our your CSP specific event request and response objects with the generic `Requester` and `Responder` interfaces defined in the handler package of this module.

```go
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
	Find(query string) (interface{}, error)
}

const findHandlerDefaultCount = 10

type AfterFindHandlerHook func(interface{}) error
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

		postcode := request.QueryByName("query")

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

```

In the case where you want to run this handler in a Mux router, call the `CreateHandler` method, pass in the generic handler defined above and pass it into the HandleFunc method on the router.

```go
r := muxRouter.NewRouter()
r.HandleFunc("/test", mux.CreateHandler(handler))

log.Fatal(http.ListenAndServe("localhost:8080", r))
```

In the case where you want to run this handler in AWS Lambda, simply pass the handler into the `Start` method found within the `aws` package of this module.
```go

aws.Start(handler)
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[MIT](https://choosealicense.com/licenses/mit/)
