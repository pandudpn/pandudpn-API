package middleware

import (
	"net/http"
	
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type client struct {
	log *logrus.Logger
}

type MiddlewareInterface interface {
	// Logger initialize first new context from request until get response
	// and will print out to terminal with json formatted
	Logger(next http.Handler) http.Handler
	// APIKey middleware for checking `x-api-key` on header with our configuration
	APIKey(next echo.HandlerFunc) echo.HandlerFunc
}

func NewMiddleware(log *logrus.Logger) MiddlewareInterface {
	return &client{
		log: log,
	}
}
