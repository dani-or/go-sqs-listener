package transactionsrepository

import (
	"nequi.com/poc-sqs-listener/internal/domain"
)

type TransactionsRepository interface {
	GetTransactions() (transaction.Transaction, error)
}