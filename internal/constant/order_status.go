package constant

type OrderStatus string

const (
	OrderStatusUnpaid OrderStatus = "UNPAID"
	OrderStatusPaid   OrderStatus = "PAID"

	OrderStatusPending OrderStatus = "PENDING"
	OrderStatusError   OrderStatus = "ERROR"
	// OrderStatusDelayed Only displayed on landing only
	OrderStatusDelayed OrderStatus = "DELAYED"

	OrderStatusExpired  OrderStatus = "EXPIRED"
	OrderStatusCanceled OrderStatus = "CANCELED"

	OrderStatusSuccess OrderStatus = "SUCCESS"

	OrderStatusPreRefund OrderStatus = "PRE_REFUNDED"
	OrderStatusRefunded  OrderStatus = "REFUNDED"
)

// This will check validity update status
var validTransitions = map[OrderStatus][]OrderStatus{
	OrderStatusUnpaid:    {OrderStatusPaid, OrderStatusExpired, OrderStatusCanceled},
	OrderStatusPaid:      {OrderStatusPending, OrderStatusError},
	OrderStatusPending:   {OrderStatusError, OrderStatusSuccess, OrderStatusPreRefund, OrderStatusRefunded},
	OrderStatusSuccess:   {OrderStatusPreRefund},
	OrderStatusPreRefund: {OrderStatusRefunded, OrderStatusSuccess},
	OrderStatusError:     {OrderStatusError, OrderStatusPending, OrderStatusPreRefund, OrderStatusSuccess},
}

func GetOrderStatusActions(currentStatus *OrderStatus) []OrderStatus {
	if currentStatus != nil {
		return validTransitions[*currentStatus]
	}
	return []OrderStatus{
		OrderStatusPaid,
		OrderStatusError,
		OrderStatusPending,
		OrderStatusCanceled,
		OrderStatusSuccess,
		OrderStatusPreRefund,
		OrderStatusRefunded,
	}
}

func (status OrderStatus) CheckEligibleStatus(nextStatus OrderStatus) bool {
	allowedNextStatuses, ok := validTransitions[status]
	if !ok {
		return false
	}

	for _, allowedStatus := range allowedNextStatuses {
		if allowedStatus == nextStatus {
			return true
		}
	}

	return false
}
