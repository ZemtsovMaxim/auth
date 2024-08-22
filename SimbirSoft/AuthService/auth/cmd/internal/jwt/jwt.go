package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	models "gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/models"
)

type MyClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

// NewToken creates new JWT token for given user.
func NewToken(user models.User, secret models.Secret, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(secret.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(accessToken string, secret models.Secret) (payload *MyClaims, err error) {

	claims := &MyClaims{}

	parsedToken, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
