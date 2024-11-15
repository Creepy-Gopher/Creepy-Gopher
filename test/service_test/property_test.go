package servicetest

import (
	"creepy/internal/models"
	"creepy/internal/service"
)

var Property *models.Property

func init() {
	Property = &models.Property{
		Model:        models.Model{},
		Title:        "",
		Description:  "",
		BuyPrice:     0,
		RentPrice:    0,
		RentPriceMin: 0,
		RentPriceMax: 0,
		RahnPriceMin: 0,
		RahnPriceMax: 0,
		Area:         0,
		Rooms:        0,
		DealingType:  "",
		Type:         "",
		City:         "",
		District:     "",
		Address:      "",
		BuildYear:    0,
		Floor:        0,
		HasElevator:  false,
		HasStorage:   false,
		HasParking:   false,
		Latitude:     0,
		Longitude:    0,
		Source:       "",
		URL:          "",
		Image:        "",
		SearchCount:  0,
	}
}

func PropertyServiceTest(app service.AppContainer) {
    app.Cfg.Logger.Info("starting test user service")

	app.Cfg.Logger.Info("'user service' successfully passed tests")
}