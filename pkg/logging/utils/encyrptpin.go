package utils

import "golang.org/x/crypto/bcrypt"

func HashPin(pin string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pin), 14)
	return string(bytes), err
}

func CheckPinHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
