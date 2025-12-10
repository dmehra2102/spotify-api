package errors

import "fmt"

type DomainError struct {
	Code    string
	Message string
	Details map[string]any
	Status  int
}

func NewDomainError(code, message string, status int) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Status:  status,
		Details: make(map[string]any, 0),
	}
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}
func (e *DomainError) WithDetails(key string, value any) *DomainError {
	e.Details[key] = value
	return e
}

var (
	ErrUnauthorized = NewDomainError("UNAUTHORIZED", "Authentication required", 401)
	ErrForbidden    = NewDomainError("FORBIDDEN", "Insufficient permissions", 403)
	ErrNotFound     = NewDomainError("NOT_FOUND", "Resource not found", 404)
	ErrConflict     = NewDomainError("CONFLICT", "Resource already exists", 409)
	ErrValidation   = NewDomainError("VALIDATION_ERROR", "Input validation failed", 400)
	ErrInternal     = NewDomainError("INTERNAL_ERROR", "Internal server error", 500)
	ErrRateLimit    = NewDomainError("RATE_LIMITED", "Rate limit exceeded", 429)
)
