package helpers

import (
	"fmt"
	"log"
)

func LogErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
