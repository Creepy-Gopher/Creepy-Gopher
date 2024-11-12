package service

import (
	"creepy/internal/storage/postgis"
	"creepy/pkg/config"
	"log"

	"gorm.io/gorm"
)

type AppContainer struct {
	Cfg                 config.Config
	DbConn              *gorm.DB
	propertyService		*PropertyService
	userService			*UserService
	filterService 		*FilterService
}

func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		Cfg: cfg,
	}

	app.mustInitDB()

	app.setPropertyService()
	app.setUserService()
	app.setFilterService()
	return app, nil
}

func (a *AppContainer) RawDBConnection() *gorm.DB {
	return a.DbConn
}

func (a *AppContainer) mustInitDB() {
	if a.DbConn != nil {
		return
	}

	db, err := postgis.NewPostgresGormConnection(a.Cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	a.DbConn = db

	err = postgis.AddExtension(a.DbConn)
	if err != nil {
		log.Fatal("Create extension failed: ", err)
	}

	err = postgis.Migrate(a.DbConn)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
}


func (a *AppContainer) PropertyService() *PropertyService {
	return a.propertyService
}

func (a *AppContainer) setPropertyService() {
	if a.propertyService != nil {
		return
	}
	a.propertyService = NewPropertyService(postgis.NewPropertyRepository(a.DbConn),
	)
}

func (a *AppContainer) UserService() *UserService {
	return a.userService
}

func (a *AppContainer) setUserService() {
	if a.userService != nil {
		return
	}
	a.userService = NewUserService(
		postgis.NewUserRepo(a.DbConn),
	)
}

func (a *AppContainer) FilterService() *FilterService {
	return a.filterService
}

func (a *AppContainer) setFilterService(){
	if a.filterService != nil {
		return
	}
	a.filterService = NewFilterService(
		postgis.NewFilterRepo(a.DbConn),
	)
}
