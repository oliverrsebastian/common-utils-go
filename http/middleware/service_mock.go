package middleware

import "context"

type mockService struct {
	err error
}

func newMockService(err error) Service {
	return &mockService{
		err: err,
	}
}

func (s mockService) Check(context.Context, string) (*User, error) {
	return &User{}, s.err
}

func (s mockService) GetAccess(context.Context, int64, ResourceCode) error {
	return s.err
}
