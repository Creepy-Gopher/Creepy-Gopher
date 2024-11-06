package mappers

import (
	"creepy/internal/filter"
	"creepy/pkg/adapters/storage/entities"
)


func FilterEntityToDomainFilter(f entities.FilterSet) filter.FilterSet {
	return filter.FilterSet{
		ID:                f.ID,
		PriceRangeMin:     f.PriceRangeMax,
		PriceRangeMax:     f.PriceRangeMax,
		AreaRangeMin:      f.AreaRangeMin,
		// AreaRangeMax:      0,
		// RoomMin:           0,
		// RoomMax:           0,
		// Cities:            "",
		// Districts:         "",
		// PropertyType:      "",
		// DealingType:       "",
		// FloorRangeMin:     0,
		// FloorRangeMax:     0,
		// HasElevator:       false,
		// HasStorage:        false,
		// BuildYearRangeMin: 0,
		// BuildYearRangeMax: 0,
		// LocationLatitude:  0,
		// LocationLongitude: 0,
		// LocationRadius:    0,
		// CreatedAfter:      "",
		// Sources:           "",
	}
}

func FilterDomainToFilterEntity(f *filter.FilterSet) *entities.FilterSet{
	return &entities.FilterSet{
		PriceRangeMin:     f.PriceRangeMin,
		// PriceRangeMax:     0,
		// AreaRangeMin:      0,
		// AreaRangeMax:      0,
		// RoomMin:           0,
		// RoomMax:           0,
		// Cities:            "",
		// Districts:         "",
		// PropertyType:      "",
		// DealingType:       "",
		// FloorRangeMin:     0,
		// FloorRangeMax:     0,
		// HasElevator:       false,
		// HasStorage:        false,
		// BuildYearRangeMin: 0,
		// BuildYearRangeMax: 0,
		// LocationLatitude:  0,
		// LocationLongitude: 0,
		// LocationRadius:    0,
		// CreatedAfter:      "",
		// Sources:           "",
	}
}