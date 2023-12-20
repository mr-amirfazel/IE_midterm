package models

import (
	"time"

	"gorm.io/datatypes"
)

type Basket struct {
    ID        uint `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Data      string `gorm:"size:2048"`
    State     string `gorm:"size:10"`
    UserID   uint
}

type BasketDTO struct {
    ID        uint `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Data      string `gorm:"size:2048"`
    State     string `gorm:"size:10"`
    JSONData  datatypes.JSON `gorm:"column:data"`
}

var BasketJson []datatypes.JSON
