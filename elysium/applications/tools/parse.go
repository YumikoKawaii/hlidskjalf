package main

import (
	"elysium.com/shared/types"
	"encoding/base64"
	"fmt"
	"time"
)

func main() {

	student := types.SyntheticStudent{
		Id:         "id_1",
		Name:       "Name",
		Age:        1,
		Sex:        "",
		Major:      "",
		Level:      "",
		GPA:        0,
		Hobbies:    nil,
		Country:    "",
		Province:   "",
		ActionTime: time.Now().UTC(),
		Subject:    "synthetic_student",
	}

	bytes := student.ToBytes()
	encoded := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(encoded)

}
