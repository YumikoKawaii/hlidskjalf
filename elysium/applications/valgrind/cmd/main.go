package main

import (
	"context"
	"elysium.com/applications/valgrind/pkg/mirgation"
	"elysium.com/shared/clickhouse"
	"elysium.com/shared/types"
)

func main() {

	ctx := context.Background()
	chCfg := clickhouse.LoadConfig()
	chClient := clickhouse.Initialize(ctx)

	migrator := mirgation.NewMigrator(chClient, chCfg)

	migrator.MigrateSchema(ctx, types.SyntheticStudent{})

}
