package error

import (
	"fmt"
)

type AppError struct {
	StatusCode   string `json:"StatusCode,omitempty"`
	ErrorCode    string
	ErrorFields  string `json:"ErrorFields,omitempty"`
	ErrorMessage string
	Err          error  `json:"Err,omitempty"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.ErrorMessage, e.Err)
	}
	return e.ErrorMessage
}
func (e *AppError) Unwrap() error {
	return e.Err
}

func NewErrorDopaInvalid(dopaMsg string) *AppError {
    return &AppError{
        ErrorCode:    "COM014",
        ErrorMessage: "Invalid Status/" + dopaMsg,
    }
}

var (
	ErrInternalServer   = &AppError{ErrorCode: "SYS500", ErrorMessage: "An unexpected internal error occurred"}
	ErrTimeOut          = &AppError{ErrorCode: "SYS003", ErrorMessage: "System Time out"}
	ErrUnauthorized     = &AppError{ErrorCode: "SYS002", ErrorMessage: "Unauthorized"}
	ErrService          = &AppError{ErrorCode: "SYS001", ErrorMessage: "System unavailable"}
	ErrDopa             = &AppError{ErrorCode: "SYS004", ErrorMessage: "DOPA Unavailable"}

	ErrApiDeivceOS      = &AppError{ErrorCode: "COM034", ErrorMessage: "Invalid Api-DeviceOS"}
	ErrApiChannel       = &AppError{ErrorCode: "COM002", ErrorMessage: "Invalid Api-Channel"}
	ErrApiRequestID     = &AppError{ErrorCode: "COM033", ErrorMessage: "Invalid Api-RequestID"}
	ErrRequiedParam     = &AppError{ErrorCode: "COM001", ErrorMessage: "Required Parameter"}


)

type ErrorResponseFormat struct {
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
}

type ValidationErrorDetail struct {
	Field        string `json:"field"`
	Tag          string `json:"tag"`
	Message      string `json:"message"`
}
