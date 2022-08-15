package comoperation

type StatusCode string

var (
	BadArgumentsCode          StatusCode = "BAD_ARGUMENTS"
	ResourceAlreadyExistsCode StatusCode = "BadArgs"
	GeneralErrorCode          StatusCode = "GeneralErrorCode"
	InternalErrorCode         StatusCode = "INTERNAL_ERROR"
)

const BadArgsMessage = "Bad arguments provided"
const ResourceAlreadyExistsMessage = "ResourceAlreadyExistsCode"
const GeneralErrorMessage = "GeneralError"

type Operation struct {
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	Data          interface{}
}

func NewOperation(statusCode StatusCode, statusMessage string) *Operation {
	return &Operation{
		StatusCode:    string(statusCode),
		StatusMessage: statusMessage,
		Data:          nil,
	}
}
func (o *Operation) Error() string {
	return o.StatusCode + " " + o.StatusMessage
}
