// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"database/sql"
)

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type Book struct {
	ID       int64
	AuthorID int64
	Title    string
}