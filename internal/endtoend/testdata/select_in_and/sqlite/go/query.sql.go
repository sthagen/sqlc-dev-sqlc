// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM
  books AS b
WHERE
  b.author NOT IN (
    SELECT
      a.name
    FROM
      authors a
    WHERE
      a.age >= ?
  )
  AND b.translator NOT IN (
    SELECT
      t.name
    FROM
      translators t
    WHERE
      t.age >= ?
  )
  AND b.year <= ?
`

type DeleteAuthorParams struct {
	Age   sql.NullInt64
	Age_2 sql.NullInt64
	Year  sql.NullInt64
}

func (q *Queries) DeleteAuthor(ctx context.Context, arg DeleteAuthorParams) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, arg.Age, arg.Age_2, arg.Year)
	return err
}
