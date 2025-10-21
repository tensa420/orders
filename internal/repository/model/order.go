package model

import (
	"errors"
	"time"
)

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrInternalError      = errors.New("internal error")
	ErrSuccessCancel      = errors.New("success cancel")
	ErrSomeDetailsMissing = errors.New("some details missing")
)

type Status string

const (
	StatusPaid           = "PAID"
	StatusCancelled      = "CANCELLED"
	StatusPendingPayment = "PENDING_PAYMENT"
)

type Category string

const (
	CategoryUnspecified Category = "UNSPECIFIED" // Категория не указана.
	CategoryEngine      Category = "ENGINE"      // Двигатели и компоненты.
	CategoryFuel        Category = "FUEL"        // Топливная система.
	CategoryPorthole    Category = "PORTHOLE"    // Иллюминаторы и окна.
	CategoryWing        Category = "WING"        // Крылья и аэродинамические поверхности.
)

type PaymentMethod string

const (
	PaymentMethodUnknown       PaymentMethod = "UNKNOWN"
	PaymentMethodCard          PaymentMethod = "CARD"
	PaymentMethodSBP           PaymentMethod = "PaymentMethodSBP"
	PaymentMethodCreditCard    PaymentMethod = "CREDIT_CARD"
	PaymentMethodInvestorMoney PaymentMethod = "INVESTOR_MONEY"
)

func PaymentMethodToString(paymentMethod PaymentMethod) string {
	switch paymentMethod {
	case PaymentMethodCard:
		return "Card"
	case PaymentMethodSBP:
		return "SBP"
	case PaymentMethodCreditCard:
		return "CreditCard"
	case PaymentMethodInvestorMoney:
		return "InvestorMoney"
	default:
		return "Unknown"
	}
}

type Order struct {
	OrderUUID       string
	UserUUID        string
	PartsUUID       []string
	TotalPrice      float64
	TransactionUUID *string
	PaymentMethod   *string
	Status          string
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
