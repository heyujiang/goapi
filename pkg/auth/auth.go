package auth

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//加密密码
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

//校验密码是否正确
func Compare(hashedPassword, password string) error {
	fmt.Println(hashedPassword)
	fmt.Println(password)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
