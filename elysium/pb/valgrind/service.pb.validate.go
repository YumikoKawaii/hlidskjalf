// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: valgrind/service.proto

package api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on EntryRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EntryRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetSubject()) < 1 {
		return EntryRequestValidationError{
			field:  "Subject",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetPayload()) < 1 {
		return EntryRequestValidationError{
			field:  "Payload",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// EntryRequestValidationError is the validation error returned by
// EntryRequest.Validate if the designated constraints aren't met.
type EntryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntryRequestValidationError) ErrorName() string { return "EntryRequestValidationError" }

// Error satisfies the builtin error interface
func (e EntryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntryRequestValidationError{}

// Validate checks the field values on EntryResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EntryResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	return nil
}

// EntryResponseValidationError is the validation error returned by
// EntryResponse.Validate if the designated constraints aren't met.
type EntryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntryResponseValidationError) ErrorName() string { return "EntryResponseValidationError" }

// Error satisfies the builtin error interface
func (e EntryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntryResponseValidationError{}