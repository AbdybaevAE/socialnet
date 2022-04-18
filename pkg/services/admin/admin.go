package admin

type AdminService interface {
	CreateLanguage(args *CreateLanguageArgs) error
	GetLanguages() ([]GetLanguageItem, error)
}
type CreateLanguageArgs struct {
	Name string
	Code string
}
type GetLanguageItem struct {
	LangId   int
	LangName string
	LangCode string
}
type adminService struct {
}

func (a *adminService) CreateLanguage(args *CreateLanguageArgs) error {

}
func (a *adminService) GetLanguages() ([]GetLanguageItem, error) {

}

func NewAdminService() AdminService {
	return &adminService{}
}
