package mirgation

import (
	"context"
	"elysium.com/shared/clickhouse"
	"elysium.com/shared/types"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
	"reflect"
	"sort"
	"strings"
)

type Migrator interface {
	// MigrateSchema use this func to migrate schema to clickhouse database
	MigrateSchema(ctx context.Context, prototype types.Prototype) error
}

type migratorImpl struct {
	chClient clickhouse.Client
	chCfg    *clickhouse.Config
}

func NewMigrator(client clickhouse.Client, config *clickhouse.Config) Migrator {
	return &migratorImpl{
		chClient: client,
		chCfg:    config,
	}
}

func (m *migratorImpl) MigrateSchema(ctx context.Context, prototype types.Prototype) error {
	descriptions, err := m.describeTable(ctx, prototype.ToClickhouseTableName())
	if err != nil {
		// parse error to detect if table is not initialized yet
		exception, ok := err.(*proto.Exception)
		if !ok {
			return err
		}
		if exception.Code != 60 {
			return err
		}
	}

	fields := reflect.TypeOf(prototype)
	protoes := make([]ColumnDescription, 0)
	for idx := 0; idx < fields.NumField(); idx++ {
		field := fields.Field(idx)
		protoes = append(
			protoes, ColumnDescription{
				Name: field.Tag.Get("db"),
				Type: field.Tag.Get("ctype"),
			},
		)
	}

	if len(descriptions) == 0 {
		if err := m.createLocalTable(ctx, prototype, protoes); err != nil {
			return err
		}

		if err := m.createDistributedTable(ctx, prototype); err != nil {
			return err
		}

		return nil
	}

	// sort descriptions
	sort.Slice(
		descriptions, func(i, j int) bool {
			return descriptions[i].Name < descriptions[j].Name
		},
	)

	// sort protoes
	sort.Slice(
		protoes, func(i, j int) bool {
			return protoes[i].Name < protoes[j].Name
		},
	)

	// get need to remove columns
	removeColumns := m.getColumnsNeedToRemove(descriptions, protoes)
	// remove columns
	if len(removeColumns) != 0 {
		if err := m.alterColumns(ctx, removeColumns, prototype, true); err != nil {
			return err
		}
	}

	// get new columns
	newColumns := m.getNewColumns(descriptions, protoes)
	// add new columns
	if len(newColumns) != 0 {
		if err := m.alterColumns(ctx, newColumns, prototype, false); err != nil {
			return err
		}
	}

	return nil
}

func (m *migratorImpl) describeTable(ctx context.Context, tableName string) ([]ColumnDescription, error) {

	table := fmt.Sprintf("%s.%s", m.chCfg.Database, tableName)
	stmt := fmt.Sprintf("desc %s", table)

	rows, err := m.chClient.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}

	descriptions := make([]ColumnDescription, 0)

	for rows.Next() {
		desc := ColumnDescription{}
		if err := rows.Scan(
			&desc.Name, &desc.Type, &desc.DefaultType, &desc.DefaultExpression, &desc.Comment, &desc.CodecExpression,
			&desc.TTLExpression,
		); err != nil {
			return nil, err
		}
		descriptions = append(descriptions, desc)
	}

	return descriptions, nil
}

func (m *migratorImpl) getNewColumns(desire []ColumnDescription, current []ColumnDescription) []ColumnDescription {

	desireMap := make(map[string]ColumnDescription)
	for _, column := range desire {
		desireMap[column.Name] = column
	}

	for _, column := range current {
		delete(desireMap, column.Name)
	}

	columns := make([]ColumnDescription, 0)
	for _, column := range desireMap {
		columns = append(columns, column)
	}
	return columns
}

func (m *migratorImpl) getColumnsNeedToRemove(desire []ColumnDescription, current []ColumnDescription) []ColumnDescription {
	currentMap := make(map[string]ColumnDescription)
	for _, column := range current {
		currentMap[column.Name] = column
	}

	for _, column := range desire {
		delete(currentMap, column.Name)
	}

	columns := make([]ColumnDescription, 0)
	for _, column := range currentMap {
		columns = append(columns, column)
	}
	return columns
}

func (m *migratorImpl) createLocalTable(ctx context.Context, prototype types.Prototype, columns []ColumnDescription) error {

	table := fmt.Sprintf("%s.%s", m.chCfg.Database, prototype.ToClickhouseTableName())

	partitionKey := ""
	if len(prototype.GetPartitionKey()) != 0 {
		partitionKey = fmt.Sprintf("partition by %s", prototype.GetPartitionKey())
	}

	ttl := ""
	if len(prototype.GetClickhouseTTL()) != 0 {
		ttl = fmt.Sprintf("ttl %s", prototype.GetClickhouseTTL())
	}

	columnStmt := m.generateColumnStmt(columns)

	stmt := fmt.Sprintf(
		""+
			"create table if not exists %s "+
			"on cluster %s "+
			"("+
			"	%s "+
			") "+
			"engine = %s "+
			"%s "+
			"order by %s "+
			"%s "+
			"settings index_granularity = 8192",
		table,
		m.chCfg.Cluster,
		columnStmt,
		prototype.GetClickhouseEngine(),
		partitionKey,
		prototype.GetOrderKey(),
		ttl,
	)

	return m.chClient.Exec(ctx, stmt)
}

func (m *migratorImpl) createDistributedTable(ctx context.Context, prototype types.Prototype) error {

	table := fmt.Sprintf("%s.%s", m.chCfg.Database, prototype.ToClickhouseTableName())
	// create distributed table
	distributedTable := clickhouse.GetDistributedTable(table)
	stmt := fmt.Sprintf(
		""+
			"create table if not exists %s "+
			"on cluster %s "+
			"as %s "+
			"engine = Distributed('%s', '%s', '%s', %s)",
		distributedTable,
		m.chCfg.Cluster,
		table,
		m.chCfg.Cluster,
		m.chCfg.Database,
		prototype.ToClickhouseTableName(),
		prototype.GetShardingKey(),
	)

	return m.chClient.Exec(ctx, stmt)
}

func (m *migratorImpl) dropTable(ctx context.Context, table string) error {
	table = fmt.Sprintf("%s.%s", m.chCfg.Database, table)
	stmt := fmt.Sprintf(
		""+
			"drop table if exists %s "+
			"on cluster %s",
		table,
		m.chCfg.Cluster,
	)
	return m.chClient.Exec(ctx, stmt)
}

func (m *migratorImpl) generateColumnStmt(columns []ColumnDescription) string {

	stmts := make([]string, 0)
	for _, column := range columns {
		stmts = append(
			stmts, fmt.Sprintf(
				"`%s` %s", column.Name, column.Type,
			),
		)
	}
	return strings.Join(stmts, ",")
}

func (m *migratorImpl) alterColumns(ctx context.Context, columns []ColumnDescription, prototype types.Prototype, isDrop bool) error {

	if err := m.dropTable(ctx, clickhouse.GetDistributedTable(prototype.ToClickhouseTableName())); err != nil {
		return err
	}

	table := fmt.Sprintf("%s.%s", m.chCfg.Database, prototype.ToClickhouseTableName())

	alterAction := m.generateAddColumnsStmt(columns)
	if isDrop {
		alterAction = m.generateDropColumnsStmt(columns)
	}

	// add new columns
	stmt := fmt.Sprintf(
		""+
			"alter table %s "+
			"on cluster %s "+
			"%s",
		table,
		m.chCfg.Cluster,
		alterAction,
	)
	if err := m.chClient.Exec(ctx, stmt); err != nil {
		return err
	}

	// recreate distributed table
	return m.createDistributedTable(ctx, prototype)
}

func (m *migratorImpl) generateAddColumnsStmt(columns []ColumnDescription) string {
	stmts := make([]string, 0)

	for _, column := range columns {
		stmts = append(
			stmts, fmt.Sprintf(
				"add column if not exists %s %s", column.Name,
				column.Type,
			),
		)
	}

	return strings.Join(stmts, ",")
}

func (m *migratorImpl) generateDropColumnsStmt(columns []ColumnDescription) string {
	stmts := make([]string, 0)

	for _, column := range columns {
		stmts = append(stmts, fmt.Sprintf("drop column if exists %s", column.Name))
	}

	return strings.Join(stmts, ",")
}
