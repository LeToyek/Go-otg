package db

import (
	"context"
	"database/sql"
	"go-otg/internal/repository/db/constants"
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
		args string
		err  error
		want User
		mock func(mock mockFields)
	}{
		{
			name: "case error getting user from DB, should return the error",
			args: "1",
			err:  assert.AnError,
			mock: func(mock mockFields) {
				mock.sql.ExpectQuery(constants.GetUserById).
					WithArgs("1").
					WillReturnError(assert.AnError)
			},
		},
		{
			name: "case success then return the result",
			args: "1",
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
				mock.sql.ExpectQuery(constants.GetUserById).
					WithArgs("1").
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

func TestRepository_CreateUser(t *testing.T) {
	type mockFields struct {
		sqlMock sqlmock.Sqlmock
	}
	tests := []struct {
		name    string
		args    User
		want    string
		errWant error
		mock    func(mock mockFields)
	}{
		{name: "success create an user",
			args: User{
				ID:       "user1",
				Username: "username1",
				Email:    "user@gmail.com",
				Password: "user123",
				Address:  "jl.Gayam",
				Age: sql.NullInt16{
					Int16: 13,
					Valid: true,
				},
			},
			want:    "1",
			errWant: nil,
			mock: func(mock mockFields) {
				mock.sqlMock.
					ExpectExec(`INSERT INTO user("id","username","email","password","address","age") 
					VALUES ($1,$2,$3,$4,$5,$6)`).
					WithArgs("user1", "username1", "user@gmail.com", "user123", "jl.Gayam", 13).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "failed to create an user",
			args: User{
				ID:       "user2",
				Username: "username1",
				Email:    "user@gmail.com",
				Password: "user123",
				Address:  "jl.Gayam",
				Age: sql.NullInt16{
					Int16: 13,
					Valid: true,
				},
			},
			want:    "",
			errWant: assert.AnError,
			mock: func(mock mockFields) {
				mock.sqlMock.
					ExpectExec(`INSERT INTO user("id","username","email","password","address","age") 
					VALUES ($1,$2,$3,$4,$5,$6)`).
					WithArgs("user2", "username1", "user@gmail.com", "user123", "jl.Gayam", 13).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, test := range tests {
		mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer mockDB.Close()

		mockField := mockFields{
			sqlMock: mockSQL,
		}

		test.mock(mockField)
		repository := &Repository{
			db: sqlx.NewDb(mockDB, "pq"),
		}

		got, err := repository.CreateUser(context.Background(), test.args)

		assert.Equal(t, test.want, got)
		assert.Equal(t, test.errWant, err)
	}
}
