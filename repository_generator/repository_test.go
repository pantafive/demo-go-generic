package repository

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDB *sqlx.DB //nolint:gochecknoglobals

func TestMain(m *testing.M) {
	var err error
	testDB, err = sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	testDB.MustExec(schema)
	testDB.MustExec("INSERT INTO user (name, age) VALUES ($1, $2)", "Alice", 18)
	testDB.MustExec("INSERT INTO user (name, age) VALUES ($1, $2)", "Bob", 21)
	testDB.MustExec("INSERT INTO user (name, age) VALUES ($1, $2)", "Charlie", 23)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestRepository_ReadUsersAge(t *testing.T) {
	repository := &Repository{testDB}
	users, err := repository.ReadUsers()
	require.NoError(t, err)

	for user, err := users.Next(); user != nil || err != nil; user, err = users.Next() {
		require.NoError(t, err)
		assert.GreaterOrEqual(t, user.Age, uint(18))
	}
}
