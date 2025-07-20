package types

type HttpMethod string

const (
	GET    HttpMethod = HttpMethod("GET")
	POST   HttpMethod = HttpMethod("POST")
	PUT    HttpMethod = HttpMethod("PUT")
	DELETE HttpMethod = HttpMethod("DELETE")
	PATCH  HttpMethod = HttpMethod("PATCH")
)
