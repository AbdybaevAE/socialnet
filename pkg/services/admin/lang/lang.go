package lang

import (
	"context"
	"github.com/abdybaevae/socialnet/pkg/entities"
	"github.com/abdybaevae/socialnet/pkg/lib/comerrors"
	"github.com/abdybaevae/socialnet/pkg/repos"
)

type CreateLangArgs struct {
	Code string
	Name string
}
type GetLangResultItem struct {
	LangId   int
	LangCode string
	LangName string
}
type Service interface {
	CreateLang(ctx context.Context, args CreateLangArgs) error
	GetLangList(ctx context.Context) (res []GetLangResultItem, err error)
}

func NewService(repo repos.LangRepo) Service {
	return &service{repo}
}

type service struct {
	repo repos.LangRepo
}

func (s *service) CreateLang(ctx context.Context, args CreateLangArgs) error {
	if len(args.Code) == 0 || len(args.Name) == 0 {
		return comerrors.ErrBadArguments
	}
	langEntity := entities.LanguageEntity{
		Name: args.Name,
		Code: args.Code,
	}
	if err := s.repo.CreateLang(ctx, langEntity); err != nil {
		return err
	}
	return nil
}
func (s *service) GetLangList(ctx context.Context) ([]GetLangResultItem, error) {
	items, err := s.repo.GetAllLang(ctx)
	if err != nil {
		return nil, nil
	}
	var res []GetLangResultItem
	for _, v := range items {
		res = append(res, GetLangResultItem{
			LangId:   v.Id,
			LangCode: v.Code,
			LangName: v.Name,
		})
	}
	return res, nil
}
