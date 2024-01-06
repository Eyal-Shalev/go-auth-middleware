package authMiddleware

import (
	"context"
	"net/http"
)

type BasicAuthenticatorFunc[T any] func(ctx context.Context, userName, password string) (T, bool, error)

func NewBasicAuthenticateMiddleware[T any](fn BasicAuthenticatorFunc[T]) Middleware {
	return NewAuthenticateMiddleware(func(r *http.Request) (value T, ok bool, err error) {
		userName, password, ok := r.BasicAuth()
		if !ok {
			return
		}
		return fn(r.Context(), userName, password)
	})
}
