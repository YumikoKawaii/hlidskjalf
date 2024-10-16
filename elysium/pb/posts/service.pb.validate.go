// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: posts/service.proto

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

// Validate checks the field values on Post with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Post) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Author

	// no validation rules for Content

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	return nil
}

// PostValidationError is the validation error returned by Post.Validate if the
// designated constraints aren't met.
type PostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostValidationError) ErrorName() string { return "PostValidationError" }

// Error satisfies the builtin error interface
func (e PostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostValidationError{}

// Validate checks the field values on UpsertPostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpsertPostRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Author

	// no validation rules for Content

	return nil
}

// UpsertPostRequestValidationError is the validation error returned by
// UpsertPostRequest.Validate if the designated constraints aren't met.
type UpsertPostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertPostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertPostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertPostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertPostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertPostRequestValidationError) ErrorName() string {
	return "UpsertPostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpsertPostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertPostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertPostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertPostRequestValidationError{}

// Validate checks the field values on UpsertPostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpsertPostResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	// no validation rules for Id

	return nil
}

// UpsertPostResponseValidationError is the validation error returned by
// UpsertPostResponse.Validate if the designated constraints aren't met.
type UpsertPostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertPostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertPostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertPostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertPostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertPostResponseValidationError) ErrorName() string {
	return "UpsertPostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpsertPostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertPostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertPostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertPostResponseValidationError{}

// Validate checks the field values on GetPostsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetPostsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Page

	// no validation rules for PageSize

	return nil
}

// GetPostsRequestValidationError is the validation error returned by
// GetPostsRequest.Validate if the designated constraints aren't met.
type GetPostsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostsRequestValidationError) ErrorName() string { return "GetPostsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetPostsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostsRequestValidationError{}

// Validate checks the field values on GetPostResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetPostResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPostResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetPostResponseValidationError is the validation error returned by
// GetPostResponse.Validate if the designated constraints aren't met.
type GetPostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostResponseValidationError) ErrorName() string { return "GetPostResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetPostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostResponseValidationError{}

// Validate checks the field values on GetPostResponse_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPostResponse_Data) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetPostResponse_DataValidationError{
					field:  fmt.Sprintf("Posts[%v]", idx),
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

// GetPostResponse_DataValidationError is the validation error returned by
// GetPostResponse_Data.Validate if the designated constraints aren't met.
type GetPostResponse_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostResponse_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostResponse_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostResponse_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostResponse_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostResponse_DataValidationError) ErrorName() string {
	return "GetPostResponse_DataValidationError"
}

// Error satisfies the builtin error interface
func (e GetPostResponse_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostResponse_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostResponse_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostResponse_DataValidationError{}
