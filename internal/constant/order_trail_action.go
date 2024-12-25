package constant

type OrderTrailAction string

const (
	ActionTrailPaymentApproved OrderTrailAction = "Payment Approved"
	ActionTrailPaymentExpired  OrderTrailAction = "Payment Expired"
	ActionTrailDistribution    OrderTrailAction = "Distribution SKU's"
	ActionTrailManualUpdate    OrderTrailAction = "Manual Update"
)
