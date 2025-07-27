package models

import "time"

type Product struct {
	ID           int       `json:"id" gorm:"column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	Description  *string   `json:"description" gorm:"column:description"`
	Sku          string    `json:"sku" gorm:"column:sku"`
	LiveQty      int       `json:"live_qty" gorm:"column:live_qty"`
	WarehouseQty int       `json:"warehouse_qty" gorm:"column:warehouse_qty"`
	CreatedBy    int       `json:"created_by" gorm:"column:created_by"`
	CreatedAt    time.Time `json:"-" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"-" gorm:"column:updated_at"`
}
