package performance

import (
	"context"
	"elysium.com/applications/pawn/pkg/students"
	pb "elysium.com/pb/pawn"
	"net/http"
)

type ServiceServer struct {
	pb.UnimplementedPerformanceServiceServer
	studentsService students.Service
}

func NewServiceServer(service students.Service) *ServiceServer {
	return &ServiceServer{
		studentsService: service,
	}
}

func (s *ServiceServer) GetStudents(ctx context.Context, request *pb.GetStudentsRequest) (*pb.GetStudentsResponse, error) {

	studentsData, err := s.studentsService.GetStudents(ctx)
	if err != nil {
		return nil, err
	}

	studentProtoes := make([]*pb.Student, 0)
	for _, student := range studentsData {
		studentProtoes = append(studentProtoes, convertStudentToProtoStudent(student))
	}

	return &pb.GetStudentsResponse{
		Code:     http.StatusOK,
		Message:  "success",
		Students: studentProtoes,
	}, nil
}
