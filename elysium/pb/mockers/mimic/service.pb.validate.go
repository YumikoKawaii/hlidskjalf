// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mockers/mimic/service.proto

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

// Validate checks the field values on MimicRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MimicRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MimicRequestValidationError is the validation error returned by
// MimicRequest.Validate if the designated constraints aren't met.
type MimicRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MimicRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MimicRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MimicRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MimicRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MimicRequestValidationError) ErrorName() string { return "MimicRequestValidationError" }

// Error satisfies the builtin error interface
func (e MimicRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMimicRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MimicRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MimicRequestValidationError{}

// Validate checks the field values on Tertiary with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Tertiary) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for First

	// no validation rules for Second

	// no validation rules for Third

	// no validation rules for Fourth

	// no validation rules for Fifth

	return nil
}

// TertiaryValidationError is the validation error returned by
// Tertiary.Validate if the designated constraints aren't met.
type TertiaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TertiaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TertiaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TertiaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TertiaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TertiaryValidationError) ErrorName() string { return "TertiaryValidationError" }

// Error satisfies the builtin error interface
func (e TertiaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTertiary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TertiaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TertiaryValidationError{}

// Validate checks the field values on Secondary with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Secondary) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for First

	// no validation rules for Second

	// no validation rules for Third

	// no validation rules for Fourth

	// no validation rules for Fifth

	for idx, item := range m.GetTertiaries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecondaryValidationError{
					field:  fmt.Sprintf("Tertiaries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SecondaryValidationError is the validation error returned by
// Secondary.Validate if the designated constraints aren't met.
type SecondaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecondaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecondaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecondaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecondaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecondaryValidationError) ErrorName() string { return "SecondaryValidationError" }

// Error satisfies the builtin error interface
func (e SecondaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecondary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecondaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecondaryValidationError{}

// Validate checks the field values on Primary with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Primary) Validate() error {
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

	for idx, item := range m.GetSecondaries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PrimaryValidationError{
					field:  fmt.Sprintf("Secondaries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PrimaryValidationError is the validation error returned by Primary.Validate
// if the designated constraints aren't met.
type PrimaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrimaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrimaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrimaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrimaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrimaryValidationError) ErrorName() string { return "PrimaryValidationError" }

// Error satisfies the builtin error interface
func (e PrimaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrimary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrimaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrimaryValidationError{}

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

	for idx, item := range m.GetPrimaries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StatValidationError{
					field:  fmt.Sprintf("Primaries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

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

// Validate checks the field values on MimicResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MimicResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MimicResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MimicResponseValidationError is the validation error returned by
// MimicResponse.Validate if the designated constraints aren't met.
type MimicResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MimicResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MimicResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MimicResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MimicResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MimicResponseValidationError) ErrorName() string { return "MimicResponseValidationError" }

// Error satisfies the builtin error interface
func (e MimicResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMimicResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MimicResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MimicResponseValidationError{}

// Validate checks the field values on MimicResponse_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MimicResponse_Data) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetStats() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MimicResponse_DataValidationError{
					field:  fmt.Sprintf("Stats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MimicResponse_DataValidationError is the validation error returned by
// MimicResponse_Data.Validate if the designated constraints aren't met.
type MimicResponse_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MimicResponse_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MimicResponse_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MimicResponse_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MimicResponse_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MimicResponse_DataValidationError) ErrorName() string {
	return "MimicResponse_DataValidationError"
}

// Error satisfies the builtin error interface
func (e MimicResponse_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMimicResponse_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MimicResponse_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MimicResponse_DataValidationError{}
