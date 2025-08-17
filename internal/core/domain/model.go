package domain

import (
	"fmt"
	"time"
)

type AppError struct {
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
	ErrServiceDown    = &AppError{Code: "SYS001", Message: "System unavailable"}
	ErrConfig         = &AppError{Code: "CFG500", Message: "Application configuration error"}
	ErrUnauthorized   = &AppError{Code: "SEC001", Message: "Unauthorized"}
	ErrForbidden      = &AppError{Code: "SEC002", Message: "Permission denied"}
	ErrInternalServer = &AppError{Code: "SYS500", Message: "An unexpected internal error occurred"}
)

type ErrorResponse struct {
	ErrorCode    string    `json:"ErrorCode"`
	ErrorMessage string    `json:"ErrorMessage"`
	Status       int       `json:"Status"`
	Timestamp    time.Time `json:"Timestamp"`
	RequestID    string    `json:"Request_ID"`
}

type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

type GetCustomerInfoRequest struct {
	IDCardNo 	string `json:"idcardno"`
	AeonID 		string `json:"aeonid"`
	AgreementNo string `json:"agreementno"`
}

type GetCustomerInfoResponse struct {
	AeonID			string `json:"aeonid"`
	CustomerNameENG string `json:"customernameeng"`
	CustomerNameTH  string `json:"customernameth"`
	Sex				string `json:"sex"`
	MobileNo		string `json:"mobileno"`
	Email			string `json:"email"`
	Nationality		string `json:"nationality"`
	DateofBirth		string `json:"dateofbirth"`
	MemberStatus	string `json:"memberstatus"`
}

type UpdateAddressRequest struct {
	CustomerID string  `json:"customerId" validate:"required,uuid4"`
	Address    Address `json:"address" validate:"required"`
}

// UpdateAddressRequest - nested structure for address details
type Address struct {
	Street     string `json:"street" validate:"required"`
	City       string `json:"city" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required,len=5"`
}

type UpdateAddressResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type GetCustomerInfo004Request struct {
	IDCardNo 	string `json:"idcardno"`
	AeonID 		string `json:"aeonid"`
	AgreementNo string `json:"agreementno"`
}

type GetCustomerInfo004Response struct {
	AeonID			string `json:"aeonid"`
	CustomerNameENG string `json:"customernameeng"`
	CustomerNameTH  string `json:"customernameth"`
	Sex				string `json:"sex"`
	MobileNo		string `json:"mobileno"`
	Email			string `json:"email"`
	Nationality		string `json:"nationality"`
	DateofBirth		string `json:"dateofbirth"`
	MemberStatus	string `json:"memberstatus"`
}

type CollectionDetailRequest struct {
	CustomerID string  `json:"customerId" validate:"required,uuid4"`
	Address    Address `json:"address" validate:"required"`
}

type CollectionDetailResponse struct {
	CustomerID string  `json:"customerId" validate:"required,uuid4"`
	Address    Address `json:"address" validate:"required"`
}

type CollectionLogRequest struct {
	CustomerID string  `json:"customerId" validate:"required,uuid4"`
	Address    Address `json:"address" validate:"required"`
}

type CollectionLogResponse struct {
	CustomerID string  `json:"customerId" validate:"required,uuid4"`
	Address    Address `json:"address" validate:"required"`
}