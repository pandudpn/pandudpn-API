package logger

import (
	"context"
	"time"

	"pandudpn/api/src/utils/config"

	"github.com/sirupsen/logrus"
)

// Finalize load from context and delete data context
func (d *DataLogger) Finalize(ctx context.Context) {
	value, ok := extract(ctx)
	if !ok {
		return
	}

	if i, ok := value.LoadAndDelete(_StatusCode); ok && i != nil {
		d.StatusCode = i.(int)
	}

	if i, ok := value.LoadAndDelete(_Response); ok && i != nil {
		d.Response = i
	}

	if i, ok := value.LoadAndDelete(_ThirdParties); ok && i != nil {
		d.ThirdParties = i.([]ThirdParty)
	}

	if i, ok := value.LoadAndDelete(_LogMessages); ok && i != nil {
		d.LogMessages = i.([]LogMessage)
	}

	if i, ok := value.LoadAndDelete(_ErrorMessage); ok && i != nil {
		d.ErrorMessage = i.(string)
	}

	d.ExecTime = time.Since(d.TimeStart).Seconds()
}

func (d *DataLogger) write() {
	var level logrus.Level

	if d.StatusCode >= 200 && d.StatusCode < 400 {
		level = logrus.InfoLevel
	} else if d.StatusCode >= 400 && d.StatusCode < 500 {
		level = logrus.WarnLevel
	} else {
		level = logrus.ErrorLevel
	}

	config.Logrus().WithField("data", d).Log(level, "apps")
}
