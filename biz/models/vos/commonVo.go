package vos

import "errors"

type CommonResp struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg,omitempty"`
	Succeed   bool   `json:"succeed"`
	Data      any    `json:"data,omitempty"`
	DataArray []any  `json:"data_array,omitempty"`
}

// operation base on id
type IDOpt struct {
	ID string `path:"id" param:"id" query:"id" json:"id" vd:"len($) > 0"`
}

const (
	SucceedCode       = 0
	InvalidParamCode  = 10001
	DataNotExistCode  = 10002
	DateNotActiveCode = 10003
	InValidBodyCode   = 10004
)

var (
	ErrInvalidParam  = errors.New("invalid params")
	ErrInvalidBody   = errors.New("invalid body")
	ErrDataNotExist  = errors.New("data not exist")
	ErrDataNotActive = errors.New("data not active")
)
