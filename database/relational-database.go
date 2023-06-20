package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// The relationalDBRepo represents a repository
// implementation for a relational database. It has a
// field db of type *sql.DB, which represents the
// connection to the database.
type relationalDBRepo struct {
	db *sql.DB
}

// The NewRelationalDBRepo function creates a new instance
// of the relationalDBRepo struct. It takes the database
// management system (e.g., MySQL, PostgreSQL) and the
// database connection URL as parameters. It uses sql.Open
// to establish a connection to the database.
// If an error occurs during the connection process,
// it returns nil and the error. Otherwise, it returns
// a pointer to the created relationalDBRepo instance.
func NewRelationalDBRepo(dbManagementSystem, url string) (*relationalDBRepo, error) {
	db, err := sql.Open(dbManagementSystem, url)
	if err != nil {
		return nil, err
	}
	return &relationalDBRepo{db}, nil
}

// The Close method is part of the Repository interface
// implementation. It closes the underlying database
// connection by invoking the Close method of the sql.DB
// struct.
func (d *relationalDBRepo) Close() error {
	return d.db.Close()
}

// The ExecuteCommand method is part of the Repository
// interface implementation. It executes a database query
// with the provided query string and arguments.
// It uses QueryContext to execute the query and retrieve
// the resulting rows. It then scans the rows and maps
// the column values to a slice of maps. Each map
// represents a row, with the column names as keys and
// the corresponding values.
func (d *relationalDBRepo) ExecuteCommand(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := make([]map[string]any, 0)
	values := make([]any, len(columns))
	valuePointers := make([]any, len(columns))
	for index := range values {
		valuePointers[index] = &values[index]
	}

	for rows.Next() {
		if err := rows.Scan(valuePointers...); err != nil {
			return nil, err
		}

		row := make(map[string]any)
		for index, column := range columns {
			row[column] = values[index]
			// You can perform any necessary type assertions or conversions here.
		}

		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// The Post method inserts data into the database using
// the provided query and arguments. It creates a new
// context with a timeout of 5 seconds. It then calls
// ExecContext on the underlying sql.DB object to execute
// the query. If an error occurs, it is returned.
// Otherwise, it returns nil to indicate success.
func (d *relationalDBRepo) Post(ctx context.Context, query string, args ...any) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// The Get method retrieves data from the database
// using the provided query and arguments. It simply calls
// the ExecuteCommand method, which executes the query
// and returns the result as a slice of maps representing
// the rows.
func (d *relationalDBRepo) Get(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	return d.ExecuteCommand(ctx, query, args...)
}

// The Put method updates data in the database using the
// provided query and arguments. It follows a similar
// pattern as the Post method, but additionally retrieves
// the number of rows affected by the update operation.
// It returns the number of affected rows and any error
// that occurred during execution.
func (d *relationalDBRepo) Put(ctx context.Context, query string, args ...any) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// The Delete method removes data from the database using
// the provided query and arguments. It follows a similar
// pattern as the Put method, returning the number of
// affected rows and any error that occurred during
// execution.
func (d *relationalDBRepo) Delete(ctx context.Context, query string, args ...any) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
