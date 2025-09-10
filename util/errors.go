package util

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Custom error types for better error handling
var (
	ErrInvalidInput     = errors.New("invalid input")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
	ErrNotFound         = errors.New("not found")
	ErrInternalError    = errors.New("internal server error")
	ErrConflict         = errors.New("conflict")
	ErrTooManyRequests  = errors.New("too many requests")
	ErrServiceUnavailable = errors.New("service unavailable")
)

// AppError represents an application error with additional context
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewAppError creates a new application error
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Details string `json:"details,omitempty"`
}

// HandleError handles errors and returns appropriate HTTP responses
func HandleError(c *gin.Context, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.Code, ErrorResponse{
			Error:   appErr.Message,
			Code:    appErr.Code,
			Details: appErr.Details,
		})
		return
	}

	// Default to internal server error for unknown errors
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Error: "Internal server error",
		Code:  http.StatusInternalServerError,
	})
}

// ConvertToGRPCError converts an AppError to a gRPC error
func ConvertToGRPCError(err error) error {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		return status.Errorf(codes.Internal, "internal server error")
	}

	switch appErr.Code {
	case http.StatusBadRequest:
		return status.Errorf(codes.InvalidArgument, appErr.Message)
	case http.StatusUnauthorized:
		return status.Errorf(codes.Unauthenticated, appErr.Message)
	case http.StatusForbidden:
		return status.Errorf(codes.PermissionDenied, appErr.Message)
	case http.StatusNotFound:
		return status.Errorf(codes.NotFound, appErr.Message)
	case http.StatusConflict:
		return status.Errorf(codes.AlreadyExists, appErr.Message)
	case http.StatusTooManyRequests:
		return status.Errorf(codes.ResourceExhausted, appErr.Message)
	case http.StatusServiceUnavailable:
		return status.Errorf(codes.Unavailable, appErr.Message)
	default:
		return status.Errorf(codes.Internal, "internal server error")
	}
}

// Common error constructors
func NewBadRequestError(message string, err error) *AppError {
	return NewAppError(http.StatusBadRequest, message, err)
}

func NewUnauthorizedError(message string, err error) *AppError {
	return NewAppError(http.StatusUnauthorized, message, err)
}

func NewForbiddenError(message string, err error) *AppError {
	return NewAppError(http.StatusForbidden, message, err)
}

func NewNotFoundError(message string, err error) *AppError {
	return NewAppError(http.StatusNotFound, message, err)
}

func NewConflictError(message string, err error) *AppError {
	return NewAppError(http.StatusConflict, message, err)
}

func NewInternalError(message string, err error) *AppError {
	return NewAppError(http.StatusInternalServerError, message, err)
}

