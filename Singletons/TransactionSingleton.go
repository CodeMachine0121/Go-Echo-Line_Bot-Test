package singletons

import (
	"go-line/Models"
	"sync"
)

var once sync.Once
var transactionHistory *Models.TransactionHistory

func GetTransactionSingleton() *Models.TransactionHistory {

	once.Do(func() {
		transactionHistory = new(Models.TransactionHistory)
	})
	return transactionHistory
}
