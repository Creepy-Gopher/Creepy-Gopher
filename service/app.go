package service

import (
	"creepy/config"
	"creepy/internal/filter"
	"creepy/pkg/adapters/storage"
	"log"

	"gorm.io/gorm"
)

type AppContainer struct {
	cfg           config.Config
	dbConn        *gorm.DB
	filterService *FilterService
}

func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		cfg: cfg,
	}
	app.mustInitDB()
	storage.Migrate(app.dbConn)

	app.setFilterService()

	return app, nil
}

func (a *AppContainer) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	db, err := storage.NewMysqlGormConnection(a.cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	a.dbConn = db
}

func (a *AppContainer) setFilterService() {
	if a.filterService != nil {
		return
	}
	a.filterService = NewFilterService(filter.NewOps(storage.NewFilterSetRepo(a.dbConn)))
}
