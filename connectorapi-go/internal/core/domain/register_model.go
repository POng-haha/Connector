package domain

import (
	"time"

	appError "connectorapi-go/pkg/error"

	"github.com/gin-gonic/gin"
)

// ---------- API CheckRegister ---------
type CheckRegisterRequest struct {
	IDCardNo                 string `json:"IDCardNo"               validate:"required,max=20"`
	MobileNo 			     string `json:"MobileNo"               validate:"required,max=10"`
	AgreementNo 		     string `json:"AgreementNo"            validate:"max=16"`
}

type CheckRegisterResponse struct {
	IDCardNo                      string	                      `json:"IDCardNo"` 
	CustomerNameTH                string 	                      `json:"CustomerNameTH"`
	CustomerNameEN                string	                      `json:"CustomerNameEN"` 
	MobileNo                      string 	                      `json:"MobileNo"`
	Email                         string	                      `json:"Email"` 
	Result                        string 	                      `json:"Result"`
	ResultCode                    string	                      `json:"ResultCode"` 
	CRRegisterFlag                string 	                      `json:"CRRegisterFlag"`
	DYCRegisterFlag               string	                      `json:"DYCRegisterFlag"` 
	AgreementRegisterFlag         string 	                      `json:"AgreementRegisterFlag"`
}

type CheckRegisterResult struct {
	Response       *CheckRegisterResponse
    AppError       *appError.AppError
    GinCtx         *gin.Context
    Timestamp      time.Time
    ReqBody        interface{}
    RespBody       interface{}
    DomainError    *appError.AppError
    ServiceName    string
	UserRef        string
    LogLine1       string
}

// ---------- API CheckRegisterSocial ---------
type CheckRegisterSocialRequest struct {
	IDCardNo                 string `json:"IDCardNo"               validate:"required,max=20"`
}

type CheckRegisterSocialResponse struct {
	IDCardNo                      string	                      `json:"IDCardNo"` 
	MobileNo                      string 	                      `json:"MobileNo"`
}

type CheckRegisterSocialResult struct {
	Response       *CheckRegisterSocialResponse
    AppError       *appError.AppError
    GinCtx         *gin.Context
    Timestamp      time.Time
    ReqBody        interface{}
    RespBody       interface{}
    DomainError    *appError.AppError
    ServiceName    string
	UserRef        string
    LogLine1       string
}