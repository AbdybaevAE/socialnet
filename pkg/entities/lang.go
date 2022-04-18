package entities

type LanguageEntity struct {
	Id   int    `db:"lang_id"`
	Name string `db:"lang_name"`
	Code string `db:"lang_code"`
}
