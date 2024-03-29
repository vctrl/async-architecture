package model

import "context"

type SessionManager interface {
	Create(ctx context.Context, userID, login, role string) (string, error)
	Check(ctx context.Context, token string) (*Session, error)
	Destroy(ctx context.Context, id string) error
	DestroyAll(ctx context.Context) error
}
