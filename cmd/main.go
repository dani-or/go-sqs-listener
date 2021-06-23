package main

import (
	"fmt"
	//"log"
	"nequi.com/poc-sqs-listener/internal/platform/storage"
	"nequi.com/poc-sqs-listener/internal/services/transactions"
)

//Ejemplo https://github.com/h2ik/go-sqs-poller/blob/v1.0.0/worker/worker.go#L54
func main() {
	fmt.Println("Hello, World! Daniela sqs listener")
	transactionsrepository := sqs.NewS3Repository()
	getTransactionsService := services.NewGetTransactionsService(transactionsrepository)
	for {
		fmt.Println("Entró")
		getTransactionsService.GetTransactions()
		fmt.Println("Salió")

	}
}