package loom

import (
	"context"
	"net/http"
)

type (
	Operation string
	Pointer   = interface{}
	Vars      = map[string]interface{}
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type ServiceClient interface {
	Do(context.Context, Operation, Vars, Pointer) error
}
