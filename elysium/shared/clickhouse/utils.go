package clickhouse

import (
	"fmt"
	"strings"
)

const (
	distributedPostfix = "_distributed"
)

func GetDistributedTable(table string) string {
	return fmt.Sprintf("%s%s", table, distributedPostfix)
}

func GetLocalTable(table string) string {
	return strings.Replace(table, distributedPostfix, "", -1)
}
