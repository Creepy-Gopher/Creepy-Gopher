package models

import "github.com/google/uuid"

type Property struct {
	Model
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
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
	HasParking   bool
	Latitude     float64
	Longitude    float64
	Source       string
	URL          string `gorm:"uniqueIndex"`
	Image        string
}
