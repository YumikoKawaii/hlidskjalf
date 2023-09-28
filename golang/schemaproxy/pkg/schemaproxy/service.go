package schemaproxy

import (
	"YumikoKawaii/hlidskjalf/golang/common/schemaregistry"
	pb "YumikoKawaii/hlidskjalf/golang/protobuf/schemaproxy"
	"context"
	"errors"
	"google.golang.org/appengine/log"
)

type Service interface {
	RegisterSchema(string, *schemaregistry.SchemaRequest) (*schemaregistry.SchemaResponse, error)
	DeleteSubject(string) error
	CheckSchemaCompatibility(string, *schemaregistry.SchemaRequest, int) (bool, error)
	GetSubjectVersions(string) ([]uint32, error)
	GetSchemaBySubjectVersion(*schemaregistry.SubjectVersion) (*schemaregistry.SchemaResponse, error)
	UpdateCompatibility(*pb.UpdateCompatibilityRequest) error
}

type serviceImpl struct {
	schemaRegistryClient schemaregistry.SchemaRegistryClient
}

func NewService() Service {
	return &serviceImpl{
		schemaRegistryClient: schemaregistry.NewSchemaRegistryClient(),
	}
}

func (s *serviceImpl) RegisterSchema(subject string, request *schemaregistry.SchemaRequest) (*schemaregistry.SchemaResponse, error) {

	isRegistered, schemaResponse, err := s.schemaRegistryClient.IsSchemaRegistered(subject, request)
	if err != nil {
		var subjectNotFoundErr schemaregistry.SubjectNotFoundError
		if errors.As(err, &subjectNotFoundErr) {
			log.Infof(context.Background(), "subject %s haven't registered yet", subject)
		} else {
			log.Errorf(context.Background(), "error while checking schema registry: %+v", err)
			return nil, err
		}
	}

	if isRegistered {
		return schemaResponse, nil
	}

	schemaResponse, err = s.schemaRegistryClient.RegisterSchema(subject, request)

	if err != nil {
		log.Errorf(context.Background(), "error while updating: %+v", err)
		return nil, err
	}

	return schemaResponse, err
}

func (s *serviceImpl) DeleteSubject(subject string) error {

	schemaIds, err := s.schemaRegistryClient.SoftDeleteSubject(subject)
	if err == nil {
		log.Infof(context.Background(), "soft delete schemas: %+v", schemaIds)
	}
	return err
}

func (s *serviceImpl) CheckSchemaCompatibility(subject string, request *schemaregistry.SchemaRequest, versionID int) (bool, error) {

	isCompatible, err := s.schemaRegistryClient.IsSchemaCompatible(subject, request, versionID)
	if err != nil {
		return false, err
	}
	return isCompatible, nil
}

func (s *serviceImpl) GetSubjectVersions(subject string) ([]uint32, error) {

	return s.schemaRegistryClient.GetVersions(subject)
}

func (s *serviceImpl) GetSchemaBySubjectVersion(version *schemaregistry.SubjectVersion) (*schemaregistry.SchemaResponse, error) {
	return s.schemaRegistryClient.GetSchemaBySubject(version)
}

func (s *serviceImpl) UpdateCompatibility(request *pb.UpdateCompatibilityRequest) error {
	return s.schemaRegistryClient.UpdateCompatibility(request)
}
