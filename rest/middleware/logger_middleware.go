package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the request details
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)

	}

}

// func Logger() Middleware {
//     return func(next http.Handler) http.Handler {
//         return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//             fmt.Println("Request:", r.Method, r.URL.Path)

//             next.ServeHTTP(w, r)
//         })
//     }
// }
