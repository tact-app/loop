package loom

import (
	"context"
	"net/http"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type ServiceClient interface {
	Do(context.Context, Operation, map[string]interface{}, interface{}) error
}
