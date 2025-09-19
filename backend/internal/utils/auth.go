package utils

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang-jwt/jwt/v5"

	middleware "github.com/oapi-codegen/nethttp-middleware"
)

var jwtKey = []byte("supersecretkey") // move to env

func GenerateJWT(userID string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	return claims.Subject, nil
}

func AuthMiddleware(spec *openapi3.T) func(http.Handler) http.Handler {
	// spec.Servers = nil
	return middleware.OapiRequestValidatorWithOptions(spec, &middleware.Options{
		SilenceServersWarning: true,
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
				if ai.SecuritySchemeName == "BearerAuth" {

					req := ai.RequestValidationInput.Request
					authHdr := req.Header.Get("Authorization")

					if authHdr == "" {
						return errors.New("authorization header is missing")
					}

					// TODO: Validate token
					if false {
						return nil
					}

					return errors.New("invalid token")
				}
				return nil
			},
		},
	})
}
