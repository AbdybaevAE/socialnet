package op

type StatusCode string

var (
	BadArgs               StatusCode = "BadArgs"
	ResourceAlreadyExists StatusCode = "BadArgs"
	GeneralErrorCode      StatusCode = "GeneralErrorCode"
)

const BadArgsMessage = "Bad arguments provided"
const ResourceAlreadyExistsMessage = "ResourceAlreadyExists"
const GeneralErrorMessage = "GeneralError"

type Operation struct {
	StatusCode string `json:"statusCode"`
	Message    string `json:"statusMessage"`
}

func NewOperation(statusCode StatusCode, statusMessage string) *Operation {
	return &Operation{
		StatusCode: string(statusCode),
		Message:    statusMessage,
	}
}
