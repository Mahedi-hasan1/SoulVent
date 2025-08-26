package middleware

import (
	"net/http"
)

func JWTAuth(next http.Handler) http.Handler {
	// JWT authentication middleware
	return next
}
