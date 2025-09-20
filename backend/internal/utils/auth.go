package utils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Mroxny/slamIt/internal/config"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang-jwt/jwt/v5"

	middleware "github.com/oapi-codegen/nethttp-middleware"
)

type contextKey string

const JWTClaimsContextKey contextKey = "userID"

var (
	ErrNoAuthHeader      = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("authorization header is malformed")
	ErrClaimsInvalid     = errors.New("provided claims do not match expected scopes")
)

var jwtKey = getJwtKey()

func getJwtKey() []byte {
	cfg, _ := config.GetConfig()
	return []byte(cfg.JWT.Key)
}

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

func GetJWSFromRequest(req *http.Request) (string, error) {
	authHdr := req.Header.Get("Authorization")
	if authHdr == "" {
		return "", ErrNoAuthHeader
	}
	prefix := "Bearer "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimPrefix(authHdr, prefix), nil
}

func AuthMiddleware(spec *openapi3.T) func(http.Handler) http.Handler {
	// spec.Servers = nil
	return middleware.OapiRequestValidatorWithOptions(spec, &middleware.Options{
		SilenceServersWarning: true,
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
				if ai.SecuritySchemeName == "bearerAuth" {
					req := ai.RequestValidationInput.Request
					jws, err := GetJWSFromRequest(req)
					if err != nil {
						return fmt.Errorf("getting JWS: %w", err)
					}

					userId, err := ValidateJWT(jws)
					if err != nil {
						return fmt.Errorf("validating JWS: %w", err)
					}

					newCtx := context.WithValue(req.Context(), JWTClaimsContextKey, userId)
					*req = *req.WithContext(newCtx)
				}
				return nil
			},
		},
	})
}
