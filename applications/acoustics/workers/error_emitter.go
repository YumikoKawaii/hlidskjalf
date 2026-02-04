package workers

import (
	"context"
	"time"

	"github.com/YumikoKawaii/shared/logger"
)

type ErrorEmitter struct {
	Interval int
}

func (e *ErrorEmitter) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(e.Interval) * time.Second)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				logger.Error("[あこーすてぃくす] - えらーはっせい: えらーえみったーからのていきてきなえらーろぐ")
			}
		}
	}()
}
