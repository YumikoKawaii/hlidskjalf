package types

import (
	"encoding/json"
	"time"
)

type SyntheticStudent struct {
	Id         string    `db:"id" ctype:"String" json:"id,omitempty"`
	Name       string    `db:"name" ctype:"String" json:"name,omitempty"`
	Age        uint32    `db:"age" ctype:"UInt32" json:"age,omitempty"`
	Sex        string    `db:"sex" ctype:"String" json:"sex,omitempty"`
	Major      string    `db:"major" ctype:"String" json:"major,omitempty"`
	Level      string    `db:"level" ctype:"String" json:"level,omitempty"`
	GPA        float32   `db:"gpa" ctype:"Float32" json:"GPA,omitempty"`
	Hobbies    []string  `db:"hobbies" ctype:"Array(String)" json:"hobbies,omitempty"`
	Country    string    `db:"country" ctype:"String" json:"country,omitempty"`
	Province   string    `db:"province" ctype:"String" json:"province,omitempty"`
	ActionTime time.Time `db:"action_time" ctype:"DateTime64(3)" json:"actionTime"`
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

func (s SyntheticStudent) ToBytes() []byte {
	bytes, _ := json.Marshal(s)
	return bytes
}
