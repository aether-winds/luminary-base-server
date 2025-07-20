package route

import (
	"context"
	"net/http"

	"github.com/aether-winds/luminary-base-server/internal/middleware"
	"github.com/aether-winds/luminary-base-server/internal/types"
)

type Route interface {
	middleware.Middleware
	GetMethod() types.HttpMethod
	GetPath() string
}

type route struct {
	Route
	method  types.HttpMethod
	path    string
	handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

func (r *route) GetMethod() types.HttpMethod {
	return r.method
}

func (r *route) GetPath() string {
	return r.path
}

func (r *route) Handler(ctx context.Context, w http.ResponseWriter, re *http.Request) error {
	return r.handler(ctx, w, re)
}

func CreateRoute(
	method types.HttpMethod,
	path string,
	handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error,
) Route {
	return &route{
		method:  method,
		path:    path,
		handler: handler,
	}
}
