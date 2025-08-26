package middleware

import (
	"net/http"
)

func RateLimit(next http.Handler) http.Handler {
	// Rate limiting middleware
	return next
}
