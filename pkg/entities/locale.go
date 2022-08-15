package entities

import "encoding/json"

type LocaleEntity struct {
	Id              int    `db:"locale_id"`
	TranslationJson []byte `db:"locale_translations"`
	Translations    map[string]string
}

func (e *LocaleEntity) InitTranslation() error {
	return json.Unmarshal(e.TranslationJson, &e.Translations)
}
