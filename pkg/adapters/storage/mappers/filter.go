package mappers

import (
	"creepy/internal/filter"
	"creepy/pkg/adapters/storage/entities"
)

func FilterEntityToDomain(f *entities.Filter) *filter.Filter {
	return &filter.Filter{
		ID:                f.ID,
		BuyPriceMin:       f.BuyPriceMin,
		BuyPriceMax:       f.BuyPriceMax,
		RentPriceMin:      f.RentPriceMin,
		RentPriceMax:      f.RentPriceMax,
		AreaMin:           f.AreaMin,
		AreaMax:           f.AreaMax,
		RoomMin:           f.RoomMin,
		RoomMax:           f.RoomMax,
		FloorMin:          f.FloorMin,
		FloorMax:          f.FloorMax,
		BuildYearMin:      f.BuildYearMin,
		BuildYearMax:      f.BuildYearMax,
		Cities:            f.Cities,
		Districts:         f.Districts,
		PropertyType:      f.PropertyType,
		DealingType:       f.DealingType,
		HasElevator:       f.HasElevator,
		HasStorage:        f.HasStorage,
		LocationLatitude:  f.LocationLatitude,
		LocationLongitude: f.LocationLongitude,
		LocationRadius:    f.LocationRadius,
		CreatedAfter:      f.CreatedAfter,
		Sources:           f.Sources,
	}
}

func FilterDomainToEntity(f *filter.Filter) *entities.Filter {
	return &entities.Filter{
		Model:             entities.Model{ID: f.ID},
		BuyPriceMin:       f.BuyPriceMin,
		BuyPriceMax:       f.BuyPriceMax,
		RentPriceMin:      f.RentPriceMin,
		RentPriceMax:      f.RentPriceMax,
		AreaMin:           f.AreaMin,
		AreaMax:           f.AreaMax,
		RoomMin:           f.RoomMin,
		RoomMax:           f.RoomMax,
		FloorMin:          f.FloorMin,
		FloorMax:          f.FloorMax,
		BuildYearMin:      f.BuildYearMin,
		BuildYearMax:      f.BuildYearMax,
		Cities:            f.Cities,
		Districts:         f.Districts,
		PropertyType:      f.PropertyType,
		DealingType:       f.DealingType,
		HasElevator:       f.HasElevator,
		HasStorage:        f.HasStorage,
		LocationLatitude:  f.LocationLatitude,
		LocationLongitude: f.LocationLongitude,
		LocationRadius:    f.LocationRadius,
		CreatedAfter:      f.CreatedAfter,
		Sources:           f.Sources,
	}
}
