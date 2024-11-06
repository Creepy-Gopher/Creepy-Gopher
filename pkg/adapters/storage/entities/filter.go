package entities

import (
	"gorm.io/gorm"
)

type FilterSet struct {
	gorm.Model
	PriceRangeMin     float64 `gorm:"uniqueIndex:idx_filterset_all"`
	PriceRangeMax     float64 `gorm:"uniqueIndex:idx_filterset_all"`
	AreaRangeMin      float64 `gorm:"uniqueIndex:idx_filterset_all"`
	AreaRangeMax      float64 `gorm:"uniqueIndex:idx_filterset_all"`
	RoomMin           int     `gorm:"uniqueIndex:idx_filterset_all"`
	RoomMax           int     `gorm:"uniqueIndex:idx_filterset_all"`
	Cities            string  `gorm:"uniqueIndex:idx_filterset_all"`
	Districts         string  `gorm:"uniqueIndex:idx_filterset_all"`
	PropertyType      string  `gorm:"uniqueIndex:idx_filterset_all"`
	DealingType       string  `gorm:"uniqueIndex:idx_filterset_all"`
	FloorRangeMin     int     `gorm:"uniqueIndex:idx_filterset_all"`
	FloorRangeMax     int     `gorm:"uniqueIndex:idx_filterset_all"`
	HasElevator       bool    `gorm:"uniqueIndex:idx_filterset_all"`
	HasStorage        bool    `gorm:"uniqueIndex:idx_filterset_all"`
	BuildYearRangeMin int     `gorm:"uniqueIndex:idx_filterset_all"`
	BuildYearRangeMax int     `gorm:"uniqueIndex:idx_filterset_all"`
	LocationLatitude  float64 `gorm:"uniqueIndex:idx_filterset_all"`
	LocationLongitude float64 `gorm:"uniqueIndex:idx_filterset_all"`
	LocationRadius    float64 `gorm:"uniqueIndex:idx_filterset_all"`
	CreatedAfter      string  `gorm:"uniqueIndex:idx_filterset_all"`
	Sources           string  `gorm:"uniqueIndex:idx_filterset_all"`
}
