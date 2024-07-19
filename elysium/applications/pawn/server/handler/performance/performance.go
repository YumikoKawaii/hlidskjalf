package performance

import (
	"context"
	"elysium.com/applications/pawn/pkg/students"
	pb "elysium.com/pb/pawn"
	"net/http"
)

type Server struct {
	studentsService students.Service
	pb.UnimplementedPerformanceServiceServer
}

func NewServer(service students.Service) *Server {
	return &Server{
		studentsService: service,
	}
}

func (s *Server) GetStudents(ctx context.Context, request pb.GetStudentsRequest) (*pb.GetStudentsResponse, error) {

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
