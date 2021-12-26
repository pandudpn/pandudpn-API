package middleware

import (
	"context"

	"pandudpn/api/src/utils/response"

	"github.com/labstack/echo"
)

const (
	invalidRequest string = "Invalid Request"
	apiKeyNotFound string = "Need API Key to access routes"
	invalidApiKey  string = "API Key invalid"
)

func (c *client) errorResponse(ctx context.Context, e echo.Context, statusCode int, message string, err error) error {
	return response.Errors(ctx, statusCode, message, err).JSON(e)
}

func (c *client) successResponse(ctx context.Context, e echo.Context, statusCode int, data interface{}) error {
	return response.Success(ctx, statusCode, data).JSON(e)
}
