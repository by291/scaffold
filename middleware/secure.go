package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func Secure() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			c.Abort()
			return
		}

		// Avoid header rewrite if response is a redirection.
		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}
		c.Next()
	}
}
