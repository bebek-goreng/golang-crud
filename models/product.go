package models

import (
	"time"
)

type Product struct {
	Id          int64      `gorm:"primaryKey" json:"id"`
	ProductName string     `gorm:"type:varchar(300); not null" json:"product_name"`
	Price       float64    `gorm:"type:decimal(10,2); not null" json:"price"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
