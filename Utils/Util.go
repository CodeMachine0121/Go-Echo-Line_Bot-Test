package Utils

import (
	"log"
	"time"
)

func ErrorHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetTimeWithFormat(t time.Time) string {
	return t.Format("2006-01-02")
}
