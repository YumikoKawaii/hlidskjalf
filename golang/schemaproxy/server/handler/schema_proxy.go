package handler

import (
	pb "YumikoKawaii/hlidskjalf/golang/protobuf/schemaproxy"
	"YumikoKawaii/hlidskjalf/golang/schemaproxy/pkg/schemaproxy"
	"context"
)

type ServiceServer struct {
	*pb.UnimplementedSchemaProxyServiceServer
	schemaService schemaproxy.Service
}

func (s *ServiceServer) RegisterSchema(ctx context.Context, request *pb.RegisterSchemaRequest) (*pb.RegistryResponse, error) {
	res, err := s.schemaService.RegisterSchema(request.Subject, convertToSchemaRequest(request))
	if err != nil {
		return nil, err
	}

	return &pb.RegistryResponse{
		Id:           int32(res.ID),
		Subject:      res.Subject,
		Version:      int32(res.Version),
		IsRegistered: true,
	}, nil
}

func (s *ServiceServer) DeleteSubject(ctx context.Context, request *pb.DeleteSubjectRequest) (*pb.DeleteSubjectResponse, error) {

	if err := s.schemaService.DeleteSubject(request.Subject); err != nil {
		return nil, err
	}

	return &pb.DeleteSubjectResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (s *ServiceServer) CheckSchemasCompatibility(ctx context.Context, request *pb.CheckSchemaCompatibilityRequest) (*pb.CheckSchemaCompatibilityResponse, error) {

	//s.schemaService.CheckSchemaCompatibility(request.Request.Subject, convertToSchemaRequest(request.Request), request.Request.)
	return nil, nil
}

func (s *ServiceServer) GetSubjectVersions(ctx context.Context, request *pb.GetSubjectVersionsRequest) (*pb.GetSubjectVersionsResponse, error) {

	versions, err := s.schemaService.GetSubjectVersions(request.Subject)
	if err != nil {
		return nil, err
	}

	return &pb.GetSubjectVersionsResponse{
		Versions: versions,
	}, nil
}

func (s *ServiceServer) GetSchemas(ctx context.Context, request *pb.GetSchemasRequest) (*pb.GetSchemasResponse, error) {
	return nil, nil
}
