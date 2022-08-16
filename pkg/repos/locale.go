package repos

import (
	"context"
	"encoding/json"
	"github.com/abdybaevae/socialnet/pkg/entities"
	"github.com/abdybaevae/socialnet/pkg/lib/comerrors"
	"github.com/jmoiron/sqlx"
)

type LocaleRepo interface {
	AddLocale(ctx context.Context, entity entities.LocaleEntity) (localeId int64, err error)
	GetLocale(ctx context.Context, id int64) (entity entities.LocaleEntity, err error)
}

type repository struct {
	Db *sqlx.DB
}

const insertQuery = `insert into locale (locale_translations) values ($1) returning locale_id`

func (r *repository) AddLocale(ctx context.Context, entity entities.LocaleEntity) (int64, error) {
	jsonBytes, err := json.Marshal(entity.Translations)
	if err != nil {
		return 0, err
	}
	res, err := r.Db.QueryContext(ctx, insertQuery, jsonBytes)
	var id int64
	if !res.Next() {
		return 0, comerrors.ErrInternalError
	}
	if err := res.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

const selectQuery = "select locale_translations from locale where locale_id = $1"

func (r *repository) GetLocale(ctx context.Context, id int64) (entities.LocaleEntity, error) {
	var entity entities.LocaleEntity
	if err := r.Db.GetContext(ctx, &entity, selectQuery, id); err != nil {
		return entity, err
	}
	if err := entity.InitTranslation(); err != nil {
		return entity, err
	}
	return entity, nil
}

func NewLocaleRepo(Db *sqlx.DB) LocaleRepo {
	r := &repository{
		Db: Db,
	}
	return r
}
