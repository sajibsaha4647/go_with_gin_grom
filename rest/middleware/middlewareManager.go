package middleware

import "github.com/gin-gonic/gin"

// type Middleware func(http.Handler) http.Handler

// type MiddlewareManager struct {
// 	middlewares []Middleware
// }

// func NewMiddlewareManager() *MiddlewareManager {
// 	return &MiddlewareManager{
// 		middlewares: []Middleware{},
// 	}
// }

// func (m *MiddlewareManager) Use(middlewae ...Middleware) {
// 	m.middlewares = append(m.middlewares, middlewae...)
// }

// func (m *MiddlewareManager) Apply(handler http.Handler) http.Handler {
// 	for i := len(m.middlewares) - 1; i >= 0; i-- {
// 		handler = m.middlewares[i](handler)
// 	}

// 	return handler
// }

// func (m *MiddlewareManager) AddMiddleware(next http.Handler, middleware ...Middleware) http.Handler {
// 	// 1. Append the new incoming middlewares to your internal slice
// 	m.Use(middleware...)
// 	m.Apply(next)
// 	return next
// }


//==== with gin framework
type Middleware = gin.HandlerFunc

type MiddlewareManager struct {
	middlewares []Middleware
}

func NewMiddlewareManager() *MiddlewareManager {
	return &MiddlewareManager{
		middlewares: []Middleware{},
	}
}

func (m *MiddlewareManager) Use(mw ...Middleware) {
	m.middlewares = append(m.middlewares, mw...)
}

func (m *MiddlewareManager) Middlewares() []gin.HandlerFunc {
	return m.middlewares
}
