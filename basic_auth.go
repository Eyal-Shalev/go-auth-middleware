package authMiddleware

import (
	"net/http"
)

type BasicAuthAuthenticator[T any] func(userName, password string) (T, bool, error)

func NewBasicAuthenticateMiddleware[T any](fn BasicAuthAuthenticator[T]) Middleware {
	return NewAuthenticateMiddleware(func(r *http.Request) (value T, ok bool, err error) {
		userName, password, ok := r.BasicAuth()
		if !ok {
			return
		}
		return fn(userName, password)
	})
}
