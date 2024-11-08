package property

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrSomeThing = errors.New("some error")
)

type Repo interface {
	Insert(ctx context.Context, p *Property) error
	DeleteByID(ctx context.Context, id *uuid.UUID) error
	UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error
	// ...
}

type Property struct {
	ID uuid.UUID

	Title        string
	Description  string
	BuyPrice     uint64
	RentPrice    uint64
	RentPriceMin uint64
	RentPriceMax uint64
	RahnPriceMin uint64
	RahnPriceMax uint64
	Area         uint64
	Rooms        uint
	DealingType  string // buy, rent, rahn
	Type         string
	City         string
	District     string
	Address      string
	BuildYear    uint
	Floor        uint
	HasElevator  bool
	HasStorage   bool
	Latitude     float64
	Longitude    float64
	Source       string
	URL          string
	Images       []string
}
