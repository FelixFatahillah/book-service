package dtos

import "time"

type LoanDto struct {
	CustomerName       string    `json:"customer_name"`
	LoanDate           time.Time `json:"loan_date"`
	ReturnDateSchedule time.Time `json:"return_date_schedule"`
	ReturnDate         time.Time `json:"return_date"`
	BookID             string    `json:"book_id"`
}

type ReturnDto struct {
	LoanID string `json:"loan_id"`
}
