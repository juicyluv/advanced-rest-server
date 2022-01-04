package validator

import (
	"fmt"
	"regexp"
)

// Email RegEx
var (
	emailRegEx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type ValidatorErrors map[string]interface{}

// Validator is used to validate fields of a given struct.
// It contains a map of fields and its error messages.
type Validator struct {
	Errors ValidatorErrors
}

// New returns a new Validator instance.
func New() *Validator {
	return &Validator{
		Errors: make(map[string]interface{}),
	}
}

// Valid checks if there is no any error in current validator state.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// addError adds a new error to error map.
func (v *Validator) addError(field, message string) {
	if _, exists := v.Errors[field]; !exists {
		v.Errors[field] = message
	}
}

// Check adds a new error to the Errors map if condition is not equal to true.
func (v *Validator) Check(condition bool, field, errorMessage string) {
	if !condition {
		v.addError(field, errorMessage)
	}
}

// IsEmail return true if a given string is email.
func IsEmail(value string) bool {
	return emailRegEx.MatchString(value)
}

// UniqueStrings checks whether given string slice contains only unique values.
// Returns true on success.
func (v *Validator) UniqueStrings(values []string) bool {
	unique := make(map[string]bool)

	for _, value := range values {
		unique[value] = true
	}

	return len(values) == len(unique)
}

// MinLength checks whether given string has provided length.
func (v *Validator) MinLength(value string, length int, field string) {
	message := fmt.Sprintf("the length must be greater than %d characters", length+1)
	v.Check(len(value) >= length, field, message)
}

// MaxLength checks whether given string has provided length.
func (v *Validator) MaxLength(value string, length int, field string) {
	message := fmt.Sprintf("the length must be not greater than %d characters", length)
	v.Check(len(value) <= length, field, message)
}

// LengthBetween checks whether given string has length between provided numbers.
func (v *Validator) LengthBetween(value string, min, max int, field string) {
	v.MinLength(value, min, field)
	v.MaxLength(value, max, field)
}

// Required checks whether given value is empty
func (v *Validator) NotEmpty(value string, field string) {
	v.Check(value != "", field, "must be provided")
}

// Min checks whether given value is greater or equal to provided value.
func (v *Validator) Min(value int, min int, field string) {
	message := fmt.Sprintf("must be greater or equal to %d", min)
	v.Check(value >= min, field, message)
}

// Max checks whether given value is less or equal to provided value.
func (v *Validator) Max(value int, max int, field string) {
	message := fmt.Sprintf("must be less or equal to %d", max)
	v.Check(value <= max, field, message)
}
