package helpers

import (
	"log"
)

func ReurnError(err error, msg string) error {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
	return err
}
