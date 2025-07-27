package models

import "time"

type StockHistory struct {
	ID            int       `json:"id" gorm:"column:id"`
	ProductID     int       `json:"product_id" gorm:"column:product_id"`
	Qty           int       `json:"qty" gorm:"column:qty"`
	OperationType string    `json:"operation_type" gorm:"column:operation_type"`
	Operation     string    `json:"operation" gorm:"column:operation"`
	CreatedBy     int       `json:"created_by" gorm:"column:created_by"`
	CreatedAt     time.Time `json:"-" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"-" gorm:"column:updated_at"`
}
