package repository

import (
	"context"
)

type Repository interface {
	Close() error
	ExecuteCommand(ctx context.Context, query string, args ...any) ([]map[string]any, error)
	Post(ctx context.Context, query string, args ...any) error
	Get(ctx context.Context, query string, args ...any) ([]map[string]any, error)
	Put(ctx context.Context, query string, args ...any) (int64, error)
	Delete(ctx context.Context, query string, args ...any) (int64, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func Close() error {
	return implementation.Close()
}

func ExecuteCommand(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	return implementation.ExecuteCommand(ctx, query, args...)
}

func Post(ctx context.Context, query string, args ...any) error {
	return implementation.Post(ctx, query, args...)
}

func Get(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	return implementation.Get(ctx, query, args...)
}

func Put(ctx context.Context, query string, args ...any) (int64, error) {
	return implementation.Put(ctx, query, args...)
}

func Delete(ctx context.Context, query string, args ...any) (int64, error) {
	return implementation.Delete(ctx, query, args...)
}
