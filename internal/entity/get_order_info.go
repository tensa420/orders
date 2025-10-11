package entity

type GetOrderInfo struct {
	OrderUUID       string
	UserUUID        string
	PartUuids       []string
	TotalPrice      float64
	TransactionUUID *string
	PaymentMethod   *string
	Status          string
}
