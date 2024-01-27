// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: paginator/v1/paginator.proto

package v1

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

// Validate checks the field values on SearchParam with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SearchParam) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchParam with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SearchParamMultiError, or
// nil if none found.
func (m *SearchParam) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchParam) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Field

	// no validation rules for Value

	// no validation rules for Exp

	// no validation rules for Logic

	if len(errors) > 0 {
		return SearchParamMultiError(errors)
	}

	return nil
}

// SearchParamMultiError is an error wrapping multiple validation errors
// returned by SearchParam.ValidateAll() if the designated constraints aren't met.
type SearchParamMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchParamMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchParamMultiError) AllErrors() []error { return m }

// SearchParamValidationError is the validation error returned by
// SearchParam.Validate if the designated constraints aren't met.
type SearchParamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchParamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchParamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchParamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchParamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchParamValidationError) ErrorName() string { return "SearchParamValidationError" }

// Error satisfies the builtin error interface
func (e SearchParamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchParam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchParamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchParamValidationError{}

// Validate checks the field values on OrderParam with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderParam) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderParam with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderParamMultiError, or
// nil if none found.
func (m *OrderParam) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderParam) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Field

	// no validation rules for Exp

	if len(errors) > 0 {
		return OrderParamMultiError(errors)
	}

	return nil
}

// OrderParamMultiError is an error wrapping multiple validation errors
// returned by OrderParam.ValidateAll() if the designated constraints aren't met.
type OrderParamMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderParamMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderParamMultiError) AllErrors() []error { return m }

// OrderParamValidationError is the validation error returned by
// OrderParam.Validate if the designated constraints aren't met.
type OrderParamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderParamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderParamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderParamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderParamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderParamValidationError) ErrorName() string { return "OrderParamValidationError" }

// Error satisfies the builtin error interface
func (e OrderParamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderParam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderParamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderParamValidationError{}

// Validate checks the field values on PaginatorReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PaginatorReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PaginatorReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PaginatorReqMultiError, or
// nil if none found.
func (m *PaginatorReq) ValidateAll() error {
	return m.validate(true)
}

func (m *PaginatorReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	for idx, item := range m.GetSearch() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PaginatorReqValidationError{
						field:  fmt.Sprintf("Search[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PaginatorReqValidationError{
						field:  fmt.Sprintf("Search[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PaginatorReqValidationError{
					field:  fmt.Sprintf("Search[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetOrder() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PaginatorReqValidationError{
						field:  fmt.Sprintf("Order[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PaginatorReqValidationError{
						field:  fmt.Sprintf("Order[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PaginatorReqValidationError{
					field:  fmt.Sprintf("Order[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PaginatorReqMultiError(errors)
	}

	return nil
}

// PaginatorReqMultiError is an error wrapping multiple validation errors
// returned by PaginatorReq.ValidateAll() if the designated constraints aren't met.
type PaginatorReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginatorReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginatorReqMultiError) AllErrors() []error { return m }

// PaginatorReqValidationError is the validation error returned by
// PaginatorReq.Validate if the designated constraints aren't met.
type PaginatorReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginatorReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginatorReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginatorReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginatorReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginatorReqValidationError) ErrorName() string { return "PaginatorReqValidationError" }

// Error satisfies the builtin error interface
func (e PaginatorReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginatorReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginatorReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginatorReqValidationError{}

// Validate checks the field values on PaginatorReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PaginatorReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PaginatorReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PaginatorReplyMultiError,
// or nil if none found.
func (m *PaginatorReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PaginatorReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	// no validation rules for Total

	if len(errors) > 0 {
		return PaginatorReplyMultiError(errors)
	}

	return nil
}

// PaginatorReplyMultiError is an error wrapping multiple validation errors
// returned by PaginatorReply.ValidateAll() if the designated constraints
// aren't met.
type PaginatorReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginatorReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginatorReplyMultiError) AllErrors() []error { return m }

// PaginatorReplyValidationError is the validation error returned by
// PaginatorReply.Validate if the designated constraints aren't met.
type PaginatorReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginatorReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginatorReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginatorReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginatorReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginatorReplyValidationError) ErrorName() string { return "PaginatorReplyValidationError" }

// Error satisfies the builtin error interface
func (e PaginatorReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginatorReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginatorReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginatorReplyValidationError{}