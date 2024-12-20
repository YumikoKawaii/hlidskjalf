// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mockers/echo/service.proto

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on EchoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EchoRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// EchoRequestValidationError is the validation error returned by
// EchoRequest.Validate if the designated constraints aren't met.
type EchoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EchoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EchoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EchoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EchoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EchoRequestValidationError) ErrorName() string { return "EchoRequestValidationError" }

// Error satisfies the builtin error interface
func (e EchoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEchoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EchoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EchoRequestValidationError{}

// Validate checks the field values on Stat with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Stat) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for First

	// no validation rules for Second

	// no validation rules for Third

	// no validation rules for Fourth

	// no validation rules for Fifth

	// no validation rules for Sixth

	// no validation rules for Seventh

	// no validation rules for Eighth

	// no validation rules for Ninth

	// no validation rules for Tenth

	return nil
}

// StatValidationError is the validation error returned by Stat.Validate if the
// designated constraints aren't met.
type StatValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatValidationError) ErrorName() string { return "StatValidationError" }

// Error satisfies the builtin error interface
func (e StatValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStat.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatValidationError{}

// Validate checks the field values on EchoResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EchoResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EchoResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// EchoResponseValidationError is the validation error returned by
// EchoResponse.Validate if the designated constraints aren't met.
type EchoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EchoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EchoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EchoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EchoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EchoResponseValidationError) ErrorName() string { return "EchoResponseValidationError" }

// Error satisfies the builtin error interface
func (e EchoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEchoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EchoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EchoResponseValidationError{}

// Validate checks the field values on EchoResponse_Data with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *EchoResponse_Data) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetStats() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EchoResponse_DataValidationError{
					field:  fmt.Sprintf("Stats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// EchoResponse_DataValidationError is the validation error returned by
// EchoResponse_Data.Validate if the designated constraints aren't met.
type EchoResponse_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EchoResponse_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EchoResponse_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EchoResponse_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EchoResponse_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EchoResponse_DataValidationError) ErrorName() string {
	return "EchoResponse_DataValidationError"
}

// Error satisfies the builtin error interface
func (e EchoResponse_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEchoResponse_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EchoResponse_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EchoResponse_DataValidationError{}
