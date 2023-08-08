package Util

import "golang.org/x/crypto/bcrypt"

// bcrpt加密与验证
func PwdHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil

}
func PwdVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil

}
