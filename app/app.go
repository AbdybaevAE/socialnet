package app

import "log"

type App interface {
	Run()
}
type appImpl struct {
}

func (a *appImpl) Run() {
	log.Println("Running app...")
}

func NewApp() App {
	return &appImpl{}
}
