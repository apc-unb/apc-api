package prometheus

import (
	"time"

	"github.com/VerasThiago/plataforma-apc/metrics"
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
