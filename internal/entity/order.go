package entity

type Order struct {
	OrderUUID       string
	UserUUID        string
	PartsUUID       []string
	TotalPrice      float64
	TransactionUUID *string
	PaymentMethod   *string
	Status          string
}
