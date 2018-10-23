package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/vetebase/husky"
)

// JWTConfig holds token information
type JWTConfig struct {
	SigningKey    interface{}
	SigningMethod string // Defaults to HS256
	Claims        jwt.Claims
}

const (
	// AlgoHS256 is 256 bit encryption
	AlgoHS256 = "HS256"

	// AlgoHS512 is 512 bit encryption
	AlgoHS512 = "HS512"

	// AlgoHS1024 is 1024 bit encryptions
	AlgoHS1024 = "HS1024"
)

// TokenParser parses out token
type TokenParser func(*husky.CTX) (string, error)

// DefaultJWT is the default settings for the JWT
var DefaultJWT = JWTConfig{
	SigningKey:    "secret",
	SigningMethod: AlgoHS256,
}

// JWTError returns a husky Handler when an error is occured
func JWTError(c husky.CTX) error {
	return c.JSON(500, "JWT Error")
}

// JWT default json web token handler
func JWT() func(next husky.Handler) husky.Handler {
	return func(next husky.Handler) husky.Handler {
		return func(ctx *husky.CTX) error {
			husky := husky.New()
			config := husky.Config.Load()

			j := DefaultJWT
			if _, ok := config["JWT_SECRET"]; ok {
				j.SigningKey = config["JWT_SECRET"]
			}

			if _, ok := config["JWT_METHOD"]; ok {
				j.SigningMethod = config["JWT_METHOD"]
			}

			parser := parseFromHeader()
			parsed, err := parser(ctx)
			if err != nil {
				log.Println(err)
			}

			token := new(jwt.Token)
			if _, ok := j.Claims.(jwt.MapClaims); ok {
				token, err = jwt.Parse(parsed, func(token *jwt.Token) (interface{}, error) {
					return []byte("AllYourBase"), nil
				})
			}

			if token.Valid && err == nil {
				return next(ctx)
			}

			return ctx.Code(http.StatusNoContent)
		}
	}
}

func parseFromHeader() TokenParser {
	return func(ctx *husky.CTX) (string, error) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			return "", errors.New("JWT token is missing")
		}

		return token, nil
	}
}
