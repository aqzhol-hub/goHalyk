package token

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	SecretKey       = "secretkey"
	Issuer          = "AuthService"
	ExpirationHours = 24
)

type JwtClaim struct {
	Username string
	jwt.StandardClaims
}

func GenerateToken(username string) (error, string) {

	ExpiresAt := time.Now().Local().Add(time.Hour * time.Duration(ExpirationHours)).Unix()

	claims := &JwtClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpiresAt,
			Issuer:    Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey))
	return err, signedToken
}

func ValidateToken(signedToken string) (error, *JwtClaim) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)
	if err != nil {
		return err, nil
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		return errors.New("Couldn't parse claims"), nil
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("JWT is expired"), nil
	}

	return nil, claims
}
