package models

type Filter struct {
	Model
	BuyPriceMin       uint64   `gorm:"uniqueIndex:idx_filterset_all"`
	BuyPriceMax       uint64   `gorm:"uniqueIndex:idx_filterset_all"`
	RentPriceMin      uint64   `gorm:"uniqueIndex:idx_filterset_all"`
	RentPriceMax      uint64   `gorm:"uniqueIndex:idx_filterset_all"`
	AreaMin           uint     `gorm:"uniqueIndex:idx_filterset_all"`
	AreaMax           uint     `gorm:"uniqueIndex:idx_filterset_all"`
	RoomMin           uint     `gorm:"uniqueIndex:idx_filterset_all"`
	RoomMax           uint     `gorm:"uniqueIndex:idx_filterset_all"`
	FloorMin          uint     `gorm:"uniqueIndex:idx_filterset_all"`
	FloorMax          uint     `gorm:"uniqueIndex:idx_filterset_all"`
	BuildYearMin      uint     `gorm:"uniqueIndex:idx_filterset_all"`
	BuildYearMax      uint     `gorm:"uniqueIndex:idx_filterset_all"`
	PropertyType      string   `gorm:"uniqueIndex:idx_filterset_all"`
	DealingType       string   `gorm:"uniqueIndex:idx_filterset_all"`
	HasElevator       bool     `gorm:"uniqueIndex:idx_filterset_all"`
	HasStorage        bool     `gorm:"uniqueIndex:idx_filterset_all"`
	HasParking        bool     `gorm:"uniqueIndex:idx_filterset_all"`
	LocationLatitude  float64  `gorm:"uniqueIndex:idx_filterset_all"`
	LocationLongitude float64  `gorm:"uniqueIndex:idx_filterset_all"`
	LocationRadius    float64  `gorm:"uniqueIndex:idx_filterset_all"`
	CreatedAfter      string   `gorm:"uniqueIndex:idx_filterset_all"`
	City              string   `gorm:"uniqueIndex:idx_filterset_all"` 
	District          string   `gorm:"uniqueIndex:idx_filterset_all"`
	Source            string   `gorm:"uniqueIndex:idx_filterset_all"`
	SearchCount	      uint     // for counting top searched filters
}
