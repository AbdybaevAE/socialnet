package comerrors

import "github.com/abdybaevae/socialnet/pkg/lib/comoperation"

var (
	ErrBadArguments         = comoperation.NewOperation(comoperation.BadArgumentsCode, comoperation.BadArgsMessage)
	ErrResourceAlreadyExist = comoperation.NewOperation(comoperation.ResourceAlreadyExistsCode, comoperation.ResourceAlreadyExistsMessage)
	ErrInternalError        = comoperation.NewOperation(comoperation.InternalErrorCode, comoperation.GeneralErrorMessage)
)
