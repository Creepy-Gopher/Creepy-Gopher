package mappers

import (
	"creepy/internal/property"
	"creepy/pkg/adapters/storage/entities"
)

func PropertyEntityToDomain(p *entities.Property) *property.Property {
	return &property.Property{
		ID:           p.ID,
		Title:        p.Title,
		Description:  p.Description,
		BuyPrice:     p.BuyPrice,
		RentPrice:    p.RentPrice,
		RentPriceMin: p.RentPriceMin,
		RentPriceMax: p.RentPriceMax,
		RahnPriceMin: p.RahnPriceMin,
		RahnPriceMax: p.RahnPriceMax,
		Area:         p.Area,
		Rooms:        p.Rooms,
		DealingType:  p.DealingType,
		Type:         p.Type,
		City:         p.City,
		District:     p.District,
		Address:      p.Address,
		BuildYear:    p.BuildYear,
		Floor:        p.Floor,
		HasElevator:  p.HasElevator,
		HasStorage:   p.HasStorage,
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		Source:       p.Source,
		URL:          p.URL,
		Images:       p.Images,
	}
}

func PropertyDomainToEntity(p *property.Property) *entities.Property {
	return &entities.Property{
		Model:        entities.Model{ID: p.ID},
		Title:        p.Title,
		Description:  p.Description,
		BuyPrice:     p.BuyPrice,
		RentPrice:    p.RentPrice,
		RentPriceMin: p.RentPriceMin,
		RentPriceMax: p.RentPriceMax,
		RahnPriceMin: p.RahnPriceMin,
		RahnPriceMax: p.RahnPriceMax,
		Area:         p.Area,
		Rooms:        p.Rooms,
		DealingType:  p.DealingType,
		Type:         p.Type,
		City:         p.City,
		District:     p.District,
		Address:      p.Address,
		BuildYear:    p.BuildYear,
		Floor:        p.Floor,
		HasElevator:  p.HasElevator,
		HasStorage:   p.HasStorage,
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		Source:       p.Source,
		URL:          p.URL,
		Images:       p.Images,
	}
}
