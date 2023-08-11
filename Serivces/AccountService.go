package serivces

import (
	"go-line/Models"
	singletons "go-line/Singletons"
)

func InsertTransaction(tx *Models.SigleTransaction) int {

	_transactionHistory := singletons.GetTransactionSingleton()
	_transactionHistory.History = append(_transactionHistory.History, tx)

	_transactionHistory.Totals = CalculateTotals()

	return _transactionHistory.Totals
}

func CalculateTotals() int {

	_transactionHistory := singletons.GetTransactionSingleton()
	sum := 0
	for i := 0; i < len(_transactionHistory.History); i++ {
		sum += _transactionHistory.History[i].Amount
	}
	return sum
}
