package models

import "gorm.io/gorm"

type Symbols struct {
	gorm.Model
	Symbol            string  `json:"symbol" gorm:"type:varchar(30); not null; uniqueIndex"`
	Name              string  `json:"name" gorm:"type:varchar(255); not null"`
	CirculatingSupply float32 `json:"circulating_supply" gorm:"type:DECIMAL(30,8); notnull"`
	TotalSupply       float32 `json:"total_supply" gorm:"type:DECIMAL(30,8); notnull"`
	MaxSupply         float32 `json:"max_supply" gorm:"type:DECIMAL(30,8); notnull"`
	LastPrice         float32 `json:"last_price" gorm:"type:DECIMAL(30,8); notnull"`
}
