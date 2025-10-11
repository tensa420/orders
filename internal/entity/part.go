package entity

import "time"

type Part struct {
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      int8
	Dimensions    Dimensions
	Manufacturer  Manufacturer
	Tags          []string
	MetaData      map[string]any
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
