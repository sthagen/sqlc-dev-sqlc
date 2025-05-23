// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const selectAllJoined = `-- name: SelectAllJoined :many
select events.id from events
    inner join handled_events
       on events.ID > handled_events.last_handled_id
where handled_events.handler = $1
    for update of handled_events skip locked
`

func (q *Queries) SelectAllJoined(ctx context.Context, handler sql.NullString) ([]sql.NullInt32, error) {
	rows, err := q.db.QueryContext(ctx, selectAllJoined, handler)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullInt32
	for rows.Next() {
		var id sql.NullInt32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectAllJoinedAlias = `-- name: SelectAllJoinedAlias :many
select e.id from events e
    inner join handled_events he
       on e.ID > he.last_handled_id
where he.handler = $1
    for update of he skip locked
`

func (q *Queries) SelectAllJoinedAlias(ctx context.Context, handler sql.NullString) ([]sql.NullInt32, error) {
	rows, err := q.db.QueryContext(ctx, selectAllJoinedAlias, handler)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullInt32
	for rows.Next() {
		var id sql.NullInt32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
