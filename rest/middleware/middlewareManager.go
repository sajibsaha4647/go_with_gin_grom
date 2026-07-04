package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type MiddlewareManager struct {
	middlewares []Middleware
}


func NewMiddlewareManager() *MiddlewareManager {
	return &MiddlewareManager{
		middlewares: []Middleware{},
	}
}


func (m *MiddlewareManager) Use(middlewae ...Middleware) {
	m.middlewares = append(m.middlewares, middlewae...)
}

func (m *MiddlewareManager) Apply(handler http.Handler) http.Handler {
	for i := len(m.middlewares) - 1; i >= 0; i-- {
		handler = m.middlewares[i](handler)
	}	

	return handler;
}

func (m *MiddlewareManager) AddMiddleware(next http.Handler, middleware ...Middleware) http.Handler {
	for _, m := range m.middlewares {
		next = m(next)
	}
	return next;
}