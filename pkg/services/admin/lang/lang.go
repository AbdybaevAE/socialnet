package lang

import (
	"github.com/abdybaevae/socialnet/pkg/entities"
	"github.com/abdybaevae/socialnet/pkg/lib/op"
	"github.com/abdybaevae/socialnet/pkg/repos"
)

type CreateLangArgs struct {
	Code string
	Name string
}
type GetLangItem struct {
	LangId   int
	LangCode string
	LangName string
}
type AdminLangSrv interface {
	CreateLang(args *CreateLangArgs) error
	GetLangList() (res []GetLangItem, err error)
}

func NewService(repo repos.LangRepo) AdminLangSrv {
	return &impl{repo}
}

type impl struct {
	repo repos.LangRepo
}

func (i *impl) CreateLang(args *CreateLangArgs) error {
	if len(args.Code) == 0 || len(args.Name) == 0 {
		return op.BadArgsOp
	}
	langEntity := &entities.LanguageEntity{
		Name: args.Name,
		Code: args.Code,
	}
	if err := i.repo.CreateLang(langEntity); err != nil {
		return err
	}
	return nil
}
func (i *impl) GetLangList() ([]GetLangItem, error) {
	items, err := i.repo.GetAllLang()
	if err != nil {
		return nil, nil
	}
	var res []GetLangItem
	for _, v := range items {
		res = append(res, GetLangItem{
			LangId:   v.Id,
			LangCode: v.Code,
			LangName: v.Name,
		})
	}
	return res, nil
}
