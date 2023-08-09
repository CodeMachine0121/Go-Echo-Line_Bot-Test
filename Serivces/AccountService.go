package serivces

import "go-line/Models"

var _transactionHistory Models.TransactionHistory

func Init() {
	_transactionHistory = Models.TransactionHistory{
		History: []Models.SigleTransaction{},
		Totals:  0,
	}
}

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
