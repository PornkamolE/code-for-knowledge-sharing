package memory

import (
	"context"
	"sync"
	"errors"

	"my-knowledge-sharing/app/core/ports"
)

type UserRepoMemory struct {
	mu   sync.Mutex
	seq  int64
	data map[int64]string
}

func NewUserRepoMemory() ports.UserRepositoryPort {
	return &UserRepoMemory{
		data: make(map[int64]string),
	}
}

func (r *UserRepoMemory) CreateUser(ctx context.Context, name string) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.seq++
	r.data[r.seq] = name
	return r.seq, nil
}

func (r *UserRepoMemory) FindByID(ctx context.Context, id int64) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	name, ok := r.data[id]
	if !ok {
		return "", errors.New("user not found")
	}
	return name, nil
}