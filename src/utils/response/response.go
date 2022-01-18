package response

import (
	"bytes"
	"context"
	"encoding/json"

	"pandudpn/api/src/utils/logger"

	"github.com/labstack/echo"
)

type response struct {
	Status       bool        `json:"status"`
	StatusCode   int         `json:"status_code"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

type OutputResponseInterface interface {
	// JSON will send the response without encrypted
	JSON(c echo.Context) error
}

func (r *response) json(c echo.Context) error {
	return c.JSON(r.StatusCode, r)
}

func (r *response) JSON(c echo.Context) error {
	return c.JSON(r.StatusCode, r)
}

func (r *response) marshal(data interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(data)
	return buffer.Bytes(), err
}

func Success(ctx context.Context, statusCode int, data interface{}) OutputResponseInterface {
	r := &response{}
	r.Status = true
	r.StatusCode = statusCode
	logger.Response(ctx, statusCode, r, nil)
	r.Data = data

	return r
}

func Errors(ctx context.Context, statusCode int, message string, err error) OutputResponseInterface {
	r := &response{}
	r.StatusCode = statusCode
	r.Status = false
	r.ErrorMessage = message

	logger.Response(ctx, statusCode, nil, err)
	return r
}
