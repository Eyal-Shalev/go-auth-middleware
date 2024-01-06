package authMiddleware

import (
	"context"
	"net/http"
)

type BasicAuthFunc[T any] func(ctx context.Context, userName, password string) (T, bool, error)

func (fn BasicAuthFunc[T]) Wrap(next http.Handler) http.Handler {
	return AuthenticateFunc[T](func(r *http.Request) (value T, ok bool, err error) {
		userName, password, ok := r.BasicAuth()
		if !ok {
			var zero T
			return zero, false, nil
		}

		return fn(r.Context(), userName, password)
	}).Wrap(next)
}

func BasicAuth[T any](fn BasicAuthFunc[T]) Middleware {
	return fn.Wrap
}
