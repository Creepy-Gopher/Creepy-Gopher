package models

type Property struct {
	Model
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
	URL          string
	Images       []string `gorm:"type:text[]"`
}
