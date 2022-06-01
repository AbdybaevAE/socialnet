package op

type StatusCode string

var (
	BadArgsCode               StatusCode = "BadArgs"
	ResourceAlreadyExistsCode StatusCode = "BadArgs"
	GeneralErrorCode          StatusCode = "GeneralErrorCode"
)

const BadArgsMessage = "Bad arguments provided"
const ResourceAlreadyExistsMessage = "ResourceAlreadyExistsCode"
const GeneralErrorMessage = "GeneralError"

type Operation struct {
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

func NewOperation(statusCode StatusCode, statusMessage string) *Operation {
	return &Operation{
		StatusCode:    string(statusCode),
		StatusMessage: statusMessage,
	}
}
func (o *Operation) Error() string {
	return o.StatusCode + " " + o.StatusMessage
}

var BadArgsOp = NewOperation(BadArgsCode, BadArgsMessage)
var ResourceAlreadyExistsOp = NewOperation(ResourceAlreadyExistsCode, ResourceAlreadyExistsMessage)
