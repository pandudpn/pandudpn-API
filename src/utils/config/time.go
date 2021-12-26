package config

import "time"

const tz = "Asia/Jakarta"

func Timezone() *time.Location {
	timezone, err := time.LoadLocation(tz)
	if err != nil {
		panic(err)
	}
	
	return timezone
}
