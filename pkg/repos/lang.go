package repos

import (
	"context"
	"github.com/abdybaevae/socialnet/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type LangRepo interface {
	CreateLang(ctx context.Context, lang entities.LanguageEntity) error
	GetAllLang(ctx context.Context) (items []entities.LanguageEntity, err error)
}
type impl struct {
	db *sqlx.DB
}

func NewLangRepo(db *sqlx.DB) LangRepo {
	return &impl{db}
}

func (s *impl) CreateLang(ctx context.Context, entity entities.LanguageEntity) error {
	_, err := s.db.NamedExecContext(ctx, `insert into lang(lang_name, lang_code) values (:lang_name, :lang_code)`, entity)
	return err
}

func (s *impl) GetAllLang(ctx context.Context) ([]entities.LanguageEntity, error) {
	var records []entities.LanguageEntity
	if err := s.db.SelectContext(ctx, &records, "select * from lang"); err != nil {
		return records, err
	}
	return records, nil
}
