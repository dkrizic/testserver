package token

import (
	"errors"
	"github.com/dkrizic/testserver/service/handler/errorh"
	"github.com/dkrizic/testserver/telemetry"
	"github.com/golang-jwt/jwt/v4"
	"go.opentelemetry.io/otel/attribute"
	"log/slog"
	"net/http"
	"strings"
)

type Config struct {
	checkToken bool
	alwaysPass AlwaysPass
}

type AlwaysPass func(r *http.Request) bool

func NewTokenHandler(options ...func(config *Config)) *Config {
	c := &Config{
		checkToken: false,
	}
	for _, option := range options {
		option(c)
	}
	slog.Debug("JWT Handler configured", "checkToken", c.checkToken)
	return c
}

func WithCheckToken(checkToken bool) func(*Config) {
	return func(c *Config) {
		c.checkToken = checkToken
	}
}

func WithAlwaysPass(ap AlwaysPass) func(*Config) {
	return func(c *Config) {
		c.alwaysPass = ap
	}
}

// http handler that checks for a valid jwt token
// and passes the request to the next handler
// if the token is valid
// if the token is not valid, it will return a 401
// unauthorized
func (c Config) TokenHandler(handler http.Handler) http.Handler {
	return errorh.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		if c.alwaysPass != nil && c.alwaysPass(r) {
			handler.ServeHTTP(w, r)
			return nil
		}

		spanName := "TokenHandler"
		if !c.checkToken {
			spanName = "TokenHandler (disabled)"
		}
		ctx, span := telemetry.Tracer().Start(r.Context(), spanName)
		defer span.End()
		r = r.WithContext(ctx)

		if c.checkToken {
			slog.DebugContext(ctx, "Checking JWT token")
			span.AddEvent("JWT check enabled")
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader == "" {
				return errors.New("Authorization header missing")
			}

			if !strings.HasPrefix(authorizationHeader, "Bearer ") {
				return errors.New("Authorization header missing")
			}

			tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
			if tokenString == "" {
				return errors.New("JWT token missing")
			}

			token, _, err := jwt.NewParser(jwt.WithoutClaimsValidation()).ParseUnverified(tokenString, jwt.MapClaims{})
			if err != nil {
				return errors.New("JWT token invalid: " + err.Error())
			}
			upn := token.Claims.(jwt.MapClaims)["upn"]
			if upn == nil {
				return errors.New("JWT token invalid: upn missing")
			}

			// set enduser.id to uid
			span.SetAttributes(attribute.String("enduser.id", upn.(string)))

			span.AddEvent("JWT token valid")
			slog.DebugContext(ctx, "JWT token valid", slog.Any("token", token))

			handler.ServeHTTP(w, r)

			return nil
		} else {
			span.AddEvent("JWT check disabled")
			slog.WarnContext(r.Context(), "JWT check disabled")
			handler.ServeHTTP(w, r)
			return nil
		}
	})
}
