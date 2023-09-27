package schemaregistry

import (
	"YumikoKawaii/hlidskjalf/golang/common/httpclient"
	pb "YumikoKawaii/hlidskjalf/golang/protobuf/schemaproxy"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/xerrors"
	"net/http"
)

const (
	SchemaLatestVersion = "latest"

	schemaAPIVersion      = "v1"
	contentTypeSchemaJSON = "application/vnd.schemaregistry." + schemaAPIVersion + "+json"

	baseURL = "hehe"
)

var (
	acceptHeaderValues = []string{contentTypeSchemaJSON, "application/vnd.schemaregistry+json", "application/json"}
)

type SchemaResponse struct {
	Subject    string       `json:"subject"`
	ID         int          `json:"id,omitempty"`
	Version    int          `json:"version"`
	Schema     string       `json:"schema"`
	SchemaType string       `json:"schemaType"`
	References []*Reference `json:"references"`
}

type SchemaRequest struct {
	Schema     string       `json:"schema"`
	SchemaType string       `json:"schemaType,omitempty"`
	References []*Reference `json:"references,omitempty"`
}

type Reference struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Version int    `json:"version"`
}

type subjectVersionOnlyJSON struct {
	Subject string `json:"subject"`
	Version int    `json:"version"`
}

type idOnlyJSON struct {
	ID int `json:"id"`
}

type isCompatibleJSON struct {
	IsCompatible bool `json:"is_compatible"`
}

type SchemaRegistryClient interface {
	GetSubjects() ([]string, error)
	GetVersions(string) ([]int, error)
	SoftDeleteSubject(string) ([]int, error)
	HardDeleteSubject(string) ([]int, error)
	IsSchemaRegistered(string, *SchemaRequest) (bool, *SchemaResponse, error)
	GetSchemaByID(int) (*SchemaResponse, error)
	GetSubjectVersionByID(int) (*string, *int, error)
	GetSchemaBySubject(string, int) (*SchemaResponse, error)
	GetLatestSchema(string) (*SchemaResponse, error)
	IsSchemaCompatible(string, *SchemaRequest, int) (bool, error)
	IsLatestSchemaCompatible(string, *SchemaRequest) (bool, error)
	UpdateCompatibility(context.Context, *pb.UpdateCompatibilityRequest) error
}

type clientImpl struct {
	httpHandler httpclient.HttpHandler
}

func NewSchemaRegistryClient() SchemaRegistryClient {

	handler := httpclient.NewHttpHandler(baseURL)

	return clientImpl{
		httpHandler: handler,
	}
}

func (c clientImpl) GetSubjects() ([]string, error) {
	response, err := c.httpHandler.Do(http.MethodGet, "/subjects", "", nil, acceptHeaderValues...)
	if err != nil {
		return nil, err
	}
	var subjects []string
	err = c.httpHandler.ReadJson(response, &subjects)
	return subjects, err
}

func (c clientImpl) GetVersions(subject string) ([]int, error) {
	if len(subject) == 0 {
		return nil, errors.New("required subject")
	}

	path := fmt.Sprintf("subjects/%s/versions", subject)
	response, err := c.httpHandler.Do(http.MethodGet, path, "", nil, acceptHeaderValues...)
	if err != nil {
		return nil, err
	}
	var versions []int
	err = c.httpHandler.ReadJson(response, versions)
	return versions, err
}

func (c clientImpl) SoftDeleteSubject(subject string) ([]int, error) {
	return c.deleteSubject(subject, false)
}

func (c clientImpl) HardDeleteSubject(subject string) ([]int, error) {
	return c.deleteSubject(subject, true)
}

func (c clientImpl) deleteSubject(subject string, isPermanent bool) ([]int, error) {
	if len(subject) == 0 {
		return nil, errors.New("required subject")
	}

	path := fmt.Sprintf("/subjects/%s", subject)
	if isPermanent {
		path += "?permanent=true"
	}
	response, err := c.httpHandler.Do(http.MethodDelete, path, "", nil, acceptHeaderValues...)
	if err != nil {
		return nil, err
	}
	var versions []int
	err = c.httpHandler.ReadJson(response, versions)
	return versions, err
}

func (c clientImpl) IsSchemaRegistered(subject string, request *SchemaRequest) (bool, *SchemaResponse, error) {
	var schemaResponse *SchemaResponse
	send, err := json.Marshal(request)
	if err != nil {
		return false, schemaResponse, err
	}

	path := fmt.Sprintf("subjects/%s", subject)
	response, err := c.httpHandler.Do(http.MethodPost, path, "", send, acceptHeaderValues...)

	if response.StatusCode == schemaNotFoundCode || err != nil {
		return false, schemaResponse, err
	}

	err = c.httpHandler.ReadJson(response, &schemaResponse)
	return true, schemaResponse, err
}

func (c clientImpl) GetSchemaByID(schemaID int) (*SchemaResponse, error) {
	path := fmt.Sprintf("/schemas/ids/%d", schemaID)
	response, err := c.httpHandler.Do(http.MethodGet, path, "", nil, acceptHeaderValues...)
	if err != nil {
		return nil, err
	}

	var schemaResponse *SchemaResponse
	if err = c.httpHandler.ReadJson(response, schemaResponse); err != nil {
		return nil, err
	}
	return schemaResponse, nil
}

func (c clientImpl) GetSubjectVersionByID(schemaID int) (*string, *int, error) {
	path := fmt.Sprintf("/schemas/ids/%d/versions", schemaID)

	response, err := c.httpHandler.Do(http.MethodGet, path, "", nil, acceptHeaderValues...)
	if err != nil {
		return nil, nil, err
	}

	var res []*subjectVersionOnlyJSON
	if err := c.httpHandler.ReadJson(response, res); err != nil {
		return nil, nil, err
	}

	return &res[0].Subject, &res[0].Version, nil
}

func (c clientImpl) getSubjectSchemaAtVersion(subject string, versionID interface{}) (*SchemaResponse, error) {
	if subject == "" {
		err := errRequired("subject")
		return nil, err
	}
	path := fmt.Sprintf("/subjects/%s/versions/%v", subject, versionID)
	resp, err := c.httpHandler.Do(http.MethodGet, path, "", nil, acceptHeaderValues...)
	if err != nil {
		return nil, err
	}

	var schema *SchemaResponse

	err = c.httpHandler.ReadJson(resp, schema)
	return schema, err
}

func (c clientImpl) GetSchemaBySubject(subject string, versionID int) (*SchemaResponse, error) {
	return c.getSubjectSchemaAtVersion(subject, versionID)
}

func (c clientImpl) GetLatestSchema(subject string) (*SchemaResponse, error) {
	return c.getSubjectSchemaAtVersion(subject, SchemaLatestVersion)
}

func (c clientImpl) isSchemaCompatibleAtVersion(subject string, schema *SchemaRequest, versionID interface{}) (bool, error) {
	if subject == "" {
		err := errRequired("subject")
		return false, err
	}
	if schema == nil {
		err := errRequired("schema")
		return false, err
	}

	send, err := json.Marshal(schema)
	if err != nil {
		return false, err
	}

	path := fmt.Sprintf("/compatibility/subjects/%s/versions/%v", subject, versionID)
	resp, err := c.httpHandler.Do(http.MethodPost, path, httpclient.ContentTypeJSON, send, acceptHeaderValues...)
	if err != nil {
		return false, err
	}

	var res isCompatibleJSON
	err = c.httpHandler.ReadJson(resp, &res)
	return res.IsCompatible, err
}

func (c clientImpl) IsSchemaCompatible(subject string, schema *SchemaRequest, versionID int) (bool, error) {
	return c.isSchemaCompatibleAtVersion(subject, schema, versionID)
}

func (c clientImpl) IsLatestSchemaCompatible(subject string, schema *SchemaRequest) (bool, error) {
	return c.isSchemaCompatibleAtVersion(subject, schema, SchemaLatestVersion)
}

func (c clientImpl) UpdateCompatibility(ctx context.Context, request *pb.UpdateCompatibilityRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return xerrors.Errorf("error when client update compatibility %w", err)
	}

	_, err = c.httpHandler.Do(http.MethodPut, "/config", "application/vnd.schemaregistry.v1+json", body, acceptHeaderValues...)
	return err
}
