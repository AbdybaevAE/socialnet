package op

import "fmt"

type StatusCode int

const (
	OkStatusMsg     = "Successfull operation."
	BadArgsMsg      = "BadArgsMessage"
	DupUserEmailMsg = "Duplicate user email found."
)

const (
	Ok StatusCode = iota
	BadArgs
	DupUserEmail
)

type Operation interface {
	Code() StatusCode
	error
	Msg() string
}
type OperationImpl struct {
	statusCode StatusCode
	statusMsg  string
}

func (o *OperationImpl) Code() StatusCode {
	return o.statusCode
}
func (o *OperationImpl) Msg() string {
	return o.statusMsg
}
func (o *OperationImpl) Error() string {
	return fmt.Sprintf("code is %v and msg is %v", o.statusCode, o.statusMsg)
}
func New(statusCode StatusCode, statusMsg string) Operation {
	return &OperationImpl{
		statusCode: statusCode,
		statusMsg:  statusMsg,
	}
}

var OkOperation = New(Ok, OkStatusMsg)
