package hash

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	pw := []byte(password)
	hashedPw, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPw), nil
}

func ComparePassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
