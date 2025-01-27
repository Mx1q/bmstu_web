package web

import (
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"ppo/domain"

	"github.com/go-chi/jwtauth/v5"
)

func ValidateAdminRoleJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			ErrorResponse(w, fmt.Errorf("getting claims from JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		role, ok := claims["role"]
		if !ok {
			ErrorResponse(w, fmt.Errorf("getting 'role' claim from JWT").Error(), http.StatusBadRequest)
			return
		}

		if role != "admin" {
			ErrorResponse(w, fmt.Errorf("only administrators can use this").Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ValidateUserRoleJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			ErrorResponse(w, fmt.Errorf("getting claims from JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		role, ok := claims["role"]
		if !ok {
			ErrorResponse(w, fmt.Errorf("getting 'role' claim from JWT").Error(), http.StatusBadRequest)
			return
		}

		if role != domain.DefaultRole && role != "admin" {
			ErrorResponse(w, fmt.Errorf("you need to login to perform this action").Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Authenticator(ja *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())

			// no token found
			if err != nil || token == nil {
				next.ServeHTTP(w, r)
				return
			}

			if jwt.Validate(token, ja.ValidateOptions()...) != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// Token is authenticated, pass it through
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}

func OnlyRatedFlagIsRaised(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		onlyRated := r.URL.Query().Get("onlyRated")
		if onlyRated != "true" {
			//ErrorResponse(w, fmt.Errorf("it's an only rated route").Error(), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
