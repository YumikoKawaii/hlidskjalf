package handlers

import (
	"context"
	"crypto/sha256"
	"sync"
	"sync/atomic"
	"time"

	"github.com/YumikoKawaii/hlidskjalf/applications/consumer/config"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

// Consumer drains a Watermill subscriber channel with a pool of N worker
// goroutines. The pool size (cfg.Concurrency) — NOT the channel buffer size —
// is what determines how many messages are processed in parallel, and therefore
// how much CPU the pod actually uses.
type Consumer struct {
	cfg       *config.KafkaConfig
	processed atomic.Uint64
}

func Initialize(cfg *config.KafkaConfig) *Consumer {
	return &Consumer{cfg: cfg}
}

// Run subscribes and fans the message channel out to cfg.Concurrency workers.
//
// The single subscriber channel below is exactly the point where Watermill's
// per-partition goroutines re-converge. With one consuming goroutine, all
// partitions collapse into one ack-gated lane — the steady-0.3-core behaviour.
// The worker pool restores parallelism above that channel.
func (c *Consumer) Run(ctx context.Context, sub *kafka.Subscriber) error {
	messages, err := sub.Subscribe(ctx, c.cfg.Topic)
	if err != nil {
		return err
	}

	concurrency := c.cfg.Concurrency
	if concurrency < 1 {
		concurrency = 1
	}

	logger.Infof("[こんしゅーまー] - わーかーきどう: ど=%d, ばっふぁ=%d, あふたーわーく=%v",
		concurrency, c.cfg.ChannelBufferSize, c.cfg.AckAfterWork)

	c.startReporter(ctx)

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for msg := range messages {
				c.handle(msg)
			}
		}(i)
	}

	wg.Wait()
	return nil
}

func (c *Consumer) handle(msg *message.Message) {
	// AckAfterWork=false: ack first, then process. The router advances
	// immediately => at-most-once (lost on crash). Smallest change, biggest
	// apparent throughput jump, but unsafe for at-least-once.
	if !c.cfg.AckAfterWork {
		msg.Ack()
		c.process(msg)
		c.processed.Add(1)
		return
	}

	// AckAfterWork=true: process, then ack => at-least-once. With a worker pool
	// this is still parallel (N messages in flight), and the offset is only
	// advanced once the work for that message has actually completed.
	c.process(msg)
	msg.Ack()
	c.processed.Add(1)
}

// process does real CPU-bound work per message so CPU usage is observable on a
// Grafana panel. It hashes the payload WorkIterations times in a tight loop.
//
// This is the key to the experiment: a sleeping goroutine consumes ZERO CPU, so
// the only way CONCURRENCY can visibly move the CPU graph is if each lane is
// genuinely on-CPU. With WorkIterations tuned so one message ~saturates one
// core for a few ms:
//   - CONCURRENCY=1  -> one lane on-CPU            -> ~1 core, regardless of buffer
//   - CONCURRENCY=12 -> twelve lanes on-CPU        -> pegs the pod up to its limit
//
// WorkDelayMs is kept as an optional I/O-bound tail (e.g. a downstream call), so
// both regimes can be demonstrated from the same binary.
func (c *Consumer) process(msg *message.Message) {
	iterations := c.cfg.WorkIterations
	digest := msg.Payload
	for i := 0; i < iterations; i++ {
		sum := sha256.Sum256(digest)
		digest = sum[:]
	}
	// Prevent the compiler from optimising the loop away.
	if len(digest) == 0 {
		logger.Debug("[こんしゅーまー] - くうだいじぇすと")
	}

	if d := c.cfg.WorkDelayMs; d > 0 {
		t := time.NewTimer(time.Duration(d) * time.Millisecond)
		defer t.Stop()
		select {
		case <-t.C:
		case <-msg.Context().Done():
		}
	}
}

// startReporter logs throughput once per second so the effect of changing
// Concurrency vs ChannelBufferSize is visible without external tooling.
func (c *Consumer) startReporter(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		var last uint64
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				now := c.processed.Load()
				logger.Infof("[こんしゅーまー] - すろーぷっと: %d msg/s (るいけい=%d)", now-last, now)
				last = now
			}
		}
	}()
}
