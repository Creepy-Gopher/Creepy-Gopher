package service

import (
	"creepy/internal/storage/postgis"
	"creepy/pkg/config"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type AppContainer struct {
	cfg                 config.Config
	dbConn              *gorm.DB
	propertyService		*PropertyService
	userService			*UserService
	filterService 		*FilterService
}

func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		cfg: cfg,
	}

	app.mustInitDB()

	if err := app.setPropertyService(); err != nil {
		app.cfg.Logger.Error(err.Error())
	}
	if err := app.setUserService(); err != nil {
		app.cfg.Logger.Error(err.Error())
	}
	if err := app.setFilterService(); err != nil {
		app.cfg.Logger.Error(err.Error())
	}
	return app, nil
}

func (a *AppContainer) RawDBConnection() *gorm.DB {
	return a.dbConn
}

func (a *AppContainer) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	db, err := postgis.NewPostgresGormConnection(a.cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	a.dbConn = db

	err = postgis.AddExtension(a.dbConn)
	if err != nil {
		log.Fatal("Create extension failed: ", err)
	}

	err = postgis.Migrate(a.dbConn)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
}


func (a *AppContainer) PropertyService() *PropertyService {
	return a.propertyService
}

func (a *AppContainer) setPropertyService() error {
	if a.propertyService != nil {
		return fmt.Errorf("application property service already exist")
	}
	a.propertyService = NewPropertyService(
		postgis.NewPropertyRepository(a.dbConn),
	)
	return nil
}

func (a *AppContainer) UserService() *UserService {
	return a.userService
}

func (a *AppContainer) setUserService() error {
	if a.userService != nil {
		return fmt.Errorf("application user service already exist")
	}
	a.userService = NewUserService(
		postgis.NewUserRepo(a.dbConn),
	)
	return nil
}

func (a *AppContainer) FilterService() *FilterService {
	return a.filterService
}

func (a *AppContainer) setFilterService() error {
	if a.filterService != nil {
		return fmt.Errorf("application filter service already exist")
	}
	a.filterService = NewFilterService(
		postgis.NewFilterRepo(a.dbConn),
	)
	return nil
}
