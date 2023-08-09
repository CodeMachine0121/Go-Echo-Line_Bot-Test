package Models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionHistory struct {
	History []SigleTransaction
	Totals  int
}

type SigleTransaction struct {
	Id          uuid.UUID
	Item        string
	Amount      int
	CreatedTime time.Time
}
