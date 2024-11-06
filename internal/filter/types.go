package filter

import (
	"context"
	"errors"
)

var (
	ErrNegativePrice  = errors.New("inserted price can not be negative")
	ErrCreateFilter = errors.New("filter could not be created")
)

type Repo interface {
	Insert(ctx context.Context, f *FilterSet) error
	//delete
	// update

}

type FilterSet struct {
	ID uint

	PriceRangeMin     float64
	PriceRangeMax     float64
	AreaRangeMin      float64
	AreaRangeMax      float64
	RoomMin           int
	RoomMax           int
	Cities            string
	Districts         string
	PropertyType      string
	DealingType       string
	FloorRangeMin     int
	FloorRangeMax     int
	HasElevator       bool
	HasStorage        bool
	BuildYearRangeMin int
	BuildYearRangeMax int
	LocationLatitude  float64
	LocationLongitude float64
	LocationRadius    float64
	CreatedAfter      string
	Sources           string
}
