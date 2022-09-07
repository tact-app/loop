package loom

import (
	"context"
	"net/http"
)

type (
	Operation string
	Pointer   interface{}
	Vars      map[string]interface{}
)

func (src Vars) Set(k string, v interface{}) Vars {
	dst := make(Vars, len(src))
	for k, v := range src {
		dst[k] = v
	}
	src[k] = v
	return src
}

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type ServiceClient interface {
	Do(context.Context, Operation, Vars, Pointer) error
}
