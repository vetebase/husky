package middleware

import (
	"net/http"
	"strings"

	"github.com/vetebase/husky"
)

// CORSConfig configuration for CORS middleware
type CORSConfig struct {
	AllowedOrigins []string `json:"allowed_origins"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedHeaders []string `json:"allowed_headers"`
	ExposedHeaders []string `json:"exposed_headers"`
}

// DefaultCORSConfig handles the default CORS configuration for Huksy
var DefaultCORSConfig = CORSConfig{
	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders: []string{"*"},
	ExposedHeaders: []string{"*"},
}

// CORSError returns a Husky Handler when an error is occured
func CORSError(ctx *husky.CTX) error {
	return ctx.JSON(500, "CORS Error")
}

// CORS middleware for Husky routes
func CORS(config CORSConfig) func(next husky.Handler) husky.Handler {
	return CORSConfigured(config)
}

// CORSConfigured returns a configured CORS middleware
// func CORSConfigured(ctx husky.Context, handler husky.Handler, config CORSConfig) husky.Handler {
func CORSConfigured(config CORSConfig) func(next husky.Handler) husky.Handler {
	if len(config.AllowedOrigins) == 0 {
		config.AllowedOrigins = DefaultCORSConfig.AllowedOrigins
	}

	if len(config.AllowedMethods) == 0 {
		config.AllowedMethods = DefaultCORSConfig.AllowedMethods
	}

	if len(config.AllowedHeaders) == 0 {
		config.AllowedHeaders = DefaultCORSConfig.AllowedHeaders
	}

	if len(config.ExposedHeaders) == 0 {
		config.ExposedHeaders = DefaultCORSConfig.ExposedHeaders
	}

	allowedOrigins := strings.Join(config.AllowedOrigins, ",")
	allowedMethods := strings.Join(config.AllowedMethods, ",")
	allowedHeaders := strings.Join(config.AllowedHeaders, ",")
	exposedHeaders := strings.Join(config.ExposedHeaders, ",")

	return func(next husky.Handler) husky.Handler {
		return func(ctx *husky.CTX) error {
			req := ctx.Request
			res := ctx.Response

			if req.Method != "OPTIONS" {
				res.Header().Add("Vary", "Origin")
				res.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
				if exposedHeaders != "" {
					res.Header().Set("Access-Control-Expose-Headers", exposedHeaders)
				}

				return next(ctx)
			}

			res.Header().Add("Vary", "Origin")
			res.Header().Add("Vary", "Access-Control-Request-Method")
			res.Header().Add("Vary", "Access-Control-Request-Headers")
			res.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
			res.Header().Set("Access-Control-Allow-Methods", allowedMethods)
			res.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

			return ctx.Code(http.StatusNoContent)
		}
	}
}
