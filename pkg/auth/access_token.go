package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

type JWTPayload struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func CreateToken(payload JWTPayload) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	claims := jwt.MapClaims{
		"exp":       jwt.NewNumericDate(expireTime),
		"iat":       jwt.NewNumericDate(nowTime),
		"sub":       payload.UserID,
		"fist_name": payload.FirstName,
		"last_name": payload.LastName,
		"email":     payload.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(viper.GetString("JWT_SECRET"))
	return token.SignedString(secret)
}

func ParseToken(accessToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error get user claims from token")
	}

	return claims, nil
}

func TokenTrimmer(token string) string {
	const prefix = "Bearer "
	if len(token) > len(prefix) && token[:len(prefix)] == prefix {
		token = token[len(prefix):]
	}

	return token
}
