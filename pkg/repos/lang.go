package repos

import "github.com/abdybaevae/socialnet/pkg/entities"

type LangRepo interface {
	CreateLang(lang *entities.LanguageEntity) error
	GetAllLang() (items []entities.LanguageEntity, err error)
}
type LangRepoImpl struct {
}

func NewLangRepo() LangRepo {
	return &LangRepoImpl{}
}
func (l *LangRepoImpl) CreateLang(args *entities.LanguageEntity) error {
	return nil
}
func (l *LangRepoImpl) GetAllLang() ([]entities.LanguageEntity, error) {
	return nil, nil
}
