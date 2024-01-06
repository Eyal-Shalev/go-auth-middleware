package authMiddleware

import (
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

type AuthenticateFunc[T any] func(r *http.Request) (value T, ok bool, err error)

func NewAuthenticateMiddleware[T any](fn AuthenticateFunc[T]) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v, ok, err := fn(r)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			ctx := r.Context()
			if ok {
				ctx = SetValue(ctx, v)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

type AuthorizeFunc[T any] func(r *http.Request, value T) bool

func NewAuthorizationMiddleware[T any](fn AuthorizeFunc[T]) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v, ok := GetValue[T](r.Context())
			if !ok || !fn(r, v) {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
