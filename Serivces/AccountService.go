package serivces

import (
	"go-line/Models"
	singletons "go-line/Singletons"
)

var _transactionHistory = singletons.GetTransactionSingleton()

func InsertTransaction(tx *Models.SigleTransaction) int {

	_transactionHistory.History = append(_transactionHistory.History, *tx)

	_transactionHistory.Totals = CalculateTotals()

	return _transactionHistory.Totals
}

func CalculateTotals() int {
	sum := 0
	for i := 0; i < len(_transactionHistory.History); i++ {
		sum += _transactionHistory.History[i].Amount
	}
	return sum
}
