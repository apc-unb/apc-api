package prometheus

import (
	"time"

	"github.com/apc-unb/apc-api/web/metrics"
)

func RecordUpTime() {

	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				metrics.Uptime.Inc()
			}
		}
	}()

}
