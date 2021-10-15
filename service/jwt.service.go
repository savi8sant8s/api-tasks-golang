package service

import (
	"os"
	"time"
	"github.com/brianvoe/sjwt"
)

var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

type JwtService struct {}

func (this *JwtService) CreateToken(email string) string {
	claims := sjwt.New()
	claims.Set("userEmail", email)
	claims.SetNotBeforeAt(time.Now().Add(time.Minute * 15))

	return claims.Generate(JWT_SECRET_KEY)
}

func (this *JwtService) VerifyToken(token string) (bool, string) {
	valid := sjwt.Verify(token, JWT_SECRET_KEY)
	if !valid {
		return false, ""
	}
	claims, _ := sjwt.Parse(token)
	userEmail, _ := claims.GetStr("userEmail")
	return true, userEmail
}

