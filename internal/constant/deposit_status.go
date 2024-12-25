package constant

type DepositStatus string

const (
	DepositStatusUnpaid DepositStatus = "UNPAID"
	DepositStatusPaid   DepositStatus = "PAID"

	DepositStatusError DepositStatus = "ERROR"
	// DepositStatusDelayed Only displayed on landing only
	DepositStatusDelayed DepositStatus = "DELAYED"

	DepositStatusExpired DepositStatus = "EXPIRED"
	DepositStatusSuccess DepositStatus = "SUCCESS"
)

// This will check validity update status
var validTransition = map[DepositStatus][]DepositStatus{
	DepositStatusUnpaid: {DepositStatusPaid, DepositStatusExpired},
	DepositStatusPaid:   {DepositStatusSuccess, DepositStatusError},
	DepositStatusError:  {DepositStatusError, DepositStatusPaid, DepositStatusSuccess},
}

func GetDepositStatusActions(currentStatus *DepositStatus) []DepositStatus {
	if currentStatus != nil {
		return validTransition[*currentStatus]
	}
	return []DepositStatus{
		DepositStatusPaid,
		DepositStatusError,
		DepositStatusSuccess,
	}
}

func (status DepositStatus) CheckEligibleStatus(nextStatus DepositStatus) bool {
	allowedNextStatuses, ok := validTransition[status]
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
