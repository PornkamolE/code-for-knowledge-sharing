package ports

import "context"

type UserRepositoryPort interface {
	CreateUser(ctx context.Context, name string) (int64, error)
	FindByID(ctx context.Context, id int64) (string, error)
}