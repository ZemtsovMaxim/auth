package jwt

import "github.com/golang-jwt/jwt/v5"

func IsValid(tokenStr string) bool {
	jwtKey := []byte("test-secret")

	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}
