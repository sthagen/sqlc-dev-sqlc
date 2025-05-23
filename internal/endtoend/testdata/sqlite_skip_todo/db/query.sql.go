// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const getFoo = `-- name: GetFoo :many
SELECT bar FROM foo
WHERE bar = ?
`

func (q *Queries) GetFoo(ctx context.Context, bar sql.NullString) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, getFoo, bar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var bar sql.NullString
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFoo = `-- name: ListFoo :many
SELECT bar FROM foo
`

func (q *Queries) ListFoo(ctx context.Context) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, listFoo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var bar sql.NullString
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
