package middleware

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-http-utils/headers"
	"net/http"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

func (m *GoMiddleware) CORS(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			headers.Accept,
			headers.Authorization,
			headers.ContentType,
			headers.XCSRFToken,
		},
		ExposedHeaders:   []string{headers.Link},
		AllowCredentials: true,
		MaxAge:           300,
	})(next)
}

func (m *GoMiddleware) Heartbeat(next http.Handler) http.Handler {
	return middleware.Heartbeat("/ping")(next)
}
