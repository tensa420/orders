package entity

type OrderPaid struct {
	EventUUID       string  `json:"event_uuid"`
	OrderUUID       string  `json:"order_uuid"`
	UserUUID        string  `json:"user_uuid"`
	PaymentMethod   *string `json:"payment_method"`
	TransactionUUID string  `json:"transaction_uuid"`
}
