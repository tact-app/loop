package loom

import (
	"context"
	"net/http"
)

type (
	Vars    = map[string]interface{}
	Pointer = interface{}
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type ServiceClient interface {
	Do(context.Context, Operation, Vars, Pointer) error
}
