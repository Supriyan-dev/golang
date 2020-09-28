package helpers

import (
	"golang.org/x/crypto/bcrypt"
	// "crypto/md5"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// func HashPasswordMd5(password string) (string, error) {
// 	data := []byte(password)
// 	var pass [16]byte
// 	pass = md5.Sum(data)
// 	return string(pass[:]), nil
// }

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false, err
	}

	return true, nil
}
