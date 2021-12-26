package middleware

import (
	"net/http"

	"pandudpn/api/src/utils/logger"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (c *client) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, data := logger.Initialize(r)

		next.ServeHTTP(w, req)

		data.Finalize(req.Context())
		c.write(&data)
	})
}

func (c *client) write(dl *logger.DataLogger) {
	var (
		level         logrus.Level
		elasticStatus = viper.GetBool("ELASTIC_ENABLED")
		errChan       = make(chan error, 1)
	)

	if elasticStatus {
		select {
		case err := <-errChan:
			logrus.Errorf("error send data to elastic %v", err)
			break
		default:
			close(errChan)
			break
		}
	}

	if dl.StatusCode >= 200 && dl.StatusCode < 400 {
		level = logrus.InfoLevel
	} else if dl.StatusCode >= 400 && dl.StatusCode < 500 {
		level = logrus.WarnLevel
	} else {
		level = logrus.ErrorLevel
	}
	c.log.WithField("data", dl).Log(level, "apps")
}
