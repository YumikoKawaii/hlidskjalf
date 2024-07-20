package students

import (
	"context"
	"elysium.com/shared/clickhouse"
	"fmt"
)

type Repository interface {
	GetStudents(ctx context.Context) ([]Student, error)
}

type repoImpl struct {
	chClient clickhouse.Client
	chCfg    *clickhouse.Config
}

func NewRepository(client clickhouse.Client, cfg *clickhouse.Config) Repository {
	return &repoImpl{
		chClient: client,
		chCfg:    cfg,
	}
}

func (r *repoImpl) GetStudents(ctx context.Context) ([]Student, error) {

	table := fmt.Sprintf("%s.%s", r.chCfg.Database, SyntheticStudentProfilesTable)

	query := fmt.Sprintf(
		""+
			"select * "+
			"from %s",
		table,
	)

	rows, err := r.chClient.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	students := make([]Student, 0)

	for rows.Next() {
		student := Student{}
		if err := rows.Scan(
			&student.Id, &student.Name, &student.Age, &student.Sex, &student.Major, &student.Level, &student.GPA,
			&student.Hobbies, &student.Country, &student.Province,
		); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
