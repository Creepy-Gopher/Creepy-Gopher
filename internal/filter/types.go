package filter

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNegativePrice = errors.New("inserted price can not be negative")
	ErrCreateFilter  = errors.New("filter could not be created")
)

type Repo interface {
	Insert(ctx context.Context, f *Filter) error
	DeleteByID(ctx context.Context, id *uuid.UUID) error
	UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error
	// ...
}

type Filter struct {
	ID uuid.UUID

	BuyPriceMin       uint64
	BuyPriceMax       uint64
	RentPriceMin      uint64
	RentPriceMax      uint64
	AreaMin           uint
	AreaMax           uint
	RoomMin           uint
	RoomMax           uint
	FloorMin          uint
	FloorMax          uint
	BuildYearMin      uint
	BuildYearMax      uint
	Cities            []string
	Districts         []string
	PropertyType      string
	DealingType       string
	HasElevator       bool
	HasStorage        bool
	LocationLatitude  float64
	LocationLongitude float64
	LocationRadius    float64
	CreatedAfter      string
	Sources           []string
}
