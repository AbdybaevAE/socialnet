package lang

import "github.com/abdybaevae/socialnet/pkg/entities"

type CreateLangArgs struct {
	Code string
	Name string
}
type GetLangsItem struct {
	LangId   int
	LangCode string
	LangName string
}
type Service interface {
	CreateLang(args *CreateLangArgs) error
	GetLangList() []GetLangsItem
}

func NewService() Service {
	return &impl{}
}

type impl struct {
}

func (i *impl) CreateLang(args *CreateLangArgs) error {
	if len(args.Code) == 0 || len(args.Name) {
		return
	}
	langEntity := entities.LanguageEntity{
		Name: args.Name,
		Code: args.Code,
	}
	if
	return nil
}
func (i *impl) GetLangList() (error, []GetLangsItem) {
	return nil, nil
}
