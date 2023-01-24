package db

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetCommonByID(t *testing.T) {
	type mockFields struct {
		sql sqlmock.Sqlmock
	}
	tests := []struct {
		name    string
		args    int64
		want    Common
		wantErr error
		// mock    func(mock mockFields)
	}{
		{
			name: "Success get data",
			args: 1,
			want: Common{
				ID:      1,
				Message: "Halo gan!",
			},
			wantErr: nil,
			// mock: func(mock mockFields) {
			// 	mock.sql.ExpectQuery("SELECT * FROM common WHERE id = $1").
			// 		WithArgs(1).
			// 		WillReturnRows(sqlmock.NewRows([]string{"id", "message"}).
			// 			AddRow(1, "Halo gan"))
			// },
		},
		// {
		// 	name: "Wrong returning value",
		// 	args: 1,
		// 	want: Common{
		// 		ID:      1,
		// 		Message: "Artinya apa bang messi",
		// 	},
		// 	wantErr: assert.AnError,
		// 	// mock: func(mock mockFields) {
		// 	// 	mock.sql.ExpectQuery("SELECT * FROM common WHERE id = $1").
		// 	// 		WithArgs(1).
		// 	// 		WillReturnError(assert.AnError)
		// 	// },
		// },
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				panic(err)
			}
			fmt.Println(mockSQL)
			// assert.NoError(t, err)
			// defer mockDB.Close()

			// mockFields := mockFields{
			// 	sql: mockSQL,
			// }
			// test.mock(mockFields)
			repository := &Repository{
				db: sqlx.NewDb(mockDB, "pq"),
			}
			// got, err := repository.GetCommonByID(test.args)
			got, err := repository.GetCommonByID(test.args)
			assert.Equal(t, test.want, got)
			// assert.Equal(t, test.wantErr, err)
		})
	}
}
