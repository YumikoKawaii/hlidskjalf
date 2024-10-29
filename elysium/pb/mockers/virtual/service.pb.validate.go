// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mockers/virtual/service.proto

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

// Validate checks the field values on VirtualRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *VirtualRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// VirtualRequestValidationError is the validation error returned by
// VirtualRequest.Validate if the designated constraints aren't met.
type VirtualRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VirtualRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VirtualRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VirtualRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VirtualRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VirtualRequestValidationError) ErrorName() string { return "VirtualRequestValidationError" }

// Error satisfies the builtin error interface
func (e VirtualRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVirtualRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VirtualRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VirtualRequestValidationError{}

// Validate checks the field values on Quaternary with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Quaternary) Validate() error {
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

// QuaternaryValidationError is the validation error returned by
// Quaternary.Validate if the designated constraints aren't met.
type QuaternaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuaternaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuaternaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuaternaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuaternaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuaternaryValidationError) ErrorName() string { return "QuaternaryValidationError" }

// Error satisfies the builtin error interface
func (e QuaternaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuaternary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuaternaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuaternaryValidationError{}

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

// Validate checks the field values on VirtualResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *VirtualResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VirtualResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// VirtualResponseValidationError is the validation error returned by
// VirtualResponse.Validate if the designated constraints aren't met.
type VirtualResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VirtualResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VirtualResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VirtualResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VirtualResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VirtualResponseValidationError) ErrorName() string { return "VirtualResponseValidationError" }

// Error satisfies the builtin error interface
func (e VirtualResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVirtualResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VirtualResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VirtualResponseValidationError{}

// Validate checks the field values on VirtualResponse_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *VirtualResponse_Data) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetStats() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return VirtualResponse_DataValidationError{
					field:  fmt.Sprintf("Stats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// VirtualResponse_DataValidationError is the validation error returned by
// VirtualResponse_Data.Validate if the designated constraints aren't met.
type VirtualResponse_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VirtualResponse_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VirtualResponse_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VirtualResponse_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VirtualResponse_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VirtualResponse_DataValidationError) ErrorName() string {
	return "VirtualResponse_DataValidationError"
}

// Error satisfies the builtin error interface
func (e VirtualResponse_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVirtualResponse_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VirtualResponse_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VirtualResponse_DataValidationError{}