package locale

import (
	"context"
	"github.com/abdybaevae/socialnet/pkg/entities"
	"github.com/abdybaevae/socialnet/pkg/lib/comerrors"
	"github.com/abdybaevae/socialnet/pkg/repos"
	"github.com/abdybaevae/socialnet/pkg/services/admin/lang"
	"strings"
)

type AddLocale struct {
	Translations map[string]string
}

type Service interface {
	AddLocale(ctx context.Context, args AddLocale) (localeId int64, err error)
	GetLocale(ctx context.Context, id int64) (entities.LocaleEntity, error)
}
type impl struct {
	langService lang.Service
	localeRepo  repos.LocaleRepo
}

func NewService(langService lang.Service, localeRepo repos.LocaleRepo) Service {
	return &impl{
		langService,
		localeRepo,
	}
}

func (s *impl) AddLocale(ctx context.Context, args AddLocale) (int64, error) {
	// retrieve all languages
	allLangs, err := s.langService.GetLangList(ctx)
	if err != nil {
		return 0, err
	}

	// trim all localeTranslation values and skip empty ones
	givenTranslations := map[string]string{}
	for key, value := range args.Translations {
		trimKey := strings.Trim(key, " ")
		trimValue := strings.Trim(value, " ")
		if len(trimKey) != 0 && len(trimValue) != 0 {
			givenTranslations[trimKey] = trimValue
		}
	}

	// ensure that localeTranslation exist for all languages
	for _, currLang := range allLangs {
		if _, ok := givenTranslations[currLang.LangCode]; !ok {
			return 0, comerrors.ErrBadArguments
		}
	}

	// create and save entity
	var localeEntity = entities.LocaleEntity{
		Translations: givenTranslations,
	}
	return s.localeRepo.AddLocale(ctx, localeEntity)
}

func (s *impl) GetLocale(ctx context.Context, id int64) (entities.LocaleEntity, error) {
	return s.localeRepo.GetLocale(ctx, id)
}
