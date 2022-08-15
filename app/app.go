package app

import (
	"context"
	"log"

	"github.com/abdybaevae/socialnet/pkg/Db/postgresql"
)

type App interface {
	Run()
}
type appImpl struct {
}

func (a *appImpl) Run() {
	ctx := context.Background()
	Db, err := postgresql.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()
	log.Println("Running app...")
}

func NewApp() App {
	return &appImpl{}
}
