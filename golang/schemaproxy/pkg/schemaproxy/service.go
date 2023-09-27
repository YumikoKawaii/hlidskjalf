package schemaproxy

import (
	"YumikoKawaii/hlidskjalf/golang/common/schemaregistry"
	pb "YumikoKawaii/hlidskjalf/golang/protobuf/schemaproxy"
)

type Service interface {
	UpdateSchema(*pb.UpdateSchemaRequest) (*pb.RegistryResponse, error)
	DeleteSubject(*pb.DeleteSubjectRequest) (*pb.DeleteSubjectResponse, error)
	CheckSchemasCompatibility(*pb.CheckSchemasCompatibilityRequest) (*pb.CheckSchemasCompatibilityResponse, error)
	GetSubjectVersions(*pb.GetSubjectVersionsRequest) (*pb.GetSubjectVersionsResponse, error)
	GetSchemas(*pb.GetSchemasRequest) (*pb.GetSchemasResponse, error)
	UpdateCompatibility(*pb.UpdateCompatibilityRequest) (*pb.UpdateCompatibilityResponse, error)
}

type serviceImpl struct {
	schemaRegistryClient schemaregistry.SchemaRegistryClient
}

func NewService() Service {
	return &serviceImpl{
		schemaRegistryClient: schemaregistry.NewSchemaRegistryClient(),
	}
}

func (s *serviceImpl) UpdateSchema(request *pb.UpdateSchemaRequest) (*pb.RegistryResponse, error) {
	return nil, nil
}

func (s *serviceImpl) DeleteSubject(request *pb.DeleteSubjectRequest) (*pb.DeleteSubjectResponse, error) {
	return nil, nil
}

func (s *serviceImpl) CheckSchemasCompatibility(request *pb.CheckSchemasCompatibilityRequest) (*pb.CheckSchemasCompatibilityResponse, error) {
	return nil, nil
}

func (s *serviceImpl) GetSubjectVersions(request *pb.GetSubjectVersionsRequest) (*pb.GetSubjectVersionsResponse, error) {
	return nil, nil
}

func (s *serviceImpl) GetSchemas(request *pb.GetSchemasRequest) (*pb.GetSchemasResponse, error) {
	return nil, nil
}

func (s *serviceImpl) UpdateCompatibility(request *pb.UpdateCompatibilityRequest) (*pb.UpdateCompatibilityResponse, error) {
	return nil, nil
}
