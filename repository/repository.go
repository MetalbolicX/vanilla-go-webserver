package repository

import (
	"context"
)

// Defines an interface Repository and a set of functions
// that act as a wrapper around the repository implementation.
// Repository defines a set of methods. Each method
// represents a database operation such as closing
// the connection, executing a command/query, inserting data
// , retrieving data, updating data, or deleting data.
type Repository interface {
	Close() error
	ExecuteCommand(ctx context.Context, query string, args ...any) ([]map[string]any, error)
	Post(ctx context.Context, query string, args ...any) error
	Get(ctx context.Context, query string, args ...any) ([]map[string]any, error)
	Put(ctx context.Context, query string, args ...any) (int64, error)
	Delete(ctx context.Context, query string, args ...any) (int64, error)
}

var implementation Repository

// SetRepository is a function that sets the global
// implementation variable to the provided repository.
// It allows the application to set the implementation
// of the Repository interface.
func SetRepository(repository Repository) {
	implementation = repository
}

// Close is a function that delegates the call to the Close
// method of the underlying repository implementation.
// It closes the connection to the database.
func Close() error {
	return implementation.Close()
}

// ExecuteCommand is a function that delegates the call
// to the ExecuteCommand method of the underlying repository
// implementation. It executes a database command/query
// and returns the result as a slice of maps,
// where each map represents a row of data.
func ExecuteCommand(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	return implementation.ExecuteCommand(ctx, query, args...)
}

// The Post function adds new data in the database.
func Post(ctx context.Context, query string, args ...any) error {
	return implementation.Post(ctx, query, args...)
}

// The Get function reads data from the database.
func Get(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	return implementation.Get(ctx, query, args...)
}

// The Put function updated data from the database.
func Put(ctx context.Context, query string, args ...any) (int64, error) {
	return implementation.Put(ctx, query, args...)
}

// The Delete function deletes data from the database.
func Delete(ctx context.Context, query string, args ...any) (int64, error) {
	return implementation.Delete(ctx, query, args...)
}
