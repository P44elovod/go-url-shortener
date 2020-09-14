package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(url []byte) (res []byte, err error) {
	res, err = bcrypt.GenerateFromPassword(url, bcrypt.MinCost)

	return

}
