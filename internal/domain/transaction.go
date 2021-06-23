package transaction

import (
)

type Transaction struct {
	Value       int
	Debenture	string
	EndDate		int
	StartDate	int
	Status		int
}

func NewTransaction(value, status, endDate, startDate int, debenture string) (Transaction, error) {

	return Transaction{
		Value	:  	value,
		Status	:	status,
		Debenture:	debenture,
		EndDate	:	endDate,
		StartDate: startDate,
	}, nil
}