package validator

// Validator is used to validate fields of a given struct.
// It contains a map of fields and its error messages.
type Validator struct {
	Errors map[string]string
}

// New returns a new Validator instance.
func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
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

// UniqueStrings checks whether given string slice contains only unique values.
// Returns true on success.
func (v *Validator) UniqueStrings(values []string) bool {
	unique := make(map[string]bool)

	for _, value := range values {
		unique[value] = true
	}

	return len(values) == len(unique)
}
