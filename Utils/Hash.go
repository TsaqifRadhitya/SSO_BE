package Utils

import "golang.org/x/crypto/bcrypt"

func CreateHash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(bytes), err
}

func CompareHash(str string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err
}
