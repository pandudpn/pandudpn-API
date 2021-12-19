package config

import "time"

const jakarta = "Asia/Jakarta"

func Timezone() *time.Location {
	timezone, err := time.LoadLocation(jakarta)
	if err != nil {
		panic(err)
	}
	
	return timezone
}
