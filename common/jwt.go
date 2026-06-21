package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrJWTExpired   = errors.New("jwt token has expired")
	ErrJWTInvalid   = errors.New("jwt token is invalid")
	ErrJWTNotIssued = errors.New("jwt token not yet valid")
)

// JWTClaims represents the custom claims embedded in the JWT
type JWTClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	Status   int    `json:"status"`
	Group    string `json:"group"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a signed JWT token with the given user claims
func GenerateJWT(userID int, username string, role int, status int, group string) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		Status:   status,
		Group:    group,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(JWTExpirationSeconds) * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}

// ParseJWT parses and validates a JWT token string, returning the claims
func ParseJWT(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(JWTSecret), nil
		})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrJWTExpired
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, ErrJWTNotIssued
		}
		return nil, ErrJWTInvalid
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrJWTInvalid
}

// GetJWTFromHeader extracts the JWT token from the Authorization header
// Expects format: "Bearer <token>"
func GetJWTFromHeader(authHeader string) string {
	const prefix = "Bearer "
	if len(authHeader) > len(prefix) && authHeader[:len(prefix)] == prefix {
		return authHeader[len(prefix):]
	}
	return ""
}
