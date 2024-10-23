// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: center/service.proto

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

// Validate checks the field values on Author with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Author) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Alias

	// no validation rules for Avatar

	return nil
}

// AuthorValidationError is the validation error returned by Author.Validate if
// the designated constraints aren't met.
type AuthorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorValidationError) ErrorName() string { return "AuthorValidationError" }

// Error satisfies the builtin error interface
func (e AuthorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorValidationError{}

// Validate checks the field values on Interaction with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Interaction) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InteractionValidationError{
				field:  "Author",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Type

	// no validation rules for Content

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	return nil
}

// InteractionValidationError is the validation error returned by
// Interaction.Validate if the designated constraints aren't met.
type InteractionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InteractionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InteractionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InteractionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InteractionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InteractionValidationError) ErrorName() string { return "InteractionValidationError" }

// Error satisfies the builtin error interface
func (e InteractionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInteraction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InteractionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InteractionValidationError{}

// Validate checks the field values on PostDetail with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PostDetail) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PostDetailValidationError{
				field:  "Author",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Content

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	for idx, item := range m.GetInteractions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PostDetailValidationError{
					field:  fmt.Sprintf("Interactions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PostDetailValidationError is the validation error returned by
// PostDetail.Validate if the designated constraints aren't met.
type PostDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostDetailValidationError) ErrorName() string { return "PostDetailValidationError" }

// Error satisfies the builtin error interface
func (e PostDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPostDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostDetailValidationError{}

// Validate checks the field values on GetPostsDetailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPostsDetailRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PostId

	return nil
}

// GetPostsDetailRequestValidationError is the validation error returned by
// GetPostsDetailRequest.Validate if the designated constraints aren't met.
type GetPostsDetailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostsDetailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostsDetailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostsDetailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostsDetailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostsDetailRequestValidationError) ErrorName() string {
	return "GetPostsDetailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPostsDetailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostsDetailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostsDetailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostsDetailRequestValidationError{}

// Validate checks the field values on GetPostsDetailResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPostsDetailResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPostsDetailResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetPostsDetailResponseValidationError is the validation error returned by
// GetPostsDetailResponse.Validate if the designated constraints aren't met.
type GetPostsDetailResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostsDetailResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostsDetailResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostsDetailResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostsDetailResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostsDetailResponseValidationError) ErrorName() string {
	return "GetPostsDetailResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPostsDetailResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostsDetailResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostsDetailResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostsDetailResponseValidationError{}

// Validate checks the field values on GetPostsDetailResponse_Data with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPostsDetailResponse_Data) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPostsDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPostsDetailResponse_DataValidationError{
				field:  "PostsDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetPostsDetailResponse_DataValidationError is the validation error returned
// by GetPostsDetailResponse_Data.Validate if the designated constraints
// aren't met.
type GetPostsDetailResponse_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostsDetailResponse_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostsDetailResponse_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostsDetailResponse_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostsDetailResponse_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostsDetailResponse_DataValidationError) ErrorName() string {
	return "GetPostsDetailResponse_DataValidationError"
}

// Error satisfies the builtin error interface
func (e GetPostsDetailResponse_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostsDetailResponse_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostsDetailResponse_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostsDetailResponse_DataValidationError{}
