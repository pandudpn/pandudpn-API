package controller

import "github.com/labstack/echo"

type JobControllerInterface interface {
	FindHandler(e echo.Context) error
}

type VisitorControllerInterface interface {
	NewVisitorHandler(c echo.Context) error
}
