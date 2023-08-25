package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func NewToken(uid, email string, role string) *claims {
	return &claims{
		Id:    uid,
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "coffeshop",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 5)),
		},
	}
}

func (c *claims) Generate() (string, error) {
	secrets := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(secrets))
}

func VerifyToken(token string) (*claims, error) {
	secrets := os.Getenv("JWT_SECRET")
	data, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets), nil
	})

	if err != nil {
		return nil, err
	}
	claimData := data.Claims.(*claims)

	return claimData, nil
}
