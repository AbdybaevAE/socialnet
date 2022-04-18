package entities

type CountryEntity struct {
	Id   int    `db:"country_id"`
	Code string `db:"country_code"`
}
