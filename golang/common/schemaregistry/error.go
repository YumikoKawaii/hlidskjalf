package schemaregistry

import (
	"fmt"
	"net/url"
)

const (
	subjectNotFoundCode = 40401
	versionNotFoundCode = 40402
	schemaNotFoundCode  = 40403

	incompatibleSchemaCode = 409

	invalidSchemaCode  = 42201
	invalidVersionCode = 42202

	errorInBackendDataStoreCode = 50001
	operationTimedOutCode       = 50002
	forwardRequestCode          = 50003
)

// ResourceError is being fired from all API calls when an error code is received.
type ResourceError struct {
	ErrorCode int    `json:"error_code"`
	Method    string `json:"method,omitempty"`
	URI       string `json:"uri,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (err ResourceError) Error() string {
	return fmt.Sprintf("client: (%s: %s) failed with error code %d %s",
		err.Method, err.URI, err.ErrorCode, err.Message)
}

func newResourceError(errCode int, uri, method, body string) ResourceError {
	unescapedURI, _ := url.QueryUnescape(uri)

	return ResourceError{
		ErrorCode: errCode,
		URI:       unescapedURI,
		Method:    method,
		Message:   body,
	}
}

var errRequired = func(field string) error {
	return fmt.Errorf("Client: %s is required", field)
}

type (
	SubjectNotFoundError struct {
		ResourceError
	}
	VersionNotFoundError struct {
		ResourceError
	}
	SchemaNotFoundError struct {
		ResourceError
	}

	IncompatibleSchemaError struct {
		ResourceError
	}

	InvalidSchemaError struct {
		ResourceError
	}
	InvalidVersionError struct {
		ResourceError
	}

	ErrorInBackendDataStore struct {
		ResourceError
	}

	OperationTimedOutError struct {
		ResourceError
	}
	ForwardRequestError struct {
		ResourceError
	}
)

func identifyResourceError(err ResourceError) error {
	switch err.ErrorCode {
	case subjectNotFoundCode:
		return SubjectNotFoundError{ResourceError: err}
	case versionNotFoundCode:
		return VersionNotFoundError{ResourceError: err}
	case schemaNotFoundCode:
		return SchemaNotFoundError{ResourceError: err}
	case incompatibleSchemaCode:
		return IncompatibleSchemaError{ResourceError: err}
	case invalidSchemaCode:
		return InvalidSchemaError{ResourceError: err}
	case invalidVersionCode:
		return InvalidVersionError{ResourceError: err}
	case errorInBackendDataStoreCode:
		return ErrorInBackendDataStore{ResourceError: err}
	case operationTimedOutCode:
		return OperationTimedOutError{ResourceError: err}
	case forwardRequestCode:
		return ForwardRequestError{ResourceError: err}
	default:
		return err
	}
}
