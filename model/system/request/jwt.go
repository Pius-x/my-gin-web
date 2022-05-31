package request

import "github.com/golang-jwt/jwt/v4"

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	ID       int64
	Account  string
	Password string
}
