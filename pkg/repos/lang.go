package repos

import "github.com/abdybaevae/socialnet/pkg/entities"

type LangRepo interface {
	CreateLang(lang *entities.LanguageEntity) error
	GetAllLang() (items []entities.LanguageEntity, err error)
}
type impl struct {
}

func NewLangRepo() LangRepo {
	return &impl{}
}
func (l *impl) CreateLang(args *entities.LanguageEntity) error {
	return nil
}
func (l *impl) GetAllLang() ([]entities.LanguageEntity, error) {
	return nil, nil
}
