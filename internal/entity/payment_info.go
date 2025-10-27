package entity

type PaymentInfo struct {
	TransactionUUID string
	OrderUUID       string
	PaymentMethod   PaymentMethod
}
