package error

import (
	"fmt"
)

type AppError struct {
	StatusCode string
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}
func (e *AppError) Unwrap() error {
	return e.Err
}

var (
	ErrNotFound       = &AppError{Code: "SVC404", Message: "Resource not found"}
	ErrValidation     = &AppError{Code: "VAL400", Message: "Invalid input provided"}
	ErrConfig         = &AppError{Code: "CFG500", Message: "Application configuration error"}
	ErrForbidden      = &AppError{Code: "SEC002", Message: "Permission denied"}

	ErrService          = &AppError{Code: "SYS001", Message: "System unavailable"}
	ErrUnauthorized     = &AppError{Code: "SYS002", Message: "Unauthorized"}
	ErrTimeOut          = &AppError{Code: "SYS003", Message: "System Time out"}
	ErrInternalServer   = &AppError{Code: "SYS500", Message: "An unexpected internal error occurred"}
	ErrInternalLength   = &AppError{Code: "SYS500", Message: "An unexpected internal error occurred: max length"}
	ErrSystemI  		= &AppError{Code: "SYS008", Message: "System-I Unavailable"}
	ErrSystemIUnexpect	= &AppError{Code: "SYS009", Message: "System-I Unexpected error occurred"}

	ErrRequiedParam     = &AppError{Code: "COM001", Message: "Required Parameter"}
	ErrApiChannel       = &AppError{Code: "COM002", Message: "Invalid Api-Channel"}
	ErrApiRequestID     = &AppError{Code: "COM033", Message: "Invalid Api-RequestID"}
	ErrApiDeviceOS      = &AppError{Code: "COM034", Message: "Invalid Api-DeviceOS"}
	ErrIDCardNotFound   = &AppError{Code: "COM067", Message: "ID Card No. Not Found"}

	ErrSUEInfoNotFound  = &AppError{Code: "COL001", Message: "SUE Information Not Found"}

	ErrAgmNoNotFound  	= &AppError{Code: "UHP003", Message: "Agreement No. Not Foundd"}

)

type ErrorResponse struct {
	ErrorCode    string    `json:"Code"`
	ErrorMessage string    `json:"Message"`
}

type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}