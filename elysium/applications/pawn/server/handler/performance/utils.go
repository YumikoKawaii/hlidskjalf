package performance

import (
	"elysium.com/applications/pawn/pkg/students"
	pb "elysium.com/pb/pawn"
)

func convertStudentToProtoStudent(student students.Student) *pb.Student {
	return &pb.Student{
		Id:       student.Id,
		FullName: student.Name,
		Age:      student.Age,
		Sex:      student.Sex,
		Major:    student.Major,
		Level:    student.Level,
		Gpa:      student.GPA,
		Hobbies:  student.Hobbies,
		Country:  student.Country,
		Province: student.Province,
	}
}
