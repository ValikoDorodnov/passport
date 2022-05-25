package jwt

import (
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"os"
	"time"
)

func BuildToken() []byte {
	tok, err := jwt.NewBuilder().
		Issuer(`some-issuer`).
		IssuedAt(time.Now()).
		Build()

	if err != nil {
		fmt.Printf("failed to build token: %s\n", err)
		return nil
	}

	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.HS256, privateToken()))
	if err != nil {
		fmt.Printf("failed to sign token: %s\n", err)
		return nil
	}

	return signed
}

func ParseToken(token []byte) jwt.Token {
	verifiedToken, err := jwt.Parse(token, jwt.WithKey(jwa.HS256, privateToken()))
	if err != nil {
		fmt.Printf("failed to verify JWS: %s\n", err)
		return nil
	}

	return verifiedToken
}

func privateToken() jwk.Key {
	secret := os.Getenv("SECRET")

	key, err := jwk.FromRaw([]byte(secret))
	if err != nil {
		fmt.Printf("failed to create key: %s\n", err)
		return nil
	}

	return key
}
