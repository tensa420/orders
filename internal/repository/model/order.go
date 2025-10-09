package model

import (
	"time"
)

type Status int8

const (
	PAID Status = iota
	PENDINGPAYMENT
	CANCELLED
)

type Category int8

const (
	CATEGORY_UNKNOWN Category = iota
	CATEGORY_ENGINE
	CATEGORY_FUEL
	CATEGORY_PORTHOLE
	CATEGORY_WING
)

type PaymentMethod int8

const (
	UNKNOWN PaymentMethod = iota
	CARD
	SBP
	CREDIT_CARD
	INVESTOR_MONEY
)

type Order struct {
	OrderUUID       string
	UserUUID        string
	PartsUUID       []string
	TotalPrice      float64
	TransactionUUID *string
	PaymentMethod   *string
	Status          Status
}

type PartsFilter struct {
	Uuids                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}

type Dimensions struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

type Manufacturer struct {
	Name    string
	Country string
	Website string
}

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

type Value struct {
	StringValue  string
	Int64Value   int64
	Float64Value float64
	BoolValue    bool
}
type GetOrderResponse struct {
	OrderUUID       string
	UserUUID        string
	PartUuids       []string
	TotalPrice      float64
	TransactionUUID *string
	PaymentMethod   *string
	Status          Status
}
