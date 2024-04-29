package middleware

import "context"

var authService Service

type Service interface {
	Check(ctx context.Context, token string) (*User, error)
	GetAccess(ctx context.Context, userID int64, code ResourceCode) error
}

func GetService() Service {
	return authService
}
