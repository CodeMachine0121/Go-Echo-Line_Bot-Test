package Models

type TransactionHistory struct {
	History []SigleTransaction
	Totals  int
}

type SigleTransaction struct {
	Id     int
	Item   string
	Amount int
}
