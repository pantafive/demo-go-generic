package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" //nolint:revive,nolintlint
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  uint   `db:"age"`
}

const schema = `
CREATE TABLE user
(
  id   INTEGER PRIMARY KEY,
  name TEXT    NOT NULL,
  age  INTEGER NOT NULL
);`

type Iterator[T any] struct {
	iterate func() (T, error)
}

func (i *Iterator[T]) Next() (T, error) {
	val, err := i.iterate()
	return val, err
}

type Repository struct {
	*sqlx.DB
}

func (r Repository) ReadUsers() (*Iterator[*User], error) {
	rows, err := r.Queryx("SELECT id, name, age FROM user;")
	if err != nil {
		return nil, err
	}

	i := &Iterator[*User]{}

	i.iterate = func() (*User, error) {
		user := User{}

		if rows.Next() {
			return &user, scanAndCloseOnError(rows, &user)
		}

		if err := rows.Close(); err != nil {
			return nil, err
		}

		return nil, nil //nolint:nilnil
	}

	return i, nil
}

func scanAndCloseOnError[T any](rows *sqlx.Rows, v T) error {
	if err := rows.StructScan(&v); err != nil {
		if err := rows.Close(); err != nil {
			return err
		}
		return err
	}
	return nil
}
