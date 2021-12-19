package logger

import (
	"context"
	"time"
)

// Finalize load from context and delete data context
func (d DataLogger) Finalize(ctx context.Context) {
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
