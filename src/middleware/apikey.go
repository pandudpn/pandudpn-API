package middleware

import (
	"fmt"
	"net/http"
	"reflect"

	"pandudpn/api/src/utils/logger"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func (c *client) APIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			req    = e.Request()
			ctx    = req.Context()
			key    = req.Header.Get("x-api-key")
			apiKey = viper.GetString("API_KEY")
			err    error
		)

		if reflect.ValueOf(key).IsZero() {
			err = fmt.Errorf("api key not found in header")
			logger.Log.Error(ctx, err.Error())

			return c.errorResponse(ctx, e, http.StatusForbidden, apiKeyNotFound, err)
		}

		if key != apiKey {
			err = fmt.Errorf("api key not same with configuration")
			logger.Log.Error(ctx, err.Error())

			return c.errorResponse(ctx, e, http.StatusForbidden, invalidApiKey, err)
		}

		return next(e)
	}
}
