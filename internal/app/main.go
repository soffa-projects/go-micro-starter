package app

import (
	"app/config"
	"github.com/soffa-projects/go-micro/adapters"
	"github.com/soffa-projects/go-micro/micro"
)

func Create(features []micro.Feature) *micro.App {
	app := adapters.NewApp(
		"service",
		"0.0.0",
		micro.Cfg{
			Locales:     "fr",
			FS:          config.RootFS,
			MultiTenant: false,
		})
	initDI(app)
	return app.Init(features)
}

func initDI(_ *micro.App) {
	// Initialize dependencies
}
