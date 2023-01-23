package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}
