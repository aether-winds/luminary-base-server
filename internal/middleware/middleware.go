package middleware

import (
	"context"
	"net/http"
)

type Middleware interface {
	Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

type middleware struct {
	handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

func (m *middleware) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return m.handler(ctx, w, r)
}

func CreateMiddleware(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error) Middleware {
	return &middleware{
		handler: handler,
	}
}
