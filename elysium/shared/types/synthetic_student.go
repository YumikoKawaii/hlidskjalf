package types

import "time"

type SyntheticStudent struct {
	Id         string    `db:"id" ctype:"String"`
	Name       string    `db:"name" ctype:"String"`
	Age        uint32    `db:"age" ctype:"UInt32"`
	Sex        string    `db:"sex" ctype:"String"`
	Major      string    `db:"major" ctype:"String"`
	Level      string    `db:"level" ctype:"String"`
	GPA        float32   `db:"gpa" ctype:"Float32"`
	Hobbies    []string  `db:"hobbies" ctype:"Array(String)"`
	Country    string    `db:"country" ctype:"String"`
	Province   string    `db:"province" ctype:"String"`
	ActionTime time.Time `db:"action_time" ctype:"DateTime64(3)"`
}

func (SyntheticStudent) ToClickhouseTableName() string {
	return "synthetic_students"
}

func (SyntheticStudent) GetOrderKey() string {
	return "id"
}

func (SyntheticStudent) GetPartitionKey() string {
	return ""
}

func (SyntheticStudent) GetShardingKey() string {
	return "murmurHash3_64(id)"
}

func (SyntheticStudent) GetClickhouseEngine() string {
	return "ReplicatedReplacingMergeTree"
}

func (SyntheticStudent) GetClickhouseTTL() string {
	return "toDate(action_time) + toIntervalDay(10)"
}

func (SyntheticStudent) ToMySQLTableName() string {
	return "synthetic_students"
}

func (SyntheticStudent) ToPostgresTableName() string {
	return "synthetic_students"
}
