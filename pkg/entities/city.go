package entities

type CityEntity struct {
	Id   int    `db:"city_id"`
	Code string `db:"city_code"`
}
