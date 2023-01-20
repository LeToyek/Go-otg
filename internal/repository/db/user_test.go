package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetUserByID(t *testing.T) {
	type mockFields struct {
		sql sqlmock.Sqlmock
	}

	tests := []struct {
		name string
		args int64
		err  error
		want User
		mock func(mock mockFields)
	}{
		{
			name: "case error getting user from DB, should return the error",
			args: 1,
			err:  assert.AnError,
			mock: func(mock mockFields) {
				mock.sql.ExpectQuery("select * from users where id = $1").
					WithArgs(1).
					WillReturnError(assert.AnError)
			},
		},
		{
			name: "case success then return the result",
			args: 1,
			want: User{
				ID:       "1",
				Username: "username",
				Email:    "email",
				Age: sql.NullInt16{
					Int16: 20,
					Valid: true,
				},
			},
			mock: func(mock mockFields) {
				mock.sql.ExpectQuery("select * from users where id = $1").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "age"}).
						AddRow("1", "username", "email", 20))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			assert.NoError(t, err)
			defer mockDB.Close()

			mockFields := mockFields{
				sql: mockSQL,
			}
			test.mock(mockFields)
			repository := &Repository{
				db: sqlx.NewDb(mockDB, "pq"),
			}

			got, err := repository.GetUserByID(context.Background(), test.args)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.err, err)
		})
	}
}
