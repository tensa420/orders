package entity

type Status string

const (
	StatusPaid           = "PAID"
	StatusCancelled      = "CANCELLED"
	StatusPendingPayment = "PENDING_PAYMENT"
)
