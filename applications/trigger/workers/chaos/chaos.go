package chaos

import (
	"context"
	"math/rand/v2"
	"time"

	"github.com/YumikoKawaii/shared/logger"
)

type Config struct {
	Interval int `json:"interval" mapstructure:"interval" yaml:"interval"`
}

type TriggerFunction struct {
	Name     string
	Function func(ctx context.Context) error
}

type Chaos interface {
	Register(functions ...TriggerFunction)
	Start(ctx context.Context)
}

type chaos struct {
	triggerFunctions []TriggerFunction
	cfg              *Config
}

func Initialize(cfg *Config) Chaos {
	return &chaos{
		cfg: cfg,
	}
}

func (c *chaos) Register(functions ...TriggerFunction) {
	c.triggerFunctions = append(c.triggerFunctions, functions...)
}

func (c *chaos) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(c.cfg.Interval) * time.Second)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				idx := rand.IntN(len(c.triggerFunctions))
				tf := c.triggerFunctions[idx]
				if err := tf.Function(ctx); err != nil {
					logger.WithFields(logger.Fields{"function": tf.Name, "error": err}).Error("[とりがー] - とりがーしっぱい")
				} else {
					logger.WithFields(logger.Fields{"function": tf.Name}).Info("[とりがー] - とりがーせいこう")
				}
			}
		}
	}()
}
