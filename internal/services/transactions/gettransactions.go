package services

import (
	"nequi.com/poc-sqs-listener/internal/repository"
	"nequi.com/poc-sqs-listener/internal/domain"
)

type GetTransactionsService struct {
	transactionsRepository transactionsrepository.TransactionsRepository
}

func NewGetTransactionsService(transactionsRepositoryIn transactionsrepository.TransactionsRepository) GetTransactionsService {
	return GetTransactionsService{
		transactionsRepository: transactionsRepositoryIn,
	}
}

func (h GetTransactionsService) GetTransactions() (transaction.Transaction, error) {
	//Acá va la logica de mi negocio
	return h.transactionsRepository.GetTransactions()
}