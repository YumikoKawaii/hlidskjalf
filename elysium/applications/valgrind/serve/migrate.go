package serve

import (
	"context"
	"elysium.com/applications/valgrind/config"
	"elysium.com/applications/valgrind/pkg/mirgation"
	"elysium.com/shared/clickhouse"
	"elysium.com/shared/types"
	"google.golang.org/appengine/log"
)

func Migrate(cfg *config.Config) {

	ctx := context.Background()

	chClient := clickhouse.Initialize(ctx, cfg.ClickhouseCfg)
	migrator := mirgation.NewMigrator(chClient, &cfg.ClickhouseCfg)

	schemas := []types.Prototype{
		types.SyntheticStudent{},
	}

	for _, schema := range schemas {
		if err := migrator.MigrateSchema(ctx, schema); err != nil {
			log.Errorf(ctx, "error migrating schema %s: %s", schema.ToClickhouseTableName(), err.Error())
		}
	}
}
