package Utils

import "log"

func ErrorHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
