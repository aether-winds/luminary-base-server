package main

import (
	"context"
	"net/http"

	"github.com/aether-winds/luminary-base-server/internal/middleware"
	"github.com/aether-winds/luminary-base-server/internal/route"
	"github.com/aether-winds/luminary-base-server/internal/server"
	"github.com/aether-winds/luminary-base-server/internal/types"
)

type Middleware = middleware.Middleware
type Route = route.Route
type Server = server.Server
type HttpMethod = types.HttpMethod

const (
	HTTP_METHOD_GET    HttpMethod = HttpMethod(types.GET)
	HTTP_METHOD_POST   HttpMethod = HttpMethod(types.POST)
	HTTP_METHOD_PUT    HttpMethod = HttpMethod(types.PUT)
	HTTP_METHOD_DELETE HttpMethod = HttpMethod(types.DELETE)
	HTTP_METHOD_PATCH  HttpMethod = HttpMethod(types.PATCH)
)

func CreateMiddleware(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error) Middleware {
	return middleware.CreateMiddleware(handler)
}

func CreateRoute(
	method HttpMethod,
	path string,
	handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error,
) Route {
	return route.CreateRoute(method, path, handler)
}

func CreateServer(port int) Server {
	return server.CreateServer(port)
}
