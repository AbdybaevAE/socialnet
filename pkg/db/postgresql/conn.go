package postgresql

import (
	"context"
	"github.com/jmoiron/sqlx"
)

const connection = "postgresql://localhost:5432/socialnet"

func Connect(ctx context.Context) (*sqlx.DB, error) {
	Db, err := sqlx.ConnectContext(ctx, "postgres", "user=cifer dbname=socialnet sslmode=disable")
	if err != nil {
		return nil, err
	}
	return Db, nil
}
