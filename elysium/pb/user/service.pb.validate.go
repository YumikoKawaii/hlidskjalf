// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/service.proto

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

// Validate checks the field values on UserInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UserInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Alias

	// no validation rules for Avatar

	// no validation rules for Introduction

	// no validation rules for Workplace

	// no validation rules for Hometown

	return nil
}

// UserInfoValidationError is the validation error returned by
// UserInfo.Validate if the designated constraints aren't met.
type UserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserInfoValidationError) ErrorName() string { return "UserInfoValidationError" }

// Error satisfies the builtin error interface
func (e UserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserInfoValidationError{}

// Validate checks the field values on UpsertUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpsertUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Alias

	// no validation rules for Avatar

	// no validation rules for Introduction

	// no validation rules for Workplace

	// no validation rules for Hometown

	return nil
}

// UpsertUserRequestValidationError is the validation error returned by
// UpsertUserRequest.Validate if the designated constraints aren't met.
type UpsertUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertUserRequestValidationError) ErrorName() string {
	return "UpsertUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpsertUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertUserRequestValidationError{}

// Validate checks the field values on UpsertUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpsertUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	// no validation rules for Id

	return nil
}

// UpsertUserResponseValidationError is the validation error returned by
// UpsertUserResponse.Validate if the designated constraints aren't met.
type UpsertUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertUserResponseValidationError) ErrorName() string {
	return "UpsertUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpsertUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertUserResponseValidationError{}

// Validate checks the field values on GetUsersInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetUsersInfoRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetIds()); l < 1 || l > 1000 {
		return GetUsersInfoRequestValidationError{
			field:  "Ids",
			reason: "value must contain between 1 and 1000 items, inclusive",
		}
	}

	_GetUsersInfoRequest_Ids_Unique := make(map[string]struct{}, len(m.GetIds()))

	for idx, item := range m.GetIds() {
		_, _ = idx, item

		if _, exists := _GetUsersInfoRequest_Ids_Unique[item]; exists {
			return GetUsersInfoRequestValidationError{
				field:  fmt.Sprintf("Ids[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_GetUsersInfoRequest_Ids_Unique[item] = struct{}{}
		}

		if utf8.RuneCountInString(item) != 36 {
			return GetUsersInfoRequestValidationError{
				field:  fmt.Sprintf("Ids[%v]", idx),
				reason: "value length must be 36 runes",
			}

		}

	}

	// no validation rules for Page

	// no validation rules for PageSize

	return nil
}

// GetUsersInfoRequestValidationError is the validation error returned by
// GetUsersInfoRequest.Validate if the designated constraints aren't met.
type GetUsersInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersInfoRequestValidationError) ErrorName() string {
	return "GetUsersInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUsersInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersInfoRequestValidationError{}

// Validate checks the field values on GetUsersInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetUsersInfoResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUsersInfoResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetUsersInfoResponseValidationError is the validation error returned by
// GetUsersInfoResponse.Validate if the designated constraints aren't met.
type GetUsersInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersInfoResponseValidationError) ErrorName() string {
	return "GetUsersInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUsersInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersInfoResponseValidationError{}

// Validate checks the field values on GetUsersInfoResponse_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetUsersInfoResponse_Data) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsersInfo() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUsersInfoResponse_DataValidationError{
					field:  fmt.Sprintf("UsersInfo[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Page

	// no validation rules for PageSize

	return nil
}

// GetUsersInfoResponse_DataValidationError is the validation error returned by
// GetUsersInfoResponse_Data.Validate if the designated constraints aren't met.
type GetUsersInfoResponse_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersInfoResponse_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersInfoResponse_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersInfoResponse_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersInfoResponse_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersInfoResponse_DataValidationError) ErrorName() string {
	return "GetUsersInfoResponse_DataValidationError"
}

// Error satisfies the builtin error interface
func (e GetUsersInfoResponse_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersInfoResponse_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersInfoResponse_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersInfoResponse_DataValidationError{}