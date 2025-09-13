package Utils

import "golang.org/x/crypto/bcrypt"

func CreateHash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(bytes), err
}

func CompareHash(str string, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str)); err != nil {
		return false
	}
	return true
}
