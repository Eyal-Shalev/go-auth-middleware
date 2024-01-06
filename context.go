package authMiddleware

import "context"

type contextKey string

const valueKey = contextKey("value")

func SetValue[T any](ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, valueKey, value)
}

func GetValue[T any](ctx context.Context) (T, bool) {
	v, ok := ctx.Value(valueKey).(T)
	if !ok {
		var zero T
		return zero, false
	}
	return v, true
}
