package entity

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
