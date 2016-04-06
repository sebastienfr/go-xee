package xee

import (
    "fmt"
)

var (
	// ErrForbidden when resource is not accessible
	ErrForbidden = fmt.Errorf("Forbidden")

	// ErrValidation appen when a struct is not valid
	ErrValidation = fmt.Errorf("Validation error")

    // ErrEntityNotFound appen when entity is not found
    ErrEntityNotFound = fmt.Errorf("Entity not found")
)
