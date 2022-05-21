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

type Repository struct {
	*sqlx.DB
}

type Iteration[T any] struct {
	Value T
	Err   error
}

func (r Repository) ReadUsers() chan Iteration[User] {
	c := make(chan Iteration[User], 1)

	go func() {
		rows, err := r.Queryx("SELECT id, name, age FROM user;")
		if err != nil {
			c <- Iteration[User]{Err: err}
			close(c)
		}

		defer func(rows *sqlx.Rows) {
			err := rows.Close()
			if err != nil {
				c <- Iteration[User]{Err: err}
			}
			close(c)
		}(rows)

		for rows.Next() {
			var user User
			err := scanAndCloseOnError(rows, &user)
			c <- Iteration[User]{Err: err, Value: user}
		}
	}()

	return c
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
