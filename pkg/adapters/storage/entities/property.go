package entities

import (
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	Title       string
	Description string
	Price       float64
	MinRent     float64
	MaxRent     float64
	Area        float64
	Rooms       int
	Type        string
	City        string
	District    string
	Address     string
	BuildYear   int
	Floor       int
	HasElevator bool
	HasStorage  bool
	Latitude    float64
	Longitude   float64
	Source      string
	URL         string
	Images      []string // ?
}
