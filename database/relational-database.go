package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type relationalDBRepo struct {
	db *sql.DB
}

func NewRelationalDBRepo(dbManagementSystem, url string) (*relationalDBRepo, error) {
	db, err := sql.Open(dbManagementSystem, url)
	if err != nil {
		return nil, err
	}
	return &relationalDBRepo{db}, nil
}

func (d *relationalDBRepo) Close() error {
	return d.db.Close()
}

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

func (d *relationalDBRepo) Post(ctx context.Context, query string, args ...any) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (d *relationalDBRepo) Get(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	return d.ExecuteCommand(ctx, query, args...)
}

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
