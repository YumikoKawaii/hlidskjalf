package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type SyntheticStudent struct {
	Id         string    `db:"id" ctype:"String" json:"id,omitempty"`
	Name       string    `db:"name" ctype:"String" json:"name,omitempty"`
	Age        uint32    `db:"age" ctype:"UInt32" json:"age,omitempty"`
	Sex        string    `db:"sex" ctype:"String" json:"sex,omitempty"`
	Major      string    `db:"major" ctype:"String" json:"major,omitempty"`
	Level      string    `db:"level" ctype:"String" json:"level,omitempty"`
	GPA        float32   `db:"gpa" ctype:"Float32" json:"gpa,omitempty"`
	Hobbies    []string  `db:"hobbies" ctype:"Array(String)" json:"hobbies,omitempty"`
	Country    string    `db:"country" ctype:"String" json:"country,omitempty"`
	Province   string    `db:"province" ctype:"String" json:"province,omitempty"`
	ActionTime time.Time `db:"action_time" ctype:"DateTime64(3)" json:"actionTime"`

	Subject string `db:"-" ctype:"-" json:"subject"`
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

func (s SyntheticStudent) ToInsertValue() string {

	hobbies := make([]string, 0)
	for _, hobby := range s.Hobbies {
		hobbies = append(hobbies, fmt.Sprintf("'%s'", hobby))
	}
	hobbiesValue := "[]"
	if len(hobbies) != 0 {
		hobbiesValue = fmt.Sprintf("[%s]", strings.Join(hobbies, ","))
	}

	return fmt.Sprintf(
		""+
			"("+
			"%[1]s, "+
			"%[2]s, "+
			"%[3]d, "+
			"%[4]s, "+
			"%[5]s, "+
			"%[6]s, "+
			"%[7]f, "+
			"%[8]s, "+
			"%[9]s, "+
			"%[10]s, "+
			"%[11]s)",
		s.Id,                  //1
		s.Name,                //2
		s.Age,                 //3
		s.Sex,                 //4
		s.Major,               //5
		s.Level,               //6
		s.GPA,                 //6
		hobbiesValue,          //7
		s.Country,             //8
		s.Province,            //9
		s.ActionTime.String(), //10
	)
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

func (SyntheticStudent) Analyze(data []byte) (Prototype, error) {
	student := SyntheticStudent{}
	err := json.Unmarshal(data, &student)
	return student, err
}

func (SyntheticStudent) GetSubject() string {
	return "synthetic_student"
}
