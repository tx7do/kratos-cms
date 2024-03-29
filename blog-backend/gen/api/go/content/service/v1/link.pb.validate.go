// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: content/service/v1/link.proto

package servicev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on Link with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Link) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Link with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in LinkMultiError, or nil if none found.
func (m *Link) ValidateAll() error {
	return m.validate(true)
}

func (m *Link) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.Url != nil {
		// no validation rules for Url
	}

	if m.Logo != nil {
		// no validation rules for Logo
	}

	if m.Description != nil {
		// no validation rules for Description
	}

	if m.Team != nil {
		// no validation rules for Team
	}

	if m.Priority != nil {
		// no validation rules for Priority
	}

	if m.CreateTime != nil {
		// no validation rules for CreateTime
	}

	if m.UpdateTime != nil {
		// no validation rules for UpdateTime
	}

	if len(errors) > 0 {
		return LinkMultiError(errors)
	}

	return nil
}

// LinkMultiError is an error wrapping multiple validation errors returned by
// Link.ValidateAll() if the designated constraints aren't met.
type LinkMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LinkMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LinkMultiError) AllErrors() []error { return m }

// LinkValidationError is the validation error returned by Link.Validate if the
// designated constraints aren't met.
type LinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LinkValidationError) ErrorName() string { return "LinkValidationError" }

// Error satisfies the builtin error interface
func (e LinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LinkValidationError{}

// Validate checks the field values on ListLinkResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListLinkResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListLinkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListLinkResponseMultiError, or nil if none found.
func (m *ListLinkResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListLinkResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListLinkResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListLinkResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListLinkResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListLinkResponseMultiError(errors)
	}

	return nil
}

// ListLinkResponseMultiError is an error wrapping multiple validation errors
// returned by ListLinkResponse.ValidateAll() if the designated constraints
// aren't met.
type ListLinkResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListLinkResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListLinkResponseMultiError) AllErrors() []error { return m }

// ListLinkResponseValidationError is the validation error returned by
// ListLinkResponse.Validate if the designated constraints aren't met.
type ListLinkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListLinkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListLinkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListLinkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListLinkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListLinkResponseValidationError) ErrorName() string { return "ListLinkResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListLinkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListLinkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListLinkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListLinkResponseValidationError{}

// Validate checks the field values on GetLinkRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetLinkRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLinkRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetLinkRequestMultiError,
// or nil if none found.
func (m *GetLinkRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLinkRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetLinkRequestMultiError(errors)
	}

	return nil
}

// GetLinkRequestMultiError is an error wrapping multiple validation errors
// returned by GetLinkRequest.ValidateAll() if the designated constraints
// aren't met.
type GetLinkRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLinkRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLinkRequestMultiError) AllErrors() []error { return m }

// GetLinkRequestValidationError is the validation error returned by
// GetLinkRequest.Validate if the designated constraints aren't met.
type GetLinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLinkRequestValidationError) ErrorName() string { return "GetLinkRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetLinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLinkRequestValidationError{}

// Validate checks the field values on CreateLinkRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateLinkRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateLinkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateLinkRequestMultiError, or nil if none found.
func (m *CreateLinkRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateLinkRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLink()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateLinkRequestValidationError{
					field:  "Link",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateLinkRequestValidationError{
					field:  "Link",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLink()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateLinkRequestValidationError{
				field:  "Link",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return CreateLinkRequestMultiError(errors)
	}

	return nil
}

// CreateLinkRequestMultiError is an error wrapping multiple validation errors
// returned by CreateLinkRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateLinkRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateLinkRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateLinkRequestMultiError) AllErrors() []error { return m }

// CreateLinkRequestValidationError is the validation error returned by
// CreateLinkRequest.Validate if the designated constraints aren't met.
type CreateLinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateLinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateLinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateLinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateLinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateLinkRequestValidationError) ErrorName() string {
	return "CreateLinkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateLinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateLinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateLinkRequestValidationError{}

// Validate checks the field values on UpdateLinkRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateLinkRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateLinkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateLinkRequestMultiError, or nil if none found.
func (m *UpdateLinkRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateLinkRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetLink()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateLinkRequestValidationError{
					field:  "Link",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateLinkRequestValidationError{
					field:  "Link",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLink()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateLinkRequestValidationError{
				field:  "Link",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return UpdateLinkRequestMultiError(errors)
	}

	return nil
}

// UpdateLinkRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateLinkRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateLinkRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateLinkRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateLinkRequestMultiError) AllErrors() []error { return m }

// UpdateLinkRequestValidationError is the validation error returned by
// UpdateLinkRequest.Validate if the designated constraints aren't met.
type UpdateLinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateLinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateLinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateLinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateLinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateLinkRequestValidationError) ErrorName() string {
	return "UpdateLinkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateLinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateLinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateLinkRequestValidationError{}

// Validate checks the field values on DeleteLinkRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteLinkRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteLinkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteLinkRequestMultiError, or nil if none found.
func (m *DeleteLinkRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteLinkRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return DeleteLinkRequestMultiError(errors)
	}

	return nil
}

// DeleteLinkRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteLinkRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteLinkRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteLinkRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteLinkRequestMultiError) AllErrors() []error { return m }

// DeleteLinkRequestValidationError is the validation error returned by
// DeleteLinkRequest.Validate if the designated constraints aren't met.
type DeleteLinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteLinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteLinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteLinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteLinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteLinkRequestValidationError) ErrorName() string {
	return "DeleteLinkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteLinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteLinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteLinkRequestValidationError{}
