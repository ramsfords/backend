// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user.proto

package user

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

// Validate checks the field values on UserData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserDataMultiError, or nil
// if none found.
func (m *UserData) ValidateAll() error {
	return m.validate(true)
}

func (m *UserData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for CognitoId

	// no validation rules for UserName

	// no validation rules for UserId

	// no validation rules for FirstName

	// no validation rules for MiddleName

	// no validation rules for LastName

	// no validation rules for Email

	// no validation rules for HashedPassword

	// no validation rules for AvatarUrl

	// no validation rules for NewPasswordRequired

	// no validation rules for PasswordChangedAt

	// no validation rules for CreatedOn

	// no validation rules for UpdatedOn

	// no validation rules for DeletedOn

	for idx, item := range m.GetPhoneNumbers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserDataValidationError{
						field:  fmt.Sprintf("PhoneNumbers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserDataValidationError{
						field:  fmt.Sprintf("PhoneNumbers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserDataValidationError{
					field:  fmt.Sprintf("PhoneNumbers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for EmailVerified

	// no validation rules for Pk

	// no validation rules for Sk

	// no validation rules for UnsuscribedToMarketingEmail

	if len(errors) > 0 {
		return UserDataMultiError(errors)
	}

	return nil
}

// UserDataMultiError is an error wrapping multiple validation errors returned
// by UserData.ValidateAll() if the designated constraints aren't met.
type UserDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserDataMultiError) AllErrors() []error { return m }

// UserDataValidationError is the validation error returned by
// UserData.Validate if the designated constraints aren't met.
type UserDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserDataValidationError) ErrorName() string { return "UserDataValidationError" }

// Error satisfies the builtin error interface
func (e UserDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserDataValidationError{}

// Validate checks the field values on MeData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MeData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MeData with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MeDataMultiError, or nil if none found.
func (m *MeData) ValidateAll() error {
	return m.validate(true)
}

func (m *MeData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MeDataMultiError(errors)
	}

	return nil
}

// MeDataMultiError is an error wrapping multiple validation errors returned by
// MeData.ValidateAll() if the designated constraints aren't met.
type MeDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MeDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MeDataMultiError) AllErrors() []error { return m }

// MeDataValidationError is the validation error returned by MeData.Validate if
// the designated constraints aren't met.
type MeDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MeDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MeDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MeDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MeDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MeDataValidationError) ErrorName() string { return "MeDataValidationError" }

// Error satisfies the builtin error interface
func (e MeDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMeData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MeDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MeDataValidationError{}

// Validate checks the field values on SignUpData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignUpData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignUpData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignUpDataMultiError, or
// nil if none found.
func (m *SignUpData) ValidateAll() error {
	return m.validate(true)
}

func (m *SignUpData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for ConfirmPassword

	// no validation rules for Email

	// no validation rules for Name

	// no validation rules for Origin

	// no validation rules for EmailVisibility

	// no validation rules for RestaurantIds

	// no validation rules for Type

	if len(errors) > 0 {
		return SignUpDataMultiError(errors)
	}

	return nil
}

// SignUpDataMultiError is an error wrapping multiple validation errors
// returned by SignUpData.ValidateAll() if the designated constraints aren't met.
type SignUpDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignUpDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignUpDataMultiError) AllErrors() []error { return m }

// SignUpDataValidationError is the validation error returned by
// SignUpData.Validate if the designated constraints aren't met.
type SignUpDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpDataValidationError) ErrorName() string { return "SignUpDataValidationError" }

// Error satisfies the builtin error interface
func (e SignUpDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpDataValidationError{}

// Validate checks the field values on UserResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserResponseMultiError, or
// nil if none found.
func (m *UserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	// no validation rules for Id

	if len(errors) > 0 {
		return UserResponseMultiError(errors)
	}

	return nil
}

// UserResponseMultiError is an error wrapping multiple validation errors
// returned by UserResponse.ValidateAll() if the designated constraints aren't met.
type UserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserResponseMultiError) AllErrors() []error { return m }

// UserResponseValidationError is the validation error returned by
// UserResponse.Validate if the designated constraints aren't met.
type UserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserResponseValidationError) ErrorName() string { return "UserResponseValidationError" }

// Error satisfies the builtin error interface
func (e UserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserResponseValidationError{}

// Validate checks the field values on Login with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Login) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Login with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in LoginMultiError, or nil if none found.
func (m *Login) ValidateAll() error {
	return m.validate(true)
}

func (m *Login) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Password

	if len(errors) > 0 {
		return LoginMultiError(errors)
	}

	return nil
}

// LoginMultiError is an error wrapping multiple validation errors returned by
// Login.ValidateAll() if the designated constraints aren't met.
type LoginMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginMultiError) AllErrors() []error { return m }

// LoginValidationError is the validation error returned by Login.Validate if
// the designated constraints aren't met.
type LoginValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginValidationError) ErrorName() string { return "LoginValidationError" }

// Error satisfies the builtin error interface
func (e LoginValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginValidationError{}

// Validate checks the field values on ResetPassword with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ResetPassword) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResetPassword with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResetPasswordMultiError, or
// nil if none found.
func (m *ResetPassword) ValidateAll() error {
	return m.validate(true)
}

func (m *ResetPassword) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for Password

	// no validation rules for ConfirmPassword

	if len(errors) > 0 {
		return ResetPasswordMultiError(errors)
	}

	return nil
}

// ResetPasswordMultiError is an error wrapping multiple validation errors
// returned by ResetPassword.ValidateAll() if the designated constraints
// aren't met.
type ResetPasswordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResetPasswordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResetPasswordMultiError) AllErrors() []error { return m }

// ResetPasswordValidationError is the validation error returned by
// ResetPassword.Validate if the designated constraints aren't met.
type ResetPasswordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetPasswordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetPasswordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetPasswordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetPasswordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetPasswordValidationError) ErrorName() string { return "ResetPasswordValidationError" }

// Error satisfies the builtin error interface
func (e ResetPasswordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetPassword.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetPasswordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetPasswordValidationError{}

// Validate checks the field values on NonAuthUser with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NonAuthUser) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NonAuthUser with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NonAuthUserMultiError, or
// nil if none found.
func (m *NonAuthUser) ValidateAll() error {
	return m.validate(true)
}

func (m *NonAuthUser) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Name

	// no validation rules for Origin

	// no validation rules for TokenKey

	// no validation rules for RestaurantIds

	// no validation rules for Type

	if len(errors) > 0 {
		return NonAuthUserMultiError(errors)
	}

	return nil
}

// NonAuthUserMultiError is an error wrapping multiple validation errors
// returned by NonAuthUser.ValidateAll() if the designated constraints aren't met.
type NonAuthUserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NonAuthUserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NonAuthUserMultiError) AllErrors() []error { return m }

// NonAuthUserValidationError is the validation error returned by
// NonAuthUser.Validate if the designated constraints aren't met.
type NonAuthUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NonAuthUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NonAuthUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NonAuthUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NonAuthUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NonAuthUserValidationError) ErrorName() string { return "NonAuthUserValidationError" }

// Error satisfies the builtin error interface
func (e NonAuthUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNonAuthUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NonAuthUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NonAuthUserValidationError{}
