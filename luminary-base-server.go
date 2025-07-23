package luminary_base_server

import (
	"context"
	"net/http"

	m "github.com/aether-winds/luminary-base-server/internal/middleware"
	r "github.com/aether-winds/luminary-base-server/internal/route"
	s "github.com/aether-winds/luminary-base-server/internal/server"
	t "github.com/aether-winds/luminary-base-server/internal/types"
)

type Middleware = m.Middleware
type Route = r.Route
type Server = s.Server
type HttpMethod = t.HttpMethod

const (
	HTTP_METHOD_GET    HttpMethod = t.GET
	HTTP_METHOD_POST   HttpMethod = t.POST
	HTTP_METHOD_PUT    HttpMethod = t.PUT
	HTTP_METHOD_DELETE HttpMethod = t.DELETE
	HTTP_METHOD_PATCH  HttpMethod = t.PATCH
)

func CreateMiddleware(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error) Middleware {
	return m.CreateMiddleware(handler)
}

func CreateRoute(
	method HttpMethod,
	path string,
	handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error,
) Route {
	return r.CreateRoute(method, path, handler)
}

func CreateServer(port int) Server {
	return s.CreateServer(port)
}
