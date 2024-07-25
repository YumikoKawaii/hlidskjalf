package clickhouse

import (
	"context"
	"fmt"
	client "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"golang.org/x/xerrors"
	"net"
	"strings"
	"time"
)

type Client interface {
	Query(ctx context.Context, query string, args ...any) (driver.Rows, error)
	Exec(ctx context.Context, query string, args ...any) error
}

type clientImpl struct {
	conn client.Conn
}

func Initialize(ctx context.Context, cfg Config) Client {

	conn, err := client.Open(
		&client.Options{
			Addr: strings.Split(cfg.Addresses, ";"),
			Auth: client.Auth{
				Database: cfg.Database,
				Username: cfg.Username,
				Password: cfg.Password,
			},
			DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
				//dialCount++
				var d net.Dialer
				return d.DialContext(ctx, "tcp", addr)
			},
			Debug: true,
			Debugf: func(format string, v ...any) {
				fmt.Printf(format+"\n", v...)
			},
			Settings: client.Settings{
				"max_execution_time": 60,
			},
			Compression: &client.Compression{
				Method: client.CompressionLZ4,
			},
			DialTimeout:          time.Second * 30,
			MaxOpenConns:         5,
			MaxIdleConns:         5,
			ConnMaxLifetime:      time.Duration(10) * time.Minute,
			ConnOpenStrategy:     client.ConnOpenInOrder,
			BlockBufferSize:      10,
			MaxCompressionBuffer: 10240,
		},
	)

	if err != nil {
		panic(xerrors.Errorf("error init clickhouse client: %w", err))
	}

	if err := conn.Ping(ctx); err != nil {
		panic(xerrors.Errorf("error while ping clickhouse: %w", err))
	}

	return &clientImpl{
		conn: conn,
	}
}

func (c *clientImpl) Query(ctx context.Context, query string, args ...any) (driver.Rows, error) {
	return c.conn.Query(ctx, query, args)
}

func (c *clientImpl) Exec(ctx context.Context, query string, args ...any) error {
	return c.conn.Exec(ctx, query, args)
}
