package servicetest

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/service"
	"fmt"
	"testing"
)

var filter *models.Filter

func init() {
	filter = &models.Filter{
		Model:             models.Model{},
		BuyPriceMin:       0,
		BuyPriceMax:       0,
		RentPriceMin:      0,
		RentPriceMax:      0,
		AreaMin:           0,
		AreaMax:           0,
		RoomMin:           0,
		RoomMax:           0,
		FloorMin:          0,
		FloorMax:          0,
		BuildYearMin:      0,
		BuildYearMax:      0,
		PropertyType:      "",
		DealingType:       "",
		HasElevator:       false,
		HasStorage:        false,
		HasParking:        false,
		LocationLatitude:  0,
		LocationLongitude: 0,
		LocationRadius:    0,
		CreatedAfter:      "",
		City:              "",
		District:          "",
		Source:            "",
		SearchCount:       0,
	}
}

func FilterServiceTest(app *service.AppContainer) {
	app.Cfg.Logger.Info("starting test user service")
	// app.DbConn = app.DbConn.Debug()

	ctx := context.Background()

	dbFilter, err := app.FilterService().GetByFilter(ctx, filter)
	if err != nil {
		app.Cfg.Logger.Error(err.Error())
	}
	if dbFilter == nil {
		err := app.FilterService().CreateFilter(ctx, filter)
		if err != nil {
			app.Cfg.Logger.Fatal(err.Error())
		}
	}

	dbFilter, err = app.FilterService().GetByFilter(ctx, filter)
	if err != nil {
		app.Cfg.Logger.Fatal(err.Error())
	}
	id := dbFilter.ID
	fmt.Printf("\nfilter id: %v \n", id)

	err = app.FilterService().DeleteFilter(ctx, id)
	if err != nil {
		app.Cfg.Logger.Error(err.Error())
	}
	app.Cfg.Logger.Info(fmt.Sprintf("filter whit id: %v deleted", id))

	app.Cfg.Logger.Info("'filter service' successfully passed tests")
}


func TestFilterService(t *testing.T) {
	cfg := readConfig()
	cfg.Logger = NewDevelopLogger()
	app, err := service.NewAppContainer(cfg)
	if err != nil {
		app.Cfg.Logger.Fatal(err.Error())
	}
	FilterServiceTest(app)
}
